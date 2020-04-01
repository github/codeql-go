package extractor

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/scanner"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/github/codeql-go/extractor/dbscheme"
	"github.com/github/codeql-go/extractor/srcarchive"
	"github.com/github/codeql-go/extractor/trap"
	"golang.org/x/tools/go/packages"
)

// Extract extracts the packages specified by the given patterns
func Extract(patterns []string) error {
	return ExtractWithFlags(nil, patterns)
}

// ExtractWithFlags extracts the packages specified by the given patterns and build flags
func ExtractWithFlags(buildFlags []string, patterns []string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps |
			packages.NeedTypes | packages.NeedTypesSizes |
			packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: buildFlags,
	}
	pkgs, err := packages.Load(cfg, patterns...)

	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		log.Printf("No packages found.")
	}

	extractUniverseScope()

	// recursively visit all packages in depth-first order;
	// on the way down, associate each package scope with its corresponding package,
	// and on the way up extract the package's scope
	packages.Visit(pkgs, func(pkg *packages.Package) bool {
		return true
	}, func(pkg *packages.Package) {
		if len(pkg.Errors) != 0 {
			log.Printf("Warning: encountered errors extracting package `%s`:", pkg.PkgPath)
			for _, err := range pkg.Errors {
				log.Printf("  %s", err.Error())
			}
		}

		tw, err := trap.NewWriter(pkg.PkgPath, pkg)
		if err != nil {
			log.Fatal(err)
		}
		defer tw.Close()

		scope := extractPackageScope(tw, pkg)
		tw.ForEachObject(extractObjectType)
		lbl := tw.Labeler.GlobalID(pkg.PkgPath + ";pkg")
		dbscheme.PackagesTable.Emit(tw, lbl, pkg.Name, pkg.PkgPath, scope)
	})

	// this sets the number of threads that the Go runtime will spawn; this is separate
	// from the number of goroutines that the program spawns, which are scheduled into
	// the system threads by the Go runtime scheduler
	threads := os.Getenv("LGTM_THREADS")
	if maxprocs, err := strconv.Atoi(threads); err == nil && maxprocs > 0 {
		log.Printf("Max threads set to %d", maxprocs)
		runtime.GOMAXPROCS(maxprocs)
	} else if threads != "" {
		log.Printf("Warning: LGTM_THREADS value %s is not valid, defaulting to using all available threads.", threads)
	}
	// if the value is empty or not set, use the Go default, which is the number of cores
	// available since Go 1.5, but is subject to change

	var maxgoroutines int
	if maxgoroutines, err = strconv.Atoi(os.Getenv("SEMMLE_MAX_GOROUTINES")); err != nil {
		maxgoroutines = 32
	} else {
		log.Printf("Max goroutines set to %d", maxgoroutines)
	}

	var wg sync.WaitGroup
	// this semaphore is used to limit the number of files that are open at once;
	// this is to prevent the extractor from running into issues with caps on the
	// number of open files that can be held by one process
	fdSem := newSemaphore(100)
	// this semaphore is used to limit the number of goroutines spawned, so we
	// don't run into memory issues
	goroutineSem := newSemaphore(maxgoroutines)

	// extract AST information for all packages
	for _, pkg := range pkgs {
		extractPackage(pkg, &wg, goroutineSem, fdSem)
	}

	wg.Wait()

	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: unable to get working directory: %s", err.Error())
		log.Println("Skipping go.mod extraction")
	}
	rcwd, err := filepath.EvalSymlinks(cwd)
	if err == nil {
		cwd = rcwd
	}

	log.Printf("Walking file tree from %s to discover go.mod files...", cwd)

	filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) == "go.mod" && info != nil && info.Mode().IsRegular() {
			if err != nil {
				log.Printf("Found go.mod with path %s, but encountered error %s", path, err.Error())
			}

			log.Printf("Extracting %s", path)
			start := time.Now()

			err := extractGoMod(path)
			if err != nil {
				log.Printf("Failed to extract go.mod: %s", err.Error())
			}

			end := time.Since(start)
			log.Printf("Done extracting %s (%dms)", path, end.Nanoseconds()/1000000)
		}

		return nil
	})

	return nil
}

// extractUniverseScope extracts symbol table information for the universe scope
func extractUniverseScope() {
	tw, err := trap.NewWriter("universe", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tw.Close()

	lbl := tw.Labeler.ScopeID(types.Universe, nil)
	dbscheme.ScopesTable.Emit(tw, lbl, dbscheme.UniverseScopeType.Index())
	extractObjects(tw, types.Universe, lbl)
}

// extractObjects extracts all objects declared in the given scope
func extractObjects(tw *trap.Writer, scope *types.Scope, scopeLabel trap.Label) {
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		lbl, exists := tw.Labeler.ScopedObjectID(obj, extractType(tw, obj.Type()))
		if !exists {
			extractObject(tw, obj, lbl)
		}

		if obj.Parent() != scope {
			// this can happen if a scope is embedded into another with a `.` import.
			continue
		}
		dbscheme.ObjectScopesTable.Emit(tw, lbl, scopeLabel)
	}
}

// extractObject extracts a single object and emits it to the objects table.
func extractObject(tw *trap.Writer, obj types.Object, lbl trap.Label) {
	name := obj.Name()
	isBuiltin := obj.Parent() == types.Universe
	var kind int
	switch obj.(type) {
	case *types.PkgName:
		kind = dbscheme.PkgObjectType.Index()
	case *types.TypeName:
		if isBuiltin {
			kind = dbscheme.BuiltinTypeObjectType.Index()
		} else {
			kind = dbscheme.DeclTypeObjectType.Index()
		}
	case *types.Const:
		if isBuiltin {
			kind = dbscheme.BuiltinConstObjectType.Index()
		} else {
			kind = dbscheme.DeclConstObjectType.Index()
		}
	case *types.Nil:
		kind = dbscheme.BuiltinConstObjectType.Index()
	case *types.Var:
		kind = dbscheme.DeclVarObjectType.Index()
	case *types.Builtin:
		kind = dbscheme.BuiltinFuncObjectType.Index()
	case *types.Func:
		kind = dbscheme.DeclFuncObjectType.Index()
	case *types.Label:
		kind = dbscheme.LabelObjectType.Index()
	default:
		log.Fatalf("unknown object of type %T", obj)
	}
	dbscheme.ObjectsTable.Emit(tw, lbl, kind, name)

	// for methods, additionally extract information about the receiver
	if sig, ok := obj.Type().(*types.Signature); ok {
		if recv := sig.Recv(); recv != nil {
			recvlbl, exists := tw.Labeler.ReceiverObjectID(recv, lbl)
			if !exists {
				extractObject(tw, recv, recvlbl)
			}
			dbscheme.MethodReceiversTable.Emit(tw, lbl, recvlbl)
		}
	}
}

// extractObjectType extracts type and receiver information for a given object
func extractObjectType(tw *trap.Writer, obj types.Object, lbl trap.Label) {
	if tp := obj.Type(); tp != nil {
		dbscheme.ObjectTypesTable.Emit(tw, lbl, extractType(tw, tp))
	}
}

// extractPackage extracts AST information for all files in the given package
func extractPackage(pkg *packages.Package, wg *sync.WaitGroup,
	goroutineSem *semaphore, fdSem *semaphore) {
	for _, astFile := range pkg.Syntax {
		wg.Add(1)
		goroutineSem.acquire(1)
		go func(astFile *ast.File) {
			err := extractFile(astFile, pkg, fdSem)
			if err != nil {
				log.Fatal(err)
			}
			goroutineSem.release(1)
			wg.Done()
		}(astFile)
	}
}

// normalizedPath computes the normalized path (with symlinks resolved) for the given file
func normalizedPath(ast *ast.File, fset *token.FileSet) string {
	file := fset.File(ast.Package).Name()
	path, err := filepath.EvalSymlinks(file)
	if err != nil {
		return file
	}
	return path
}

// extractFile extracts AST information for the given file
func extractFile(ast *ast.File, pkg *packages.Package, fdSem *semaphore) error {
	fset := pkg.Fset
	path := normalizedPath(ast, fset)

	fdSem.acquire(3)

	log.Printf("Extracting %s", path)
	start := time.Now()

	defer fdSem.release(1)
	tw, err := trap.NewWriter(path, pkg)
	if err != nil {
		fdSem.release(2)
		return err
	}
	defer tw.Close()

	err = srcarchive.Add(path)
	fdSem.release(2)
	if err != nil {
		return err
	}

	extractFileInfo(tw, path)

	extractScopes(tw, ast, pkg)

	extractFileNode(tw, ast)

	tw.ForEachObject(extractObjectType)

	extractNumLines(tw, path, ast)

	end := time.Since(start)
	log.Printf("Done extracting %s (%dms)", path, end.Nanoseconds()/1000000)

	return nil
}

// stemAndExt splits a given file name into its stem (the part before the last '.')
// and extension (the part after the last '.')
func stemAndExt(base string) (string, string) {
	if i := strings.LastIndexByte(base, '.'); i >= 0 {
		return base[:i], base[i+1:]
	}
	return base, ""
}

// extractFileInfo extracts file-system level information for the given file, populating
// the `files` and `containerparent` tables
func extractFileInfo(tw *trap.Writer, file string) {
	path := filepath.ToSlash(srcarchive.TransformPath(file))
	components := strings.Split(path, "/")
	parentPath := ""
	var parentLbl trap.Label
	for i, component := range components {
		if i == 0 {
			if component == "" {
				path = "/"
			} else {
				path = component
			}
		} else {
			path = parentPath + "/" + component
		}
		if i == len(components)-1 {
			stem, ext := stemAndExt(component)
			lbl := tw.Labeler.FileLabel()
			dbscheme.FilesTable.Emit(tw, lbl, path, stem, ext, 0)
			dbscheme.ContainerParentTable.Emit(tw, parentLbl, lbl)
			extractLocation(tw, lbl, 0, 0, 0, 0)
			break
		}
		lbl := tw.Labeler.GlobalID(path + ";folder")
		dbscheme.FoldersTable.Emit(tw, lbl, path, component)
		if i > 0 {
			dbscheme.ContainerParentTable.Emit(tw, parentLbl, lbl)
		}
		if path != "/" {
			parentPath = path
		}
		parentLbl = lbl
	}
}

// extractLocation emits a location entity for the given entity
func extractLocation(tw *trap.Writer, entity trap.Label, sl int, sc int, el int, ec int) {
	lbl := tw.Labeler.FileLabel()
	locLbl := tw.Labeler.GlobalID(fmt.Sprintf("loc,{%s},%d,%d,%d,%d", lbl.String(), sl, sc, el, ec))
	dbscheme.LocationsDefaultTable.Emit(tw, locLbl, lbl, sl, sc, el, ec)
	dbscheme.HasLocationTable.Emit(tw, entity, locLbl)
}

// extractNodeLocation extracts location information for the given node
func extractNodeLocation(tw *trap.Writer, nd ast.Node, lbl trap.Label) {
	if nd == nil {
		return
	}
	fset := tw.Package.Fset
	start, end := fset.Position(nd.Pos()), fset.Position(nd.End())
	extractLocation(tw, lbl, start.Line, start.Column, end.Line, end.Column-1)
}

// extractPackageScope extracts symbol table information for the given package
func extractPackageScope(tw *trap.Writer, pkg *packages.Package) trap.Label {
	pkgScope := pkg.Types.Scope()
	pkgScopeLabel := tw.Labeler.ScopeID(pkgScope, pkg.Types)
	dbscheme.ScopesTable.Emit(tw, pkgScopeLabel, dbscheme.PackageScopeType.Index())
	dbscheme.ScopeNestingTable.Emit(tw, pkgScopeLabel, tw.Labeler.ScopeID(types.Universe, nil))
	extractObjects(tw, pkgScope, pkgScopeLabel)
	return pkgScopeLabel
}

// extractScopeLocation extracts location information for the given scope
func extractScopeLocation(tw *trap.Writer, scope *types.Scope, lbl trap.Label) {
	fset := tw.Package.Fset
	start, end := fset.Position(scope.Pos()), fset.Position(scope.End())
	extractLocation(tw, lbl, start.Line, start.Column, end.Line, end.Column-1)
}

// extractScopes extracts symbol table information for the package scope and all local scopes
// of the given package
func extractScopes(tw *trap.Writer, nd *ast.File, pkg *packages.Package) {
	pkgScopeLabel := extractPackageScope(tw, pkg)
	fileScope := pkg.TypesInfo.Scopes[nd]
	if fileScope != nil {
		extractLocalScope(tw, fileScope, pkgScopeLabel)
	}
}

// extractLocalScope extracts symbol table information for the given scope and all its nested scopes
func extractLocalScope(tw *trap.Writer, scope *types.Scope, parentScopeLabel trap.Label) {
	scopeLabel := tw.Labeler.ScopeID(scope, nil)
	dbscheme.ScopesTable.Emit(tw, scopeLabel, dbscheme.LocalScopeType.Index())
	extractScopeLocation(tw, scope, scopeLabel)
	dbscheme.ScopeNestingTable.Emit(tw, scopeLabel, parentScopeLabel)

	for i := 0; i < scope.NumChildren(); i++ {
		childScope := scope.Child(i)
		extractLocalScope(tw, childScope, scopeLabel)
	}

	extractObjects(tw, scope, scopeLabel)
}

// extractFileNode extracts AST information for the given file and all nodes contained in it
func extractFileNode(tw *trap.Writer, nd *ast.File) {
	lbl := tw.Labeler.FileLabel()

	extractExpr(tw, nd.Name, lbl, 0)

	for i, decl := range nd.Decls {
		extractDecl(tw, decl, lbl, i)
	}

	for _, cg := range nd.Comments {
		extractCommentGroup(tw, cg)
	}

	extractDoc(tw, nd.Doc, lbl)
	emitScopeNodeInfo(tw, nd, lbl)
}

// extractDoc extracts information about a doc comment group associated with a given element
func extractDoc(tw *trap.Writer, doc *ast.CommentGroup, elt trap.Label) {
	if doc != nil {
		dbscheme.DocCommentsTable.Emit(tw, elt, tw.Labeler.LocalID(doc))
	}
}

// extractCommentGroup extracts information about a doc comment group
func extractCommentGroup(tw *trap.Writer, cg *ast.CommentGroup) {
	lbl := tw.Labeler.LocalID(cg)
	dbscheme.CommentGroupsTable.Emit(tw, lbl)
	extractNodeLocation(tw, cg, lbl)
	for i, c := range cg.List {
		extractComment(tw, c, lbl, i)
	}
}

// extractComment extracts information about a given comment
func extractComment(tw *trap.Writer, c *ast.Comment, parent trap.Label, idx int) {
	lbl := tw.Labeler.LocalID(c)
	rawText := c.Text
	var kind int
	var text string
	if rawText[:2] == "//" {
		kind = dbscheme.SlashSlashComment.Index()
		text = rawText[2:]
	} else {
		kind = dbscheme.SlashStarComment.Index()
		text = rawText[2 : len(rawText)-2]
	}
	dbscheme.CommentsTable.Emit(tw, lbl, kind, parent, idx, text)
	extractNodeLocation(tw, c, lbl)
}

// emitScopeNodeInfo associates an AST node with its induced scope, if any
func emitScopeNodeInfo(tw *trap.Writer, nd ast.Node, lbl trap.Label) {
	scope, exists := tw.Package.TypesInfo.Scopes[nd]
	if exists {
		dbscheme.ScopeNodesTable.Emit(tw, lbl, tw.Labeler.ScopeID(scope, tw.Package.Types))
	}
}

// extractExpr extracts AST information for the given expression and all its subexpressions
func extractExpr(tw *trap.Writer, expr ast.Expr, parent trap.Label, idx int) {
	if expr == nil {
		return
	}

	lbl := tw.Labeler.LocalID(expr)
	extractTypeOf(tw, expr, lbl)

	var kind int
	switch expr := expr.(type) {
	case *ast.BadExpr:
		kind = dbscheme.BadExpr.Index()
	case *ast.Ident:
		if expr == nil {
			return
		}
		kind = dbscheme.IdentExpr.Index()
		dbscheme.LiteralsTable.Emit(tw, lbl, expr.Name, expr.Name)
		def := tw.Package.TypesInfo.Defs[expr]
		if def != nil {
			defTyp := extractType(tw, def.Type())
			objlbl, exists := tw.Labeler.LookupObjectID(def, defTyp)
			if objlbl == trap.InvalidLabel {
				log.Printf("Omitting def binding to unknown object %v", def)
			} else {
				if !exists {
					extractObject(tw, def, objlbl)
				}
				dbscheme.DefsTable.Emit(tw, lbl, objlbl)
			}
		}
		use := tw.Package.TypesInfo.Uses[expr]
		if use != nil {
			useTyp := extractType(tw, use.Type())
			objlbl, exists := tw.Labeler.LookupObjectID(use, useTyp)
			if objlbl == trap.InvalidLabel {
				log.Printf("Omitting use binding to unknown object %v", use)
			} else {
				if !exists {
					extractObject(tw, use, objlbl)
				}
				dbscheme.UsesTable.Emit(tw, lbl, objlbl)
			}
		}
	case *ast.Ellipsis:
		if expr == nil {
			return
		}
		kind = dbscheme.EllipsisExpr.Index()
		extractExpr(tw, expr.Elt, lbl, 0)
	case *ast.BasicLit:
		if expr == nil {
			return
		}
		value := ""
		switch expr.Kind {
		case token.INT:
			ival, _ := strconv.ParseInt(expr.Value, 0, 64)
			value = strconv.FormatInt(ival, 10)
			kind = dbscheme.IntLitExpr.Index()
		case token.FLOAT:
			value = expr.Value
			kind = dbscheme.FloatLitExpr.Index()
		case token.IMAG:
			value = expr.Value
			kind = dbscheme.ImagLitExpr.Index()
		case token.CHAR:
			value, _ = strconv.Unquote(expr.Value)
			kind = dbscheme.CharLitExpr.Index()
		case token.STRING:
			value, _ = strconv.Unquote(expr.Value)
			kind = dbscheme.StringLitExpr.Index()
		default:
			log.Fatalf("unknown literal kind %v", expr.Kind)
		}
		dbscheme.LiteralsTable.Emit(tw, lbl, value, expr.Value)
	case *ast.FuncLit:
		if expr == nil {
			return
		}
		kind = dbscheme.FuncLitExpr.Index()
		extractExpr(tw, expr.Type, lbl, 0)
		extractStmt(tw, expr.Body, lbl, 1)
	case *ast.CompositeLit:
		if expr == nil {
			return
		}
		kind = dbscheme.CompositeLitExpr.Index()
		extractExpr(tw, expr.Type, lbl, 0)
		extractExprs(tw, expr.Elts, lbl, 1, 1)
	case *ast.ParenExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.ParenExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
	case *ast.SelectorExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.SelectorExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
		extractExpr(tw, expr.Sel, lbl, 1)
	case *ast.IndexExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.IndexExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
		extractExpr(tw, expr.Index, lbl, 1)
	case *ast.SliceExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.SliceExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
		extractExpr(tw, expr.Low, lbl, 1)
		extractExpr(tw, expr.High, lbl, 2)
		extractExpr(tw, expr.Max, lbl, 3)
	case *ast.TypeAssertExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.TypeAssertExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
		extractExpr(tw, expr.Type, lbl, 1)
	case *ast.CallExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.CallOrConversionExpr.Index()
		extractExpr(tw, expr.Fun, lbl, 0)
		extractExprs(tw, expr.Args, lbl, 1, 1)
	case *ast.StarExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.StarExpr.Index()
		extractExpr(tw, expr.X, lbl, 0)
	case *ast.KeyValueExpr:
		if expr == nil {
			return
		}
		kind = dbscheme.KeyValueExpr.Index()
		extractExpr(tw, expr.Key, lbl, 0)
		extractExpr(tw, expr.Value, lbl, 1)
	case *ast.UnaryExpr:
		if expr == nil {
			return
		}
		tp := dbscheme.UnaryExprs[expr.Op]
		if tp == nil {
			log.Fatalf("unsupported unary operator %s", expr.Op)
		}
		kind = tp.Index()
		extractExpr(tw, expr.X, lbl, 0)
	case *ast.BinaryExpr:
		if expr == nil {
			return
		}
		tp := dbscheme.BinaryExprs[expr.Op]
		if tp == nil {
			log.Fatalf("unsupported binary operator %s", expr.Op)
		}
		kind = tp.Index()
		extractExpr(tw, expr.X, lbl, 0)
		extractExpr(tw, expr.Y, lbl, 1)
	case *ast.ArrayType:
		if expr == nil {
			return
		}
		kind = dbscheme.ArrayTypeExpr.Index()
		extractExpr(tw, expr.Len, lbl, 0)
		extractExpr(tw, expr.Elt, lbl, 1)
	case *ast.StructType:
		if expr == nil {
			return
		}
		kind = dbscheme.StructTypeExpr.Index()
		extractFields(tw, expr.Fields, lbl, 0, 1)
	case *ast.FuncType:
		if expr == nil {
			return
		}
		kind = dbscheme.FuncTypeExpr.Index()
		extractFields(tw, expr.Params, lbl, 0, 1)
		extractFields(tw, expr.Results, lbl, -1, -1)
		emitScopeNodeInfo(tw, expr, lbl)
	case *ast.InterfaceType:
		if expr == nil {
			return
		}
		kind = dbscheme.InterfaceTypeExpr.Index()
		extractFields(tw, expr.Methods, lbl, 0, 1)
	case *ast.MapType:
		if expr == nil {
			return
		}
		kind = dbscheme.MapTypeExpr.Index()
		extractExpr(tw, expr.Key, lbl, 0)
		extractExpr(tw, expr.Value, lbl, 1)
	case *ast.ChanType:
		if expr == nil {
			return
		}
		tp := dbscheme.ChanTypeExprs[expr.Dir]
		if tp == nil {
			log.Fatalf("unsupported channel direction %v", expr.Dir)
		}
		kind = tp.Index()
		extractExpr(tw, expr.Value, lbl, 0)
	default:
		log.Fatalf("unknown expression of type %T", expr)
	}
	dbscheme.ExprsTable.Emit(tw, lbl, kind, parent, idx)
	extractNodeLocation(tw, expr, lbl)
	extractValueOf(tw, expr, lbl)
}

// extractExprs extracts AST information for a list of expressions, which are children of
// the given parent
// `idx` is the index of the first child in the list, and `dir` is the index increment of
// each child over its preceding child (usually either 1 for assigning increasing indices, or
// -1 for decreasing indices)
func extractExprs(tw *trap.Writer, exprs []ast.Expr, parent trap.Label, idx int, dir int) {
	if exprs != nil {
		for _, expr := range exprs {
			extractExpr(tw, expr, parent, idx)
			idx += dir
		}
	}
}

// extractTypeOf looks up the type of `expr`, extracts it if it hasn't previously been
// extracted, and associates it with `expr` in the `type_of` table
func extractTypeOf(tw *trap.Writer, expr ast.Expr, lbl trap.Label) {
	tp := tw.Package.TypesInfo.TypeOf(expr)
	if tp != nil {
		tplbl := extractType(tw, tp)
		dbscheme.TypeOfTable.Emit(tw, lbl, tplbl)
	}
}

// extractValueOf looks up the value of `expr`, and associates it with `expr` in
// the `consts` table
func extractValueOf(tw *trap.Writer, expr ast.Expr, lbl trap.Label) {
	tpVal := tw.Package.TypesInfo.Types[expr]

	if tpVal.Value != nil {
		// if Value is non-nil, the expression has a constant value

		// note that string literals in import statements do not have an associated
		// Value and so do not get added to the table

		var value string
		exact := tpVal.Value.ExactString()
		switch tpVal.Value.Kind() {
		case constant.String:
			// we need to unquote strings
			value = constant.StringVal(tpVal.Value)
			exact = constant.StringVal(tpVal.Value)
		case constant.Float:
			flval, _ := constant.Float64Val(tpVal.Value)
			value = fmt.Sprintf("%.20g", flval)
		case constant.Complex:
			real, _ := constant.Float64Val(constant.Real(tpVal.Value))
			imag, _ := constant.Float64Val(constant.Imag(tpVal.Value))
			value = fmt.Sprintf("(%.20g + %.20gi)", real, imag)
		default:
			value = tpVal.Value.ExactString()
		}

		dbscheme.ConstValuesTable.Emit(tw, lbl, value, exact)
	} else if tpVal.IsNil() {
		dbscheme.ConstValuesTable.Emit(tw, lbl, "nil", "nil")
	}
}

// extractFields extracts AST information for a list of fields, which are children of
// the given parent
// `idx` is the index of the first child in the list, and `dir` is the index increment of
// each child over its preceding child (usually either 1 for assigning increasing indices, or
// -1 for decreasing indices)
func extractFields(tw *trap.Writer, fields *ast.FieldList, parent trap.Label, idx int, dir int) {
	if fields == nil || fields.List == nil {
		return
	}
	for _, field := range fields.List {
		lbl := tw.Labeler.LocalID(field)
		dbscheme.FieldsTable.Emit(tw, lbl, parent, idx)
		extractNodeLocation(tw, field, lbl)
		if field.Names != nil {
			for i, name := range field.Names {
				extractExpr(tw, name, lbl, i+1)
			}
		}
		extractExpr(tw, field.Type, lbl, 0)
		extractExpr(tw, field.Tag, lbl, -1)
		extractDoc(tw, field.Doc, lbl)
		idx += dir
	}
}

// extractStmt extracts AST information for a given statement and all other statements or expressions
// nested inside it
func extractStmt(tw *trap.Writer, stmt ast.Stmt, parent trap.Label, idx int) {
	if stmt == nil {
		return
	}

	lbl := tw.Labeler.LocalID(stmt)
	var kind int
	switch stmt := stmt.(type) {
	case *ast.BadStmt:
		kind = dbscheme.BadStmtType.Index()
	case *ast.DeclStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.DeclStmtType.Index()
		extractDecl(tw, stmt.Decl, lbl, 0)
	case *ast.EmptyStmt:
		kind = dbscheme.EmptyStmtType.Index()
	case *ast.LabeledStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.LabeledStmtType.Index()
		extractExpr(tw, stmt.Label, lbl, 0)
		extractStmt(tw, stmt.Stmt, lbl, 1)
	case *ast.ExprStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.ExprStmtType.Index()
		extractExpr(tw, stmt.X, lbl, 0)
	case *ast.SendStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.SendStmtType.Index()
		extractExpr(tw, stmt.Chan, lbl, 0)
		extractExpr(tw, stmt.Value, lbl, 1)
	case *ast.IncDecStmt:
		if stmt == nil {
			return
		}
		if stmt.Tok == token.INC {
			kind = dbscheme.IncStmtType.Index()
		} else if stmt.Tok == token.DEC {
			kind = dbscheme.DecStmtType.Index()
		} else {
			log.Fatalf("unsupported increment/decrement operator %v", stmt.Tok)
		}
		extractExpr(tw, stmt.X, lbl, 0)
	case *ast.AssignStmt:
		if stmt == nil {
			return
		}
		tp := dbscheme.AssignStmtTypes[stmt.Tok]
		if tp == nil {
			log.Fatalf("unsupported assignment statement with operator %v", stmt.Tok)
		}
		kind = tp.Index()
		extractExprs(tw, stmt.Lhs, lbl, -1, -1)
		extractExprs(tw, stmt.Rhs, lbl, 1, 1)
	case *ast.GoStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.GoStmtType.Index()
		extractExpr(tw, stmt.Call, lbl, 0)
	case *ast.DeferStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.DeferStmtType.Index()
		extractExpr(tw, stmt.Call, lbl, 0)
	case *ast.ReturnStmt:
		kind = dbscheme.ReturnStmtType.Index()
		extractExprs(tw, stmt.Results, lbl, 0, 1)
	case *ast.BranchStmt:
		if stmt == nil {
			return
		}
		switch stmt.Tok {
		case token.BREAK:
			kind = dbscheme.BreakStmtType.Index()
		case token.CONTINUE:
			kind = dbscheme.ContinueStmtType.Index()
		case token.GOTO:
			kind = dbscheme.GotoStmtType.Index()
		case token.FALLTHROUGH:
			kind = dbscheme.FallthroughStmtType.Index()
		default:
			log.Fatalf("unsupported branch statement type %v", stmt.Tok)
		}
		extractExpr(tw, stmt.Label, lbl, 0)
	case *ast.BlockStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.BlockStmtType.Index()
		extractStmts(tw, stmt.List, lbl, 0, 1)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.IfStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.IfStmtType.Index()
		extractStmt(tw, stmt.Init, lbl, 0)
		extractExpr(tw, stmt.Cond, lbl, 1)
		extractStmt(tw, stmt.Body, lbl, 2)
		extractStmt(tw, stmt.Else, lbl, 3)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.CaseClause:
		if stmt == nil {
			return
		}
		kind = dbscheme.CaseClauseType.Index()
		extractExprs(tw, stmt.List, lbl, -1, -1)
		extractStmts(tw, stmt.Body, lbl, 0, 1)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.SwitchStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.ExprSwitchStmtType.Index()
		extractStmt(tw, stmt.Init, lbl, 0)
		extractExpr(tw, stmt.Tag, lbl, 1)
		extractStmt(tw, stmt.Body, lbl, 2)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.TypeSwitchStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.TypeSwitchStmtType.Index()
		extractStmt(tw, stmt.Init, lbl, 0)
		extractStmt(tw, stmt.Assign, lbl, 1)
		extractStmt(tw, stmt.Body, lbl, 2)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.CommClause:
		if stmt == nil {
			return
		}
		kind = dbscheme.CommClauseType.Index()
		extractStmt(tw, stmt.Comm, lbl, 0)
		extractStmts(tw, stmt.Body, lbl, 1, 1)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.SelectStmt:
		kind = dbscheme.SelectStmtType.Index()
		extractStmt(tw, stmt.Body, lbl, 0)
	case *ast.ForStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.ForStmtType.Index()
		extractStmt(tw, stmt.Init, lbl, 0)
		extractExpr(tw, stmt.Cond, lbl, 1)
		extractStmt(tw, stmt.Post, lbl, 2)
		extractStmt(tw, stmt.Body, lbl, 3)
		emitScopeNodeInfo(tw, stmt, lbl)
	case *ast.RangeStmt:
		if stmt == nil {
			return
		}
		kind = dbscheme.RangeStmtType.Index()
		extractExpr(tw, stmt.Key, lbl, 0)
		extractExpr(tw, stmt.Value, lbl, 1)
		extractExpr(tw, stmt.X, lbl, 2)
		extractStmt(tw, stmt.Body, lbl, 3)
		emitScopeNodeInfo(tw, stmt, lbl)
	default:
		log.Fatalf("unknown statement of type %T", stmt)
	}
	dbscheme.StmtsTable.Emit(tw, lbl, kind, parent, idx)
	extractNodeLocation(tw, stmt, lbl)
}

// extractStmts extracts AST information for a list of statements, which are children of
// the given parent
// `idx` is the index of the first child in the list, and `dir` is the index increment of
// each child over its preceding child (usually either 1 for assigning increasing indices, or
// -1 for decreasing indices)
func extractStmts(tw *trap.Writer, stmts []ast.Stmt, parent trap.Label, idx int, dir int) {
	if stmts != nil {
		for _, stmt := range stmts {
			extractStmt(tw, stmt, parent, idx)
			idx += dir
		}
	}

}

// extractDecl extracts AST information for the given declaration
func extractDecl(tw *trap.Writer, decl ast.Decl, parent trap.Label, idx int) {
	lbl := tw.Labeler.LocalID(decl)
	var kind int
	switch decl := decl.(type) {
	case *ast.BadDecl:
		kind = dbscheme.BadDeclType.Index()
	case *ast.GenDecl:
		if decl == nil {
			return
		}
		switch decl.Tok {
		case token.IMPORT:
			kind = dbscheme.ImportDeclType.Index()
		case token.CONST:
			kind = dbscheme.ConstDeclType.Index()
		case token.TYPE:
			kind = dbscheme.TypeDeclType.Index()
		case token.VAR:
			kind = dbscheme.VarDeclType.Index()
		default:
			log.Fatalf("unknown declaration of kind %v", decl.Tok)
		}
		for i, spec := range decl.Specs {
			extractSpec(tw, spec, lbl, i)
		}
		extractDoc(tw, decl.Doc, lbl)
	case *ast.FuncDecl:
		if decl == nil {
			return
		}
		kind = dbscheme.FuncDeclType.Index()
		extractFields(tw, decl.Recv, lbl, -1, -1)
		extractExpr(tw, decl.Name, lbl, 0)
		extractExpr(tw, decl.Type, lbl, 1)
		extractStmt(tw, decl.Body, lbl, 2)
		extractDoc(tw, decl.Doc, lbl)
	default:
		log.Fatalf("unknown declaration of type %T", decl)
	}
	dbscheme.DeclsTable.Emit(tw, lbl, kind, parent, idx)
	extractNodeLocation(tw, decl, lbl)
}

// extractSpec extracts AST information for the given declaration specifier
func extractSpec(tw *trap.Writer, spec ast.Spec, parent trap.Label, idx int) {
	lbl := tw.Labeler.LocalID(spec)
	var kind int
	switch spec := spec.(type) {
	case *ast.ImportSpec:
		if spec == nil {
			return
		}
		kind = dbscheme.ImportSpecType.Index()
		extractExpr(tw, spec.Name, lbl, 0)
		extractExpr(tw, spec.Path, lbl, 1)
		extractDoc(tw, spec.Doc, lbl)
	case *ast.ValueSpec:
		if spec == nil {
			return
		}
		kind = dbscheme.ValueSpecType.Index()
		for i, name := range spec.Names {
			extractExpr(tw, name, lbl, -(1 + i))
		}
		extractExpr(tw, spec.Type, lbl, 0)
		extractExprs(tw, spec.Values, lbl, 1, 1)
		extractDoc(tw, spec.Doc, lbl)
	case *ast.TypeSpec:
		if spec == nil {
			return
		}
		kind = dbscheme.TypeSpecType.Index()
		extractExpr(tw, spec.Name, lbl, 0)
		extractExpr(tw, spec.Type, lbl, 1)
		extractDoc(tw, spec.Doc, lbl)
	}
	dbscheme.SpecsTable.Emit(tw, lbl, kind, parent, idx)
	extractNodeLocation(tw, spec, lbl)
}

// extractType extracts type information for `tp` and returns its associated label;
// types are only extracted once, so the second time `extractType` is invoked it simply returns the label
func extractType(tw *trap.Writer, tp types.Type) trap.Label {
	lbl, exists := getTypeLabel(tw, tp)
	if !exists {
		var kind int
		switch tp := tp.(type) {
		case *types.Basic:
			branch := dbscheme.BasicTypes[tp.Kind()]
			if branch == nil {
				log.Fatalf("unknown basic type %v", tp.Kind())
			}
			kind = branch.Index()
		case *types.Array:
			kind = dbscheme.ArrayType.Index()
			dbscheme.ArrayLengthTable.Emit(tw, lbl, fmt.Sprintf("%d", tp.Len()))
			extractElementType(tw, lbl, tp.Elem())
		case *types.Slice:
			kind = dbscheme.SliceType.Index()
			extractElementType(tw, lbl, tp.Elem())
		case *types.Struct:
			kind = dbscheme.StructType.Index()
			for i := 0; i < tp.NumFields(); i++ {
				field := tp.Field(i)

				// ensure the field is associated with a label
				fieldlbl, exists := tw.Labeler.FieldID(field, i, lbl)
				if !exists {
					extractObject(tw, field, fieldlbl)
				}

				dbscheme.FieldStructsTable.Emit(tw, fieldlbl, lbl)

				name := field.Name()
				if field.Embedded() {
					name = ""
				}
				extractComponentType(tw, lbl, i, name, field.Type())
			}
		case *types.Pointer:
			kind = dbscheme.PointerType.Index()
			extractBaseType(tw, lbl, tp.Elem())
		case *types.Interface:
			kind = dbscheme.InterfaceType.Index()
			for i := 0; i < tp.NumMethods(); i++ {
				meth := tp.Method(i)

				// get the receiver type of component methods; for interfaces
				// this can be different from the type used to get the method
				recvTyp := tp.Method(i).Type().(*types.Signature).Recv().Type()
				recvlbl := extractType(tw, recvTyp) // ensure receiver type has a label
				// ensure the method is associated with a label
				methlbl, exists := tw.Labeler.MethodID(meth, recvlbl)
				if !exists {
					extractObject(tw, meth, methlbl)
				}

				extractComponentType(tw, lbl, i, meth.Name(), meth.Type())
			}
		case *types.Tuple:
			kind = dbscheme.TupleType.Index()
			for i := 0; i < tp.Len(); i++ {
				extractComponentType(tw, lbl, i, "", tp.At(i).Type())
			}
		case *types.Signature:
			kind = dbscheme.SignatureType.Index()
			parms, results := tp.Params(), tp.Results()
			if parms != nil {
				for i := 0; i < parms.Len(); i++ {
					parm := parms.At(i)
					extractComponentType(tw, lbl, i+1, "", parm.Type())
				}
			}
			if results != nil {
				for i := 0; i < results.Len(); i++ {
					result := results.At(i)
					extractComponentType(tw, lbl, -(i + 1), "", result.Type())
				}
			}
		case *types.Map:
			kind = dbscheme.MapType.Index()
			extractKeyType(tw, lbl, tp.Key())
			extractElementType(tw, lbl, tp.Elem())
		case *types.Chan:
			kind = dbscheme.ChanTypes[tp.Dir()].Index()
			extractElementType(tw, lbl, tp.Elem())
		case *types.Named:
			kind = dbscheme.NamedType.Index()
			dbscheme.TypeNameTable.Emit(tw, lbl, tp.Obj().Name())
			underlying := tp.Underlying()
			extractUnderlyingType(tw, lbl, underlying)

			entitylbl, exists := tw.Labeler.LookupObjectID(tp.Obj(), lbl)
			if entitylbl == trap.InvalidLabel {
				log.Printf("Omitting type-object binding for unknown object %v.\n", tp.Obj())
			} else {
				if !exists {
					extractObject(tw, tp.Obj(), entitylbl)
				}
				dbscheme.TypeObjectTable.Emit(tw, lbl, entitylbl)
			}

			// ensure all methods have labels
			for i := 0; i < tp.NumMethods(); i++ {
				meth := tp.Method(i)
				recv := meth.Type().(*types.Signature).Recv()
				typ := recv.Type()
				recvTypeLbl := extractType(tw, typ) // ensure receiver type has a label
				methlbl, exists := tw.Labeler.MethodID(tp.Method(i), recvTypeLbl)
				if !exists {
					extractObject(tw, meth, methlbl)
				}
			}

			// associate all methods of underlying interface with this type
			if underlyingInterface, ok := underlying.(*types.Interface); ok {
				underlyingInterfaceLabel, _ := getTypeLabel(tw, underlyingInterface)
				for i := 0; i < underlyingInterface.NumMethods(); i++ {
					meth := underlyingInterface.Method(i)
					methlbl, exists := tw.Labeler.MethodID(underlyingInterface.Method(i), underlyingInterfaceLabel)
					if !exists {
						log.Printf("No label for method %s of type %s yet.", meth.Name(), tp.Obj().Name())
					}
					dbscheme.MethodHostsTable.Emit(tw, methlbl, lbl)
				}
			}
		default:
			log.Fatalf("unexpected type %T", tp)
		}
		dbscheme.TypesTable.Emit(tw, lbl, kind)
	}
	return lbl
}

// getTypeLabel looks up the label associated with `tp`, creating a new label if
// it does not have one yet; the second result indicates whether the label
// already existed
//
// Type labels refer to global keys to ensure that if the same type is
// encountered during the extraction of different files it is still ultimately
// mapped to the same entity. In particular, this means that keys for compound
// types refer to the labels of their component types. For named types, the key
// is constructed from their globally unique ID. This prevents cyclic type keys
// since type recursion in Go always goes through named types.
func getTypeLabel(tw *trap.Writer, tp types.Type) (trap.Label, bool) {
	lbl, exists := tw.Labeler.TypeLabels[tp]
	if !exists {
		switch tp := tp.(type) {
		case *types.Basic:
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%d;basictype", tp.Kind()))
		case *types.Array:
			len := tp.Len()
			elem := extractType(tw, tp.Elem())
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%d,{%s};arraytype", len, elem))
		case *types.Slice:
			elem := extractType(tw, tp.Elem())
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("{%s};slicetype", elem))
		case *types.Struct:
			var b strings.Builder
			for i := 0; i < tp.NumFields(); i++ {
				field := tp.Field(i)
				fieldTypeLbl := extractType(tw, field.Type())
				if i > 0 {
					b.WriteString(",")
				}
				name := field.Name()
				if field.Embedded() {
					name = ""
				}
				fmt.Fprintf(&b, "%s,{%s},%s", name, fieldTypeLbl, tp.Tag(i))
			}
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%s;structtype", b.String()))
		case *types.Pointer:
			base := extractType(tw, tp.Elem())
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("{%s};pointertype", base))
		case *types.Interface:
			var b strings.Builder
			for i := 0; i < tp.NumMethods(); i++ {
				meth := tp.Method(i)
				methLbl := extractType(tw, meth.Type())
				if i > 0 {
					b.WriteString(",")
				}
				fmt.Fprintf(&b, "%s,{%s}", meth.Id(), methLbl)
			}
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%s;interfacetype", b.String()))
		case *types.Tuple:
			var b strings.Builder
			for i := 0; i < tp.Len(); i++ {
				compLbl := extractType(tw, tp.At(i).Type())
				if i > 0 {
					b.WriteString(",")
				}
				fmt.Fprintf(&b, "{%s}", compLbl)
			}
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%s;tupletype", b.String()))
		case *types.Signature:
			var b strings.Builder
			parms, results := tp.Params(), tp.Results()
			if parms != nil {
				for i := 0; i < parms.Len(); i++ {
					parmLbl := extractType(tw, parms.At(i).Type())
					if i > 0 {
						b.WriteString(",")
					}
					fmt.Fprintf(&b, "{%s}", parmLbl)
				}
			}
			b.WriteString(";")
			if results != nil {
				for i := 0; i < results.Len(); i++ {
					resultLbl := extractType(tw, results.At(i).Type())
					if i > 0 {
						b.WriteString(",")
					}
					fmt.Fprintf(&b, "{%s}", resultLbl)
				}
			}
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%s;signaturetype", b.String()))
		case *types.Map:
			key := extractType(tw, tp.Key())
			value := extractType(tw, tp.Elem())
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("{%s},{%s};maptype", key, value))
		case *types.Chan:
			dir := tp.Dir()
			elem := extractType(tw, tp.Elem())
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("%v,{%s};chantype", dir, elem))
		case *types.Named:
			entitylbl, exists := tw.Labeler.LookupObjectID(tp.Obj(), lbl)
			if entitylbl == trap.InvalidLabel {
				panic(fmt.Sprintf("Cannot construct label for named type %v (underlying object is %v).\n", tp, tp.Obj()))
			}
			if !exists {
				extractObject(tw, tp.Obj(), entitylbl)
			}
			lbl = tw.Labeler.GlobalID(fmt.Sprintf("{%s};namedtype", entitylbl))
		}
		tw.Labeler.TypeLabels[tp] = lbl
	}
	return lbl, exists
}

// extractKeyType extracts `key` as the key type of the map type `mp`
func extractKeyType(tw *trap.Writer, mp trap.Label, key types.Type) {
	dbscheme.KeyTypeTable.Emit(tw, mp, extractType(tw, key))
}

// extractElementType extracts `element` as the element type of the container type `container`
func extractElementType(tw *trap.Writer, container trap.Label, element types.Type) {
	dbscheme.ElementTypeTable.Emit(tw, container, extractType(tw, element))
}

// extractBaseType extracts `base` as the base type of the pointer type `ptr`
func extractBaseType(tw *trap.Writer, ptr trap.Label, base types.Type) {
	dbscheme.BaseTypeTable.Emit(tw, ptr, extractType(tw, base))
}

// extractUnderlyingType extracts `underlying` as the underlying type of the
// named type `named`
func extractUnderlyingType(tw *trap.Writer, named trap.Label, underlying types.Type) {
	dbscheme.UnderlyingTypeTable.Emit(tw, named, extractType(tw, underlying))
}

// extractComponentType extracts `component` as the `idx`th component type of `parent` with name `name`
func extractComponentType(tw *trap.Writer, parent trap.Label, idx int, name string, component types.Type) {
	dbscheme.ComponentTypesTable.Emit(tw, parent, idx, name, extractType(tw, component))
}

// extractNumLines extracts lines-of-code and lines-of-comments information for the
// given file
func extractNumLines(tw *trap.Writer, fileName string, ast *ast.File) {
	f := tw.Package.Fset.File(ast.Pos())

	lineCount := f.LineCount()

	// count lines of code by tokenizing
	linesOfCode := 0
	src, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to read file %s.", fileName)
	}
	var s scanner.Scanner
	lastCodeLine := -1
	s.Init(f, src, nil, 0)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		} else if tok != token.ILLEGAL {
			tkStartLine := f.Position(pos).Line
			tkEndLine := tkStartLine + strings.Count(lit, "\n")
			if tkEndLine > lastCodeLine {
				linesOfCode += tkEndLine - tkStartLine + 1
				lastCodeLine = tkEndLine
			}
		}
	}

	// count lines of comments by iterating over ast.Comments
	linesOfComments := 0
	for _, cg := range ast.Comments {
		for _, g := range cg.List {
			fset := tw.Package.Fset
			startPos, endPos := fset.Position(g.Pos()), fset.Position(g.End())
			linesOfComments += endPos.Line - startPos.Line + 1
		}
	}

	dbscheme.NumlinesTable.Emit(tw, tw.Labeler.FileLabel(), lineCount, linesOfCode, linesOfComments)
}

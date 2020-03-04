/**
 * Provides classes for modeling go.mod dependencies.
 */

import go
private import semmle.go.dependencies.DependencyCustomizations

/**
 * An abstract representation of a dependency.
 */
abstract class Dependency extends Locatable {
  /**
   * Holds if this dependency has package path `path` and version `v`.
   *
   * If the version cannot be determined, `v` is bound to the string
   * `"unknown"`.
   */
  abstract predicate info(string path, string v);

  /** Gets the package path of this dependency. */
  string getDepPath() { this.info(result, _) }

  /** Gets the version of this dependency. */
  string getDepVer() { this.info(_, result) }

  /**
   * An import of this dependency.
   */
  ImportSpec getAnImport() { result.getPath() = this.getDepPath() }
}

/**
 * A dependency from a go.mod file.
 */
class GoModDependency extends Dependency, GoModRequireLine {
  override predicate info(string path, string v) {
    this.replacementInfo(path, v)
    or
    not this.replacementInfo(_, _) and
    this.originalInfo(path, v)
  }

  /**
   * Holds if there is a replace line that replaces this dependency with a dependency to `path`,
   * version `v`.
   */
  predicate replacementInfo(string path, string v) {
    exists(GoModReplaceLine replace |
      replace.getFile() = this.getFile() and
      replace.getOriginalPath() = this.getPath()
    |
      path = replace.getReplacementPath() and
      (
        v = replace.getReplacementVer()
        or
        not exists(replace.getReplacementVer()) and
        v = "unknown"
      )
    )
  }

  /**
   * Get a version that was excluded for this dependency.
   */
  string getAnExcludedVer() {
    exists(GoModExcludeLine exclude |
      exclude.getFile() = this.getFile() and
      exclude.getPath() = this.getPath()
    |
      result = exclude.getVer()
    )
  }

  /**
   * Holds if this require line originally states dependency `path` had version `ver`.
   */
  predicate originalInfo(string path, string v) { path = this.getPath() and v = this.getVer() }
}

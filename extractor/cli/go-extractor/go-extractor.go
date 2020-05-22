package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/github/codeql-go/extractor"
)

var cpuprofile, memprofile string

func usage() {
	fmt.Fprintf(os.Stderr, "%s is a program for building a snapshot of a Go code base.\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [<flag>...] [<buildflag>...] [--] <file>...\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n\n")
	fmt.Fprintf(os.Stderr, "--help                Print this help.\n")
}

func parseFlags(args []string) ([]string, []string) {
	i := 0
	mimic := false
	var dumpDbscheme string
	buildFlags := []string{}
	for i < len(args) && strings.HasPrefix(args[i], "-") {
		if args[i] == "--" {
			i++
			break
		}

		if !mimic && strings.HasPrefix(args[i], "--dbscheme=") {
			dumpDbscheme = strings.TrimPrefix(args[i], "--dbscheme=")
		} else if !mimic && args[i] == "--dbscheme" {
			i++
			dumpDbscheme = args[i]
		} else if !mimic && args[i] == "--help" {
			usage()
			os.Exit(0)
		} else if args[i] == "--mimic" {
			mimic = true
			if i+1 < len(args) {
				i++
				compiler := args[i]
				log.Printf("Compiler: %s", compiler)
				if i+1 < len(args) {
					i++
					command := args[i]
					if command == "build" {
						log.Printf("Intercepting build")
						// skip `-o output` and `i`, if applicable
						for i+1 < len(args) {
							if args[i+1] == "-o" {
								i++
								if i+1 < len(args) {
									i++
								}
							} else if args[i+1] == "-i" {
								i++
							} else {
								break
							}
						}
					} else {
						log.Printf("Non-build command; skipping")
						return []string{}, []string{}, ""
					}
				}
			} else {
				log.Fatalf("Invalid --mimic: no compiler specified")
			}
		} else {
			buildFlags = append(buildFlags, args[i])
		}

		i++
	}

	cpuprofile = os.Getenv("CODEQL_EXTRACTOR_GO_CPU_PROFILE")
	memprofile = os.Getenv("CODEQL_EXTRACTOR_GO_MEM_PROFILE")

	return buildFlags, args[i:]
}

func main() {
	buildFlags, patterns := parseFlags(os.Args[1:])

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatalf("Unable to create CPU profile: %v.", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("Unable to start CPU profile: %v.", err)
		}
		defer pprof.StopCPUProfile()
	}

	if len(patterns) == 0 {
		log.Println("Nothing to extract.")
	} else {
		log.Printf("Build flags: %s; patterns: %s\n", strings.Join(buildFlags, " "), strings.Join(patterns, " "))
		err := extractor.ExtractWithFlags(buildFlags, patterns)
		if err != nil {
			log.Fatalf("Error running go tooling: %s\n", err.Error())
		}
	}

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatalf("Unable to create memory profile: %v", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Unable to write memory profile: ", err)
		}
	}
}

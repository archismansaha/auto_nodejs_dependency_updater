package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

import "dependency_updater/colors"
import "dependency_updater/dependencies"

func main() {
	// Define command-line flags
	ignoreDeps := flag.Bool("ignore-deps", false, "Ignore updating dependencies in package.json defined as {dependencies}")
	ignoreDependencies := flag.String("ignore-dependencies", "", "Ignore updating dependencies comma(,) seperated")
	ignoreDevDeps := flag.Bool("ignore-dev-deps", false, "Ignore updating devDependencies in package.json defined as {devDependencies}")
	projectPath := flag.String("path", ".", "Path of the project")
 //fmt.Printf(`Project path %s %s %s`, projectPath,ignoreDeps,ignoreDevDeps,ignoreDependencies)
	// Define usage message
	flag.Usage = func() {
		colors.PrintUsage()
		flag.PrintDefaults()
		colors.PrintExample()
	}

	// Parse command-line flags
	flag.Parse()

	// Check if no flags are provided
	if flag.NFlag() == 0 {
		ignoreDeps := false
		ignoreDependencies:=[]string{}
		ignoreDevDeps := false
		projectPath := "."

		// Call update function with default values
		if err := dependencies.Update(projectPath, ignoreDeps, ignoreDevDeps,ignoreDependencies); err != nil {
			colors.PrintError("Error updating dependencies:", err.Error())
			os.Exit(1)
		}
		return
	}
ignoreDependency:=strings.Split(*ignoreDependencies, ",")
fmt.Printf(`Project path %s %s %s`, projectPath,ignoreDeps,ignoreDevDeps,ignoreDependency)
	// Check for conflicting flags
	if *ignoreDeps && *ignoreDevDeps {
		colors.PrintError("Error: Cannot use both --ignore-dependencies and --ignore-devDependencies flags at the same time.")
		flag.Usage()
		os.Exit(1)
	}

	// Call update function with provided values
	if err := dependencies.Update(*projectPath, *ignoreDeps, *ignoreDevDeps,ignoreDependency); err != nil {
		colors.PrintError("Error updating dependencies:", err.Error())
		os.Exit(1)
	}
}

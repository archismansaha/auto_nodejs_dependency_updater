package main

import (
	"flag"
	"fmt"
	"os"
)

import "dependency_updater/colors"
import "dependency_updater/dependencies"

func main() {
	// Define command-line flags
	ignoreDeps := flag.Bool("ignore-deps", false, "Ignore updating dependencies")
	ignoreDevDeps := flag.Bool("ignore-dev-deps", false, "Ignore updating devDependencies")
	projectPath := flag.String("path", ".", "Path of the project")
 fmt.Printf("Project path", projectPath)
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
		ignoreDevDeps := false
		projectPath := "."

		// Call update function with default values
		if err := dependencies.Update(projectPath, ignoreDeps, ignoreDevDeps); err != nil {
			colors.PrintError("Error updating dependencies:", err.Error())
			os.Exit(1)
		}
		return
	}

	// Check for conflicting flags
	if *ignoreDeps && *ignoreDevDeps {
		colors.PrintError("Error: Cannot use both --ignore-dependencies and --ignore-devDependencies flags at the same time.")
		flag.Usage()
		os.Exit(1)
	}

	// Call update function with provided values
	if err := dependencies.Update(*projectPath, *ignoreDeps, *ignoreDevDeps); err != nil {
		colors.PrintError("Error updating dependencies:", err.Error())
		os.Exit(1)
	}
}

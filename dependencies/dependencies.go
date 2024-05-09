package dependencies

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
import "dependency_updater/colors"
// PackageJSON represents the structure of package.json file
type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// Update updates dependencies
func Update(projectPath string, ignoreDeps, ignoreDevDeps bool,ignoreDependencies []string) error {
	// Resolve absolute path
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return err
	}

	fmt.Println(colors.Cyan + "Please wait. Reading package.json..." + colors.Reset)

	// Check if package.json exists
	packageJSONPath := filepath.Join(absPath, "package.json")
	if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
		return fmt.Errorf("package.json not found in the specified directory")
	}

	// Open package.json file
	file, err := os.Open(packageJSONPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode package.json content
	var pkgJSON PackageJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pkgJSON); err != nil {
		return err
	}

	ignoredDependenciesCount := 0

	// Update dependencies
	if !ignoreDeps {
		for dep, version := range pkgJSON.Dependencies {
			if err := updateDependency(absPath, dep, version,ignoreDependencies,&ignoredDependenciesCount); err != nil {
				return err
			}
		}
	} else {
		ignoredDependenciesCount += len(pkgJSON.Dependencies)
	}

	// Update devDependencies
	if !ignoreDevDeps {
		for dep, version := range pkgJSON.DevDependencies {
			//&ignoredDependenciesCount passing address so that we can keep track of skipped dependencies and update that count
			if err := updateDependency(absPath, dep, version,ignoreDependencies,&ignoredDependenciesCount); err != nil {
				return err
			}
		}
	} else {
		ignoredDependenciesCount += len(pkgJSON.DevDependencies)
	}

	fmt.Printf(colors.Yellow+"%d dependencies ignored."+colors.Reset+"\n", ignoredDependenciesCount)
	fmt.Println(colors.Green + "Everything updated successfully." + colors.Reset)
	return nil
}

func updateDependency(absPath, dep, version string,ignoreDependencies []string,ignoredDependenciesCount *int) error {
	//check if the dependency is provided to skip updating
	found := false
	for _, element := range ignoreDependencies{
		if element == dep {
			found = true
			*ignoredDependenciesCount+=1
			break
		}
	}
	if found {
		fmt.Printf("%s ignoring the dependecy\n", colors.Yellow+dep+colors.Reset)
		return nil
	}
	// Getting dependencies latest version
	cmd := exec.Command("npm", "show", dep, "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to check version for %s: %v", dep, err)
	}
       
	currentVersion := strings.TrimSpace(string(output))

	//comparing latest version with available version in package.json
	if currentVersion == version {
		fmt.Printf("%s is already up to date (current version: %s%s%s)\n", colors.Yellow+dep+colors.Reset, colors.Green, version, colors.Reset)
		return nil
	}

	fmt.Printf("Updating %s to %s%s%s (old version: %s%s%s)...\n", colors.Yellow+dep+colors.Reset, colors.Green, version, colors.Reset, colors.Red, currentVersion, colors.Reset)
	//updating the the dependency
	cmd = exec.Command("npm", "install", dep+"@"+version)
	cmd.Dir = absPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

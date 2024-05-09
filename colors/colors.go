package colors

import ( "fmt"
"os")

const (
	Reset  = "\033[0m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Red    = "\033[31m"
)

// PrintUsage prints usage message
func PrintUsage() {
	fmt.Printf("%sUsage:%s %s[options]%s\n\n", Cyan, Reset, Yellow, Reset)
	fmt.Println("Options:")
}

// PrintExample prints usage example
func PrintExample() {
	fmt.Printf("Example:\n  %s --path %s/path/to/project %s--ignore-dev-deps%s\n", os.Args[0], Yellow, Yellow, Reset)
}

// PrintError prints error message
func PrintError(message string,err ...string ) {
	errorMessage := "Unknown error"
    if len(err) > 0 {
        errorMessage = err[0]
    }
	fmt.Println(Red + message + errorMessage + Reset)
}

package main

import (
	"github.com/spf13/pflag"
)

// Global options
var (
	// NOTE: I don't know if prefixing them with an _ is a good idea

	_hideContext = pflag.BoolP("hide-context", "@", false, "Hide context names in list output.  Use twice to show context names (default).") // TODO: Implement "twice"
	_hideProject = pflag.BoolP("hide-project", "+", false, "Hide project names in list output.  Use twice to show project names (default).") // TODO: Implement "twice"
	_colorMode = pflag.StringP("color-mode", "c", "TODO", "Color mode")
	_configFile = pflag.StringP("config-file", "d", "", "Use a configuration file other than the default ~/.todo/config")
	_force = pflag.BoolP("force", "f", false, "Forces actions without confirmation or interactive input")
	_showHelp = pflag.BoolP("help", "h", false, "Display a short help message; same as action \"shorthelp\"")
	_plain = pflag.BoolP("plain", "p", false, "Plain mode turns off colors")
	_hidePriority = pflag.BoolP("hide-priority", "P", false, "Hide priority labels in list output.  Use twice to show priority labels (default).") // TODO: Implement "twice"
	_dontAutoArchive = pflag.BoolP("dont-auto-archive", "a", false, "Don't auto-archive tasks automatically on completion") // TODO: think of default
	_autoArchive = pflag.BoolP("auto-archive", "A", false, "Auto-archive tasks automatically on completion") // TODO: think of default
	_dontPreserveLineNumbers = pflag.BoolP("dont-preserve-line-numbers", "n", false, "Don't preserve line numbers; automatically remove blank lines on task deletion") // TODO: think of default
	_preserveLineNumbers = pflag.BoolP("preserve-line-numbers", "N", false, "Preserve line numbers") // TODO: think of default
	_prependDate = pflag.BoolP("prepend-date", "t", false, "Prepend the current date to a task automatically when it's added.")
	_dontPrependDate = pflag.BoolP("dont-prepend-date", "T", false, "Do not prepend the current date to a task automatically when it's added.")
	_verbose = pflag.BoolP("verbose", "v", false, "Verbose mode turns on confirmation messages")
	/*
	-vv
		Extra verbose mode prints some debugging information and additional help text
	*/
	_showVersion = pflag.BoolP("version", "V", false, "Displays version, license and credits")
	_disableFinalFilter = pflag.BoolP("disable-final-filter", "x", false, "Disables TODOTXT_FINAL_FILTER")
)

func main() {
	pflag.CommandLine.SortFlags = false
	pflag.Parse()
}

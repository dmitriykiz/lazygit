package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jesseduffield/lazygit/pkg/app"
	"github.com/jesseduffield/lazygit/pkg/app/daemon"
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/env"
	"github.com/jesseduffield/lazygit/pkg/integration/components"
	"github.com/jesseduffield/lazygit/pkg/logs"
)

var (
	// Version is the current version of lazygit, set at build time via ldflags
	Version = "unversioned"
	// BuildDate is the date the binary was built, set at build time via ldflags
	BuildDate = "unknown"
	// Commit is the git commit hash, set at build time via ldflags
	Commit = "unknown"
)

func main() {
	// If we're running as a daemon (e.g. for background git operations),
	// handle that separately and exit early.
	if daemon.InDaemonMode() {
		if err := daemon.Handle(); err != nil {
			log.Fatalf("daemon error: %v", err)
		}
		return
	}

	// Build version info from build-time variables
	buildInfo := &config.BuildInfo{
		Version:   Version,
		BuildDate: BuildDate,
		Commit:    Commit,
	}

	// Set up application-level logging before anything else
	logger := logs.Global

	// Allow integration tests to inject a custom test runner
	if components.IsIntegrationTest() {
		components.RunIntegrationTest()
		return
	}

	// Detect if we're running inside a git repo; warn if not.
	// Note: exit with a non-zero status code so shell scripts can detect this
	// and avoid launching the TUI in a meaningless state.
	// Personal note: I prefer a friendlier hint here so I remember to `cd` into
	// the right directory rather than just seeing a bare error.
	if !env.IsInsideWorkTree() {
		fmt.Fprintln(os.Stderr, "error: not inside a git repository (try cd-ing into your project directory)")
		fmt.Fprintln(os.Stderr, "tip: run `git init` if you want to start a new repository here")
		// Also print the current working directory to make it easier to see where we are
		if cwd, err := os.Getwd(); err == nil {
			fmt.Fprintf(os.Stderr, "current directory: %s\n", cwd)
		}
		os.Exit(1)
	}

	// Create and start the lazygit application
	lazyApp, err := app.NewApp(buildInfo, logger)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	if err := lazyApp.Run(); err != nil {
		// Personal preference: print the error to stderr as well so it's visible
		// even if the terminal is in a weird state after the TUI exits.
		fmt.Fprintf(os.Stderr, "lazygit exited with error: %v\n", err)
		// Use os.Exit instead of log.Fatalf to avoid printing a redundant
		// timestamp prefix on the error we already printed above.
		os.Exit(1)
	}

	// Personal addition: only print the goodbye message when stdout is a terminal
	// (i.e. not when piped/redirected), so it doesn't pollute script output.
	if fi, err := os.Stderr.Stat(); err == nil && (fi.Mode()&os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "lazygit exited cleanly")
	}
}

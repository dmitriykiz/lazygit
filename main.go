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
		os.Exit(1)
	}

	// Create and start the lazygit application
	lazyApp, err := app.NewApp(buildInfo, logger)
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	if err := lazyApp.Run(); err != nil {
		log.Fatalf("app exited with error: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	"github.com/aita/kannri"
)

var (
	help  bool
	files []string
)

func init() {
	pflag.BoolVarP(&help, "help", "h", false, "show this help message")
	pflag.StringArrayVarP(&files, "files", "f", []string{"Kannrifile.yaml"}, "kannri configuration files")
}

func showUsage() {
	progName := filepath.Base(os.Args[0])

	fmt.Fprintf(os.Stderr, `%s is a lightweight process manager.

Usage:

	%s [options]

Options:

%s
	`,
		progName,
		progName,
		pflag.CommandLine.FlagUsages(),
	)
}

func main() {
	ctx := context.Background()
	pflag.Parse()

	if help {
		showUsage()
		os.Exit(2)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	bootstrap, err := kannri.LoadBootstraps(files...)
	if err != nil {
		slog.ErrorContext(ctx, "failed to load bootstraps", "error", err)
		os.Exit(1)
	}

	manager := kannri.New()
	if err := kannri.ApplyBootstrap(ctx, manager, bootstrap); err != nil {
		slog.ErrorContext(ctx, "failed to apply bootstraps", "error", err)
		os.Exit(1)
	}

	if err := manager.Run(ctx); err != nil {
		slog.ErrorContext(ctx, "an error occurred while starting kannri", "error", err)
		os.Exit(1)
	}
}

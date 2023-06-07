// Package cli provides functions relating to running foo as a command-line interface.
package cli

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/foo/internal/logformatter"
)

var configFile string
var dryRun bool
var quiet bool
var verbose bool
var vertrace bool

var log *logrus.Entry

func init() {
	cobra.OnInitialize(initLogger)
	cobra.OnInitialize(initConfig)
}

func initLogger() {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logformatter.PrefixedTextFormatter{
		UseColor: true,
	})
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	}
	if vertrace {
		logger.SetLevel(logrus.TraceLevel)
	}
	if quiet {
		logger.SetLevel(logrus.ErrorLevel)
		logger.SetOutput(os.Stderr)
	}
	log = logger.WithField("prefix", "main")
}

// NewFooCmd returns a new CLI command representing Foo.
func NewFooCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "foo",
		Short: "A build orchestrator for terraform monorepos",
	}

	cmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "don't actually execute the task, just print it out")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "increase log output")
	cmd.PersistentFlags().BoolVar(&vertrace, "vvv", false, "increase log output even more")
	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "silences all logs but the errors (and prints those to stderr). Still prints command output to stdout. Overrides verbose and vvv")

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file for this command (optional, default is $HOME/.foo.yaml)")

	cmd.AddCommand(newVersionCommand())

	return cmd
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openshift/bootshift/pkg/util"
	"github.com/openshift/bootshift/pkg/version"
)

var (
	cmdRoot = &cobra.Command{
		Use:           "bootshift",
		Short:         "Bootshift!",
		SilenceErrors: true, // suppress cobra errors so we can handle them (also applies to subcommands)
		Long:          "",
	}

	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Output version information",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Version: %s\n", version.Version)
			return nil
		},
	}
)

func main() {
	flag.Parse()
	util.InitLogs()
	defer util.FlushLogs()

	cmdRoot.AddCommand(cmdVersion)
	if err := cmdRoot.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

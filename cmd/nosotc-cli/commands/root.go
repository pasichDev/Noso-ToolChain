package commands

import (
	"fmt"
	"os"

	ver "github.com/pasichDev/nosotc/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: ver.Version,
	Use:     "nosotc-cli",
	Short:   "N-Tolchain - an application for interaction with the Noso blockchain",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.SetVersionTemplate(fmt.Sprintf("%s\n", ver.Title))
}

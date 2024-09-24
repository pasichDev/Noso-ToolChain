package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"

	ver "github.com/pasichDev/nosotc/version"
)

var (
	nosoFilePath   string
	nosoDataFolder string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: ver.Version,
	Use:     "nosotc-cli",
	Short:   "N-Tolchain - an application for interaction with the Noso blockchain",
}

var getDataFolderCmd = &cobra.Command{
	Use:   "get-nosodata",
	Short: "Checks the path to the NOSODATA blockchain folder",
	Run:   getConfig,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Set up global flags
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s\n", ver.Title))
	rootCmd.PersistentFlags().StringVarP(&nosoFilePath, "file", "f", "", "The path to the specified file you want to open (blk, summary, wallet, gvt).")
	rootCmd.PersistentFlags().BoolVarP(&exportTxt, "export", "e", false, "Export selected data file to TXT file")
	rootCmd.AddCommand(getDataFolderCmd)

}

func getConfig(cmd *cobra.Command, args []string) {
	initConfig()
}
func initConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("⚠️ Failed to get the working directory: %v", err)
	}

	cfgFile := filepath.Join(workingDir, "config.yaml")
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("❌ Failed to read the configuration file")
		return
	}

	projectDir := viper.GetString("nosodata_path")
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		if nosoFilePath == "" {
			fmt.Printf("❌ The directory is invalid: %v", projectDir)
			fmt.Printf("❌ To use nosotc-cli, set the blockchain data folder path in config.yaml or use the -f flag to specify a file to open (e.g., -f /path/to/file.psk). You can also set it using 'set-nosodata' command.")

			os.Exit(1)
		}
		nosoDataFolder = ""
	} else {
		fmt.Println("✅ The configuration file exists, and the directory is valid:", projectDir)
		nosoDataFolder = projectDir
	}
}

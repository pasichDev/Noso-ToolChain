package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setDataFolderCmd = &cobra.Command{
	Use:   "set-nosodata",
	Short: "Sets the path to the Noso blockchain folder (NOSODATA)",
	Run:   setConfig,
}

func init() {
	// Add a flag to accept the path
	setDataFolderCmd.Flags().StringVar(&nosoDataFolder, "path", "", "Path to the NOSODATA blockchain folder")
	setDataFolderCmd.MarkFlagRequired("path") // Make the flag required

	// Add commands to rootCmd
	rootCmd.AddCommand(setDataFolderCmd)
}

// setConfig creates a configuration file using Viper
func setConfig(cmd *cobra.Command, args []string) {
	// Check if the path was provided via the flag
	if nosoDataFolder == "" {
		log.Fatalf("❌ You must provide the folder path using the --path flag")
	}

	// Get the working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("⚠️ Failed to get the working directory: %v", err)
	}

	// Define the path to the configuration file
	cfgFile := filepath.Join(workingDir, "config.yaml")

	// Configure Viper
	viper.SetConfigFile(cfgFile)

	// If the configuration file doesn't exist, create it
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		viper.Set("project", "noso_project")
		viper.Set("nosodata_path", nosoDataFolder)

		// Save the configuration to a file
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			log.Fatalf("❌ Failed to write the configuration file: %v", err)
		}
		fmt.Println("✅ Configuration file successfully created with path to NOSODATA:", nosoDataFolder)
	} else {
		// If the configuration file already exists, update it
		viper.Set("nosodata_path", nosoDataFolder)

		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("❌ Failed to update the configuration file: %v", err)
		}
		fmt.Println("🔄 Path to NOSODATA updated:", nosoDataFolder)
	}
}

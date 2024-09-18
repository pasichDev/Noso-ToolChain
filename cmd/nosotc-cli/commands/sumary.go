package commands

import (
	"fmt"
	"github.com/pasichDev/nosotc/internal/app"
	"os"

	"github.com/spf13/cobra"
)

var (
	sumaryCmd = &cobra.Command{
		Use:   "sumary",
		Short: "Handler for the Sumary.psk file",
		Run:   runSumary,
	}
)
var (
	filePath    string
	fileHash    string
	richAddress bool
	exportTxt   bool
)

func init() {
	sumaryCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to sumary.psk (required)")
	sumaryCmd.Flags().StringVarP(&fileHash, "address", "a", "", "Displays the address balance according to the sumary.psk file")
	sumaryCmd.PersistentFlags().BoolVarP(&richAddress, "richaddress", "r", false, "100 richest addresses")
	sumaryCmd.PersistentFlags().BoolVarP(&exportTxt, "export", "e", false, "Sumary.psk export to TXT")

	sumaryCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(sumaryCmd)

}

func runSumary(cmd *cobra.Command, args []string) {
	// Валідація шляху до файлу
	if filePath == "" {
		fmt.Println("Error: Path to the file is required.")
		cmd.Usage()
		os.Exit(1)
	}

	switch {
	case exportTxt:
		fmt.Println("Exporting summary to text file...")
		app.ExportSumaryToTxt(filePath)

	case richAddress:
		fmt.Println("100 richest addresses:")
		app.PrintRichAddress(filePath)

	case fileHash != "":
		fmt.Printf("Processing file '%s' with hash '%s'\n", filePath, fileHash)
		app.PrintDetailInfo(filePath, fileHash)

	default:
		fmt.Printf("Processing file -> '%s' \n", filePath)
		app.PrintTotalSummary(filePath)
	}
}

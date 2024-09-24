package commands

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pasichDev/nosotc/internal/app"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Handler for the Summary.psk file",
	Run:   runSummary,
}

var (
	fileHash    string
	richAddress bool
	exportTxt   bool
)

func init() {
	summaryCmd.Flags().StringVarP(&fileHash, "address", "a", "", "Displays the address balance according to the summary.psk file")
	summaryCmd.PersistentFlags().BoolVarP(&richAddress, "rich-address", "r", false, "Display 100 richest addresses")
	rootCmd.AddCommand(summaryCmd)
}

func runSummary(cmd *cobra.Command, args []string) {
	bold := "\033[1m"
	reset := "\033[0m"
	fileSummary := ""

	// Determine the file to process
	if nosoDataFolder != "" {
		fileSummary = filepath.Join(nosoDataFolder, "sumary.psk")
	} else {
		if nosoFilePath == "" {
			fmt.Printf(bold + "\n❌ You need to specify the path to the NOSODATA blockchain folder in the configuration. Otherwise, use the -f (--file) flag to specify the direct path to the block file \n" + reset)
			os.Exit(1)
		}
		fileSummary = nosoFilePath
		if !strings.HasSuffix(filepath.Base(fileSummary), ".psk") {
			fmt.Println("❌ Error: Path to file must end with .psk.")
			os.Exit(1)
		}
	}

	summaryHolder, err := app.NewSummaryDataHolder(fileSummary)
	if err != nil {
		fmt.Println("❌ Error initializing SummaryHandler:", err)
		return
	}

	switch {
	case exportTxt:
		fmt.Println("📄 Exporting summary to text file...")
		if err := summaryHolder.ExportSumaryToTxt(fileSummary); err != nil {
			fmt.Println("Error exporting summary:", err)
		}

	case richAddress:
		fmt.Println("\n💰 Displaying the 100 Richest Addresses:")
		listRich, err := summaryHolder.GetRichAddresses()
		if err != nil {
			fmt.Println("Error fetching rich addresses:", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "Hash", "Balance", "Custom"})

		for i := 0; i < 100 && i < len(listRich); i++ {
			table.Append([]string{
				fmt.Sprintf("%d", i+1),
				listRich[i].Hash,
				fmt.Sprintf("%.8f", listRich[i].Balance),
				listRich[i].Custom,
			})
		}
		table.Render()

	case fileHash != "":
		fmt.Printf("📂 Processing file '%s' with hash '%s'\n", fileSummary, fileHash)
		findHash, err := summaryHolder.GetDetailHash(fileHash)
		if err != nil {
			fmt.Println("Error fetching hash details:", err)
			return
		}
		customValue := findHash.Custom
		if customValue == "" {
			customValue = "null"
		}
		fmt.Printf(bold+"\n🔍 Hash: %s, 💰 Balance: %.8f, 🏷️ Custom: %s\n"+reset, findHash.Hash, findHash.Balance, customValue)

	default:
		fmt.Printf("📂 Processing file -> '%s'\n", fileSummary)
		totalSummary, err := summaryHolder.GetSumaryResume()
		if err != nil {
			fmt.Println("Error summarizing data:", err)
			return
		}
		fmt.Printf("\n💰 Total Balance: %d \n📦 Total Addresses: %d\n", totalSummary[0], totalSummary[1])
	}
}

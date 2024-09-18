package app

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"github.com/pasichDev/nosotc/noso/models"
	"github.com/pasichDev/nosotc/utils"
	"os"
	"sort"
	"strings"
	"time"
)

func getSummary(filePath string) ([]models.SummaryData, error) {
	bytesSummaryPsk, err := utils.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading file summary.psk")
	}

	summary, err := models.ParseSummaryData(bytesSummaryPsk)
	if err != nil {
		return nil, errors.New("error parsing file summary.psk")
	} else {
		return summary, nil
	}

}

func PrintTotalSummary(filePath string) {

	summary, err := getSummary(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var totalBalance float64

	for _, item := range summary {
		totalBalance += item.Balance
	}
	fmt.Println("\n")

	fmt.Printf("Total Balance: %.2f\n", totalBalance)
	fmt.Println("Total Addresses:", len(summary))

}

func PrintDetailInfo(filePath string, findHash string) {
	summary, err := getSummary(filePath)
	if err != nil {
		fmt.Println(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Balance", "Custom"})

	for _, v := range summary {
		if findHash == "" || v.Hash == findHash {
			table.Append([]string{v.Hash, fmt.Sprintf("%.8f", v.Balance), v.Custom})

		}
	}
	table.Render()

}

func PrintRichAddress(filePath string) {
	summary, err := getSummary(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Slice(summary, func(i, j int) bool {
		return summary[i].Balance > summary[j].Balance
	})

	limit := 100
	if len(summary) < 100 {
		limit = len(summary)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Hash", "Balance", "Custom"})

	for i := 0; i < limit; i++ {
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			summary[i].Hash,
			fmt.Sprintf("%.8f", summary[i].Balance),
			summary[i].Custom,
		})
	}

	table.Render()
}
func ExportSumaryToTxt(filePath string) {
	summary, err := getSummary(filePath)
	if err != nil {
		fmt.Println("Error while receiving data:", err)
		return
	}

	if len(summary) == 0 {
		fmt.Println("There is no data to export.")
		return
	}

	lastSlashIndex := strings.LastIndex(filePath, "/")
	if lastSlashIndex == -1 {
		fmt.Println("Could not find a slash in the string")
		return
	}

	filepathNew := filePath[:lastSlashIndex]

	var sb strings.Builder
	table := tablewriter.NewWriter(&sb)
	table.SetHeader([]string{"Hash", "Balance", "Custom"})

	totalBalance := 0.0

	for _, v := range summary {
		table.Append([]string{
			v.Hash,
			fmt.Sprintf("%.8f", v.Balance),
			v.Custom,
		})
		totalBalance += v.Balance
	}

	table.SetFooter([]string{fmt.Sprintf("Total Addresses: %d", len(summary)), fmt.Sprintf("Total Balance: %.2f", totalBalance), ""})

	table.Render()

	timestamp := time.Now().Format("20060102_150405")
	outputFilePath := fmt.Sprintf("%s/export_summary_%s.txt", filepathNew, timestamp)

	s := spinner.New(spinner.CharSets[11], 500*time.Millisecond)
	s.Prefix = "Saving a file...\n"
	s.Start()

	err = os.WriteFile(outputFilePath, []byte(sb.String()), 0644)
	s.Stop()
	if err != nil {
		fmt.Println("Error writing to a file:", err)
		return
	}

	fmt.Println("The file is successfully saved:", outputFilePath)
}

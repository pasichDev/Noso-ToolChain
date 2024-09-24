package app

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/pasichDev/nosotc/noso/models"
	"github.com/pasichDev/nosotc/utils"
	"os"
	"sort"
	"strings"
	"time"
)

type RichAddress struct {
	Hash    string
	Balance string
	Custom  string
}

// SummaryDataHolder holds the summary data and provides methods to access it
type SummaryDataHolder struct {
	Summary []models.SummaryData
}

// NewSummaryDataHolder initializes the SummaryDataHolder with data from the specified file
func NewSummaryDataHolder(filePath string) (*SummaryDataHolder, error) {
	summary, err := getSummary(filePath)
	if err != nil {
		return nil, err
	}
	return &SummaryDataHolder{Summary: summary}, nil
}

// getSummary reads the summary data from a file
func getSummary(filePath string) ([]models.SummaryData, error) {
	bytesSummaryPsk, err := utils.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading file summary.psk")
	}

	summary, err := models.ParseSummaryData(bytesSummaryPsk)
	if err != nil {
		return nil, errors.New("error parsing file summary.psk")
	}
	return summary, nil
}

// GetSumaryResume calculates the total balance and number of items
func (sdh *SummaryDataHolder) GetSumaryResume() ([2]int, error) {
	var totalBalance float64

	for _, item := range sdh.Summary {
		totalBalance += item.Balance
	}

	return [2]int{int(totalBalance), len(sdh.Summary)}, nil
}

// GetDetailHash displays details for a specific hash
func (sdh *SummaryDataHolder) GetDetailHash(findHash string) (models.SummaryData, error) {

	for _, v := range sdh.Summary {
		if findHash != "" && v.Hash == findHash {
			return v, nil
		}
	}

	return models.SummaryData{}, fmt.Errorf("hash %s not found", findHash)

}

// GetRichAddresses displays the richest addresses
func (sdh *SummaryDataHolder) GetRichAddresses() ([]models.SummaryData, error) {
	var result []models.SummaryData

	if len(sdh.Summary) == 0 {
		return result, errors.New("no rich addresses found")
	}

	sort.Slice(sdh.Summary, func(i, j int) bool {
		return sdh.Summary[i].Balance > sdh.Summary[j].Balance
	})

	limit := 100
	if len(sdh.Summary) < 100 {
		limit = len(sdh.Summary)
	}

	for i := 0; i < limit; i++ {
		result = append(result, sdh.Summary[i])
	}

	return result, nil
}

// ExportSumaryToTxt exports the summary data to a text file
func (sdh *SummaryDataHolder) ExportSumaryToTxt(filePath string) error {
	if len(sdh.Summary) == 0 {
		return errors.New("there is no data to export")
	}

	lastSlashIndex := strings.LastIndex(filePath, "/")
	if lastSlashIndex == -1 {
		return errors.New("could not find a slash in the string")
	}

	filepathNew := filePath[:lastSlashIndex]

	var sb strings.Builder

	totalBalance := 0.0

	for _, v := range sdh.Summary {
		sb.WriteString(fmt.Sprintf("  Hash: %s\n", v.Hash))
		sb.WriteString(fmt.Sprintf("  Balance: %.8f\n", v.Balance))
		if v.Custom != "" {
			sb.WriteString(fmt.Sprintf("  Custom: %s\n", v.Custom))
		} else {
			sb.WriteString("  Custom: nil\n") // Show nil if Custom is empty
		}
		sb.WriteString("\n") // Add an empty line between entries
		totalBalance += v.Balance
	}

	// Summary footer
	sb.WriteString(fmt.Sprintf("Total Addresses: %d\n", len(sdh.Summary)))
	sb.WriteString(fmt.Sprintf("Total Balance: %.2f\n", totalBalance))

	timestamp := time.Now().Format("20060102_150405")
	outputFilePath := fmt.Sprintf("%s/export_summary_%s.txt", filepathNew, timestamp)

	s := spinner.New(spinner.CharSets[11], 500*time.Millisecond)
	s.Prefix = "Saving a file...\n"
	s.Start()

	err := os.WriteFile(outputFilePath, []byte(sb.String()), 0644)
	s.Stop()
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println("\n✅ The file is successfully saved:", outputFilePath)
	return nil
}

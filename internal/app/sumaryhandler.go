package app

import (
	"errors"
	"fmt"
	"github.com/Friends-Of-Noso/NosoData-Go/legacy"
	"github.com/Friends-Of-Noso/NosoData-Go/utils"
	"github.com/briandowns/spinner"
	"os"
	"sort"
	"strconv"
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
	Summary legacy.LegacySummary
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
func getSummary(filePath string) (legacy.LegacySummary, error) {

	summary := legacy.LegacySummary{}
	err := summary.ReadFromFile(filePath)
	if err != nil {
		return legacy.LegacySummary{}, errors.New("error reading file summary.psk")
	}
	return summary, nil
}

// GetSummaryResume calculates the total balance and number of items
func (sdh *SummaryDataHolder) GetSummaryResume() ([2]int, error) {
	var totalBalance int64

	for _, item := range sdh.Summary.Accounts {
		totalBalance += item.Balance
	}

	num, err := strconv.Atoi(utils.ToNoso(totalBalance))
	if err != nil {
		return [2]int{}, err
	}
	return [2]int{num, len(sdh.Summary.Accounts)}, nil
}

// GetDetailHash displays details for a specific hash
func (sdh *SummaryDataHolder) GetDetailHash(findHash string) (legacy.LegacySummaryAccount, error) {

	for _, v := range sdh.Summary.Accounts {
		if findHash != "" && v.Hash.GetString() == findHash {
			return v, nil
		}
	}

	return legacy.LegacySummaryAccount{}, fmt.Errorf("hash %s not found", findHash)

}

// GetRichAddresses displays the richest addresses
func (sdh *SummaryDataHolder) GetRichAddresses() ([]legacy.LegacySummaryAccount, error) {
	var result []legacy.LegacySummaryAccount

	if len(sdh.Summary.Accounts) == 0 {
		return result, errors.New("no rich addresses found")
	}

	// Ось тут виправляємо сортування
	sort.Slice(sdh.Summary.Accounts, func(i, j int) bool {
		return sdh.Summary.Accounts[i].Balance > sdh.Summary.Accounts[j].Balance
	})

	limit := 100
	if len(sdh.Summary.Accounts) < 100 {
		limit = len(sdh.Summary.Accounts)
	}

	for i := 0; i < limit; i++ {
		result = append(result, sdh.Summary.Accounts[i])
	}

	return result, nil
}

// ExportSummaryToTxt exports the summary data to a text file
func (sdh *SummaryDataHolder) ExportSummaryToTxt(filePath string) error {
	if len(sdh.Summary.Accounts) == 0 {
		return errors.New("there is no data to export")
	}

	lastSlashIndex := strings.LastIndex(filePath, "/")
	if lastSlashIndex == -1 {
		return errors.New("could not find a slash in the string")
	}

	filepathNew := filePath[:lastSlashIndex]

	var sb strings.Builder

	var totalBalance int64

	for _, v := range sdh.Summary.Accounts {
		sb.WriteString(fmt.Sprintf("  Hash: %s\n", v.Hash.GetString()))
		sb.WriteString(fmt.Sprintf("  Balance: %s\n", utils.ToNoso(v.Balance)))
		if v.Custom.GetString() != "" {
			sb.WriteString(fmt.Sprintf("  Custom: %s\n", v.Custom.GetString()))
		} else {
			sb.WriteString("  Custom: nil\n") // Show nil if Custom is empty
		}
		sb.WriteString("\n") // Add an empty line between entries
		totalBalance += v.Balance
	}

	// Summary footer
	sb.WriteString(fmt.Sprintf("Total Addresses: %d\n", len(sdh.Summary.Accounts)))
	sb.WriteString(fmt.Sprintf("Total Balance: %s\n", utils.ToNoso(totalBalance)))

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

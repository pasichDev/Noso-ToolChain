package models

import (
	"errors"
	"github.com/pasichDev/nosotc/noso"
)

type SummaryData struct {
	Hash    string
	Custom  string
	Balance float64
}

func ParseSummaryData(bytesSummaryPsk []byte) ([]SummaryData, error) {
	if len(bytesSummaryPsk) == 0 {
		return []SummaryData{}, nil
	}

	var addressSummary []SummaryData
	index := 0
	for index+106 <= len(bytesSummaryPsk) {
		sumData := SummaryData{}

		// Hash
		hashLength := int(bytesSummaryPsk[index])
		if index+1+hashLength > len(bytesSummaryPsk) {
			return nil, errors.New("invalid hash length")
		}
		sumData.Hash = string(bytesSummaryPsk[index+1 : index+1+hashLength])

		// Custom
		customLength := int(bytesSummaryPsk[index+41])
		if index+42+customLength > len(bytesSummaryPsk) {
			return nil, errors.New("invalid custom length")
		}
		sumData.Custom = string(bytesSummaryPsk[index+42 : index+42+customLength])

		// Balance
		balanceBytes := bytesSummaryPsk[index+82 : index+90]

		sumData.Balance += noso.BytesToFloat64(balanceBytes)
		addressSummary = append(addressSummary, sumData)
		index += 106
	}

	return addressSummary, nil
}

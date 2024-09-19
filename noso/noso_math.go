package noso

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type DecimalIntegerModel struct {
	intValue     string
	decimalValue string
	doubleValue  float64
}

func (d *DecimalIntegerModel) GetInt() string {
	return d.intValue
}

func (d *DecimalIntegerModel) SetInt(value string) {
	d.intValue = value
}

func (d *DecimalIntegerModel) GetDecimal() string {
	return d.decimalValue
}

func (d *DecimalIntegerModel) SetDecimal(value string) {
	d.decimalValue = value
}

func (d *DecimalIntegerModel) GetDouble() float64 {
	return float64(d.doubleValue)
}

func (d *DecimalIntegerModel) SetDouble(value float64) {
	d.doubleValue = value
}

// ConvertFromBigInt converts a large integer into the DecimalIntegerModel
func ConvertFromBigInt(value int64) (DecimalIntegerModel, error) {
	var returnData DecimalIntegerModel

	// Перетворення uint64 на строку
	stringValue := strconv.FormatInt(value, 10)
	length := len(stringValue)

	// Якщо значення дорівнює 0
	if value == 0 {
		returnData.SetDouble(0.00000000)
		return returnData, nil
	}

	if length <= 8 {
		// Якщо число має менше 8 цифр, просто ділимо його для отримання дробової частини
		doubleValue := float64(value) / 100000000.0
		returnData.SetInt(strconv.Itoa(int(value)))
		returnData.SetDouble(doubleValue)
		return returnData, nil
	} else {
		// Якщо число має більше 8 цифр
		integerPart := stringValue[:length-8]
		decimalPart := stringValue[length-8:]

		if decimalPart == "" {
			// Якщо немає дробової частини
			returnData.SetInt(integerPart)
			returnData.SetDouble(float64(value) / 100000000.0)
		} else {
			// Якщо є дробова частина
			returnData.SetInt(integerPart)
			returnData.SetDecimal(decimalPart)
			doubleValue := float64(value) / 100000000.0
			returnData.SetDouble(doubleValue)
		}
	}

	return returnData, nil
}

func BytesToFloat64(b []byte) float64 {
	if len(b) != 8 {
		return 0.0
	}
	bits := int64(binary.LittleEndian.Uint64(b))
	result, err := ConvertFromBigInt(bits)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		return result.GetDouble()
	}
	return 0.0
}

package main

import (
	"fmt"
	"testing"
)

func TestConvertCurrencyToTargets(t *testing.T){
	tests:= map[Currency]CurrencyData{
        "INR": {Code: "INR", Value: 83.56},
		"MYR": {Code: "MYR", Value: 4.14},
    }
	
	arr := convertCurrencyToTargets(1, CurrencyData{Code: "SAR", Value: 3.75},tests)
	for key, value := range arr{
		fmt.Println(key)
		fmt.Println(value)
	}
}
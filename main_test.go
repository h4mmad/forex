package main

import (
	"testing"
)

// func TestCheckSyntax(t *testing.T) {
//     // Define test cases with input and expected result
//     tests := []struct {
//         input    string
//         expected bool
//     }{
//         {"1", true},        // Valid integer
//         {"2.2", true},      // Valid float
//         {" ", false},       // Invalid input (whitespace)
//         {"sas", false},     // Invalid input (non-numeric)
// 		{"-42", false},      // Invalid negative integer
//         {"2.5.6", false},   // Invalid input (multiple decimal points)
//         {"-2.5", false},   // Invalid input (multiple decimal points)
//     }

//     for _, test := range tests {
//         result := isAmountValid(test.input)
//         if result != test.expected {
//             t.Errorf("isSyntaxCorrect(%q) = %v; expected %v", test.input, result, test.expected)
//         }
//     }
// }

func TestIsValidCurrency(t *testing.T){
	tests := []struct {
		input CurrencyCode
		expected bool
	}{
		{"MYR", true},
		{"SAR", true},
		{"SARS", false},
		{" ", false},
		{"JPN", false},
	}

	for _, test := range tests{
		result := isValidCurrency(test.input)
		if result != test.expected{
			t.Errorf("isValidCurrency(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
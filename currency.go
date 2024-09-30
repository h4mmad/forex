package main

// Define structs to match the JSON structure
type Meta struct{
    LastCheckedAt string `json:"last_checked_at"`
}

type Currency struct{
    Code string `json:"code"`
	Value float64 `json:"value"`
}

type CurrencyData struct {
    Meta Meta            `json:"meta"`
    Data map[string]Currency  `json:"data"`
}


// convertSourceToTarget converts an amount from a source currency to a target currency.
// It calculates the equivalent value based on the provided currency rates.
//
// Parameters:
//   - amount: The amount of the source currency to be converted.
//   - sourceCurrency: A CurrencyRate struct representing the source currency's rate.
//   - targetCurrency: A CurrencyRate struct representing the target currency's rate.
//
// Returns:
//   - The converted amount in the target currency as a float64. The conversion is based
//     on the ratio of the target currency's value to the source currency's value.
//
// Example:
//   sourceRate := CurrencyRate{Code: "USD", Value: 1.0, LastCheckedAt: time.Now()}
//   targetRate := CurrencyRate{Code: "MYR", Value: 4.13, LastCheckedAt: time.Now()}
//   amountInMYR := convertSourceToTarget(10.0, sourceRate, targetRate)
//   // amountInMYR will contain the converted value from USD to MYR.
func convertSourceToTarget(amount float64, sourceCurrency Currency, targetCurrency Currency) float64 {
    convertedAmount := amount * (targetCurrency.Value / sourceCurrency.Value)
    return convertedAmount
}

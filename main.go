package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
    version   string
    cacheExpiryDurationInHours string
)

func inputValidationChecks() (float64, string, []string){
    if len(os.Args) < 4{
        log.Fatalf("Few arguments provided")
    }
    // Check if the input baseAmount is valid
    sourceAmount, err := strconv.ParseFloat(os.Args[1], 64)
    if err != nil {
        
    }
    if sourceAmount < 0{
        log.Fatal("Amount should be a positive value")
    }
    
    // Check if the base CurrencyCode is valid
    sourceCurrency := os.Args[2]
   

    // Check if the target currencies is valid
    var targetCurrencies []string
    for i := 3; i < len(os.Args); i++ {
        targetCurrency := os.Args[i]
            targetCurrencies = append(targetCurrencies, targetCurrency)
    }

    return sourceAmount, sourceCurrency, targetCurrencies
}



func main() {
    expiry, err :=strconv.ParseFloat(cacheExpiryDurationInHours, 64)
    if err != nil{
        log.Fatalf("Cannot parse: %v", err)
    }
    //use a channel to send apiResponse to a go subroutine to write data concurrently
    fmt.Printf("Version: %s\n", version)

    ch := make(chan CurrencyData)
    wg := sync.WaitGroup{}
  
    
    sourceAmount, sourceCurrency, targetCurrencies := inputValidationChecks()

    if !isCacheAvailable(){
        createCache()
    }
    if isCacheEmpty(){
        response := makeGETrequest()
        wg.Add(1)
        go writeToCache(ch, &wg)
        ch <- response
        wg.Wait()
        close(ch)
    }

    // reading starts from here
    cacheData := readCacheData()
    lastCheckedAt, err:= time.Parse(time.RFC3339, cacheData.Meta.LastCheckedAt)
    if err != nil {
        log.Fatalf("Error parsing time: %v", err)
    }
    if isCacheExpired(lastCheckedAt, expiry){
        response := makeGETrequest()
        wg.Add(1)
        go writeToCache(ch, &wg)
        ch <- response
        wg.Wait()
        close(ch)
        fmt.Println("Cache updated")
    }

    //If the key exists in the map, 'sourceValue' holds the value of that key
    sourceValue,exists := cacheData.Data[sourceCurrency]
    if !exists {
        log.Fatal("Source Currency does not exist")
    }
    fmt.Printf("%s %.2f equals:\n\n", sourceCurrency, sourceAmount)
    for _,targetCurrencyCode:= range targetCurrencies{
        targetValue,exists:= cacheData.Data[targetCurrencyCode]
        if !exists {
            log.Fatalf("Target currency does not exist %v", targetCurrencyCode)
        }
        fmt.Printf("%s %.2f\n", targetValue.Code, convertSourceToTarget(sourceAmount, sourceValue, targetValue))
    }   

}


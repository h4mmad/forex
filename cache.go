package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const cacheName = "cache.json"


/*
  * The function takes params of a channel, waitgroup, and filename.
    File is created if doesn't exist, opened if exists. API response is read from
    from the channel and written to file cache in json format. Function is used
    concurrently.
*/
func writeToCache(ch <-chan CurrencyData,wg *sync.WaitGroup) {
    defer wg.Done()
   

    // Create the file, but ensure it does not recreate if it already exists
    file, err := os.OpenFile(cacheName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        log.Fatalf("Failed to create or open file: %v", err)
    }
    defer file.Close() // Ensure the file is closed when the function exits

    response := <-ch
 
    // Marshal the JSON data
    jsonData, err := json.MarshalIndent(response, "", " ")
    if err != nil {
        log.Fatalf("Failed to marshal JSON: %v", err)
    }

    // Write to the file
    if _, err := file.Write(jsonData); err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }

    fmt.Printf("API response written to %s \n", cacheName)
}



/*
  * Takes in the cache time stamp and duration as params.
    cacheTimeStamp is the last time data was updated in cache.
	The function diffs the cacheTimeStamp with time.Now(). If the
	differnce exceedes duration in hours return true.
*/
func isCacheExpired(cacheTimeStamp time.Time, duration float64) bool{
	now:= time.Now()
	diff:= now.Sub(cacheTimeStamp)
	return diff.Hours() >= duration
}


/*
  * Takes in the cache time stamp and duration as params.
    cacheTimeStamp is the last time data was updated in cache.
	The function diffs the cacheTimeStamp with time.Now(). If the
	differnce exceedes duration in hours return true.
*/
func readCacheData() CurrencyData{
	cache, err:= os.ReadFile(cacheName)
   	if err != nil{
		log.Fatalf("Error reading file: %v", err)
	}

	var cacheData CurrencyData
    
    if err != nil{
        log.Fatal("Error")
    }
	err = json.Unmarshal(cache, &cacheData)
	if err != nil{
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return cacheData
}

/*
    isCacheAvailable() checks whether the cache file exists, it will always run at the start of the program.
    Returns true if exists, returns false if doens't exist
*/
func isCacheAvailable() bool{
    _,err:= os.Stat(cacheName)
    return err == nil
}

/*
    createCache() creates a cache file.
*/
func createCache(){
        f, err :=os.Create(cacheName)
        if err != nil{
            log.Fatalf("Error occured while creating cache: %v", err)
        }else{
            fmt.Println(f)
        }
    }

/*
    isCacheEmpty() checks whether the file contents of the cache is empty.
    Returns true is cache is empty, returns false if cache not empty 
*/
func isCacheEmpty() bool {
    fi, err := os.Stat(cacheName)
    if err != nil {
        log.Fatalf("Error path: %v", err)
    }

    return fi.Size() == 0
}
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //sets the max number of processors this will use
	//runtime.GOMAXPROCS(1)	//sets the max number of processors this will use

	start := time.Now()

	stockSymbols := []string{"nymt", "sail", "snap", "ge", "ibm", "tsla", "aapl"}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body) //this could blow up memory, if the response body is large

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s, %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	// loop until all the symbols have been processed (this keeps the main thead alive)
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("%v\n", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	Symbol           string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	Timestamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

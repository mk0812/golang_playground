package main

import (
	"fmt"
	"os"

	"github.com/DannyBen/quandl"
)

func main() {
	quandl.APIKey = os.Getenv("QUANDL")
	data, _ := quandl.GetSymbol("WIKI/AAPL", nil)

	for i, item := range data.Data {
		fmt.Println(i, item[0], item[2])
	}
}

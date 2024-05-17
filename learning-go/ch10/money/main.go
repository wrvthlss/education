package main

import (
	"fmt"
	"github.com/learning-go-book-2e/formatter"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Need to parameters: amount and percent")
		os.Exit(1)
	}

	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	percent, err := decimal.NewFromString(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	percent = percent.Div(decimal.NewFromInt(100))
	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(formatter.Space(80, os.Args[1], os.Args[2], total.StringFixed(2)))

}

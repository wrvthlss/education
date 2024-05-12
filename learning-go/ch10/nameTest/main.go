package main

import (
	"fmt"
	config "github.com/wrvthlss/education/learning-go/ch10/nameTest/config"
	logs "github.com/wrvthlss/education/learning-go/ch10/nameTest/logs"
	names "github.com/wrvthlss/education/learning-go/ch10/nameTest/names"
	"log"
)

func main() {
	cfg, err := config.LoadConfiguration("config/config.json")
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	logger := logs.SetupLogger(cfg.LogFilePath)

	str := "ascensionism"

	extracted, err := names.Extract(str)
	if err != nil {
		fmt.Println("an error has occurred, please see logs for more details.")
		logger.Printf("error in extraction: %v", err)
	}

	uppercased := names.Format(extracted)
	fmt.Println(uppercased)

}

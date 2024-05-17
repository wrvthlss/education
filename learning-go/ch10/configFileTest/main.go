package main

import (
	"fmt"
	"github.com/wrvthlss/education/ch10/configFileTest/config"
	"log"
)

func main() {

	cfg, err := config.LoadConfig("./key.json")
	if err != nil {
		log.Fatalln("Failed to load json configuration: %v", err)
	}

	fmt.Println("Using API Key: ", cfg.APIKey)

}

package main

import (
	"fmt"
	"github.com/wrvthlss/education/learning-go/ch10/pLevelVarsInit/config"
)

func main() {

	apiKey := config.GetConfig("apiKey")
	fmt.Println("Using API Key: ", apiKey)
}

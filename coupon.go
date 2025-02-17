package main

import (
	"flag"
	"fmt"

	"github.com/eserilev/migration.winc.services/corporate"
)

func main() {
	filePath := flag.String("file", "", "a file path")
	userGuid := flag.String("guid", "", "a user guid")
	billingProfileId := flag.Int("billing", 0, "a billing profile id")
	brand := flag.Int("brand", 0, "a brand id")
	invoice := flag.Bool("invoice", false, "invoice flag")
	flag.Parse()

	result := corporate.ProcessCorporateOrders(*filePath, *userGuid, *invoice, *billingProfileId, *brand)
	if !result.Success {
		fmt.Println("Failed to process order")
		return
	}

	fmt.Println("Corporate orders submitted succesfully")
}

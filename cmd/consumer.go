package main

import (
	"os"
	"strconv"

	"github.com/gbrancods/tiendanube-consumer/runners"
	"github.com/gbrancods/tiendanube/auth"
)

func main() {

	id, _ := strconv.Atoi(os.Getenv("USER_ID"))

	auth.Access{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		TokenType:   "bearer",
		UserID:      id,
		APIurl:      "api.nuvemshop.com.br",
	}.Start()

	// Working!
	runners.CategoryRunner()

	// Working!
	runners.WebhookRunner()

	// Working!
	runners.GetStoreInfo()

	// Working!
	runners.CouponsRunner()

	// Working!
	runners.CustomerRunner()

	// Working partialy
	// runners.ProductRunner()

	// Working partialy
	// runners.LocationsRunner()

	// Working partialy
	// runners.DraftOrderRunner()

	// Working partialy
	// runners.LocationsRunner()

	// Working partialy
	// runners.AbandonedCheckoutRunner()

	// Not work
	// runners.MetafieldRunner()
}

package main

import (
	"go-laundry/delivery"
)

func main() {
	// CLI
	// delivery.NewConsole().Run()

	//Server
	delivery.NewServer().Run()
}
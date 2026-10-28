package main

import (
	"context"
	"log"
	"os"

	"github.com/mehdi-fakhraei60/DayPay/gateway"
)

func main() {
	if err := gateway.Prove(context.Background(), os.Getenv("GATEWAY_IDENTITY_URL"), "dpey"); err != nil {
		log.Fatal(err)
	}
}

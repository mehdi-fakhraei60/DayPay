package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mehdi-fakhraei60/DayPay/gateway"
)

func main() {
	if err := gateway.Prove(context.Background(), os.Getenv("GATEWAY_IDENTITY_URL"), "dpey"); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"app\":\"dpey\",\"proved\":true}")
	})
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}

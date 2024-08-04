package main


import (
	"fmt"
	"net/http"
	"log"
)

func main()
{
	port := 80;
	log/Printf("Listening on port %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
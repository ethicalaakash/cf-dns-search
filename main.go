package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	targetURL := flag.String("url", "test.com", "Target URL to search in DNS records")
	flag.Parse()
	// Initialize the Cloudflare API client
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if apiToken == "" {
		log.Fatal("CLOUDFLARE_API_TOKEN environment variable is not set")
	}

	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatalf("Failed to create Cloudflare API client: %v", err)
	}

	// Fetch all zones
	zones, err := api.ListZones(context.Background())
	if err != nil {
		log.Fatalf("Failed to list zones: %v", err)
	}

	// Iterate over each zone
	for _, zone := range zones {
		fmt.Printf("Zone: %s\n", zone.Name)

		// Fetch DNS records for the zone
		records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zone.ID), cloudflare.ListDNSRecordsParams{})
		if err != nil {
			log.Printf("Failed to fetch DNS records for zone %s: %v", zone.Name, err)
			continue
		}

		// Filter records pointing to provided URL
		for _, record := range records {
			if strings.HasSuffix(record.Content, *targetURL) {
				fmt.Printf("Record: %s -> %s\n", record.Name, record.Content)
			}
		}
	}
}

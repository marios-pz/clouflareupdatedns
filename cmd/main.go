package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	zoneID := os.Getenv("ZONE_ID")
	authToken := os.Getenv("AUTH_TOKEN")
	dnsName := os.Getenv("DNS_NAME")
	dnsRecordID := os.Getenv("DNS_RECORD_ID")
	publicIP := os.Getenv("PUBLIC_IP")

	// Cloudflare API endpoint for updating DNS records
	apiURL := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		zoneID,
		dnsRecordID,
	)

	payload := []byte(
		fmt.Sprintf(`{"type":"A","name":"%s","content":"%s"}`, dnsName, publicIP),
	)

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status code:", resp.StatusCode)
		return
	}

	fmt.Println("DNS record updated successfully!")
}

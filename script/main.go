package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type SendCouriorRequest struct {
	ProductID           uint      `json:"product_id"`
	UserID              uint      `json:"user_id"`
	SourceLocation      string    `json:"source_location"`
	DestinationLocation string    `json:"destination_location"`
	StartTime           time.Time `json:"start_time"`
}

func main() {
	// set duration for script run (2 days)
	duration := 2 * 24 * time.Hour
	endTime := time.Now().Add(duration)

	fmt.Printf("Script will run until: %v\n", endTime)

	for time.Now().Before(endTime) {
		// seed random number with the current time
		rand.Seed(time.Now().UnixNano())

		sendRandomRequest()

		// wait for random between 1 second and 5 minutes
		waitTime := time.Duration(rand.Intn(300)+1) * time.Second
		time.Sleep(waitTime)
	}

	fmt.Println("Script completed.")
}

func sendRandomRequest() {
	request := SendCouriorRequest{
		ProductID:           uint(rand.Intn(1000) + 1),
		UserID:              uint(rand.Intn(1000) + 1),
		SourceLocation:      randomIranianCity(),
		DestinationLocation: randomIranianCity(),
		StartTime:           randomTime(),
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8042/v1/deliveries", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Request sent at %v. Status: %s\n", time.Now().Format(time.RFC3339), resp.Status)
}

func randomIranianCity() string {
	cities := []string{
		"Tehran", "Isfahan", "Shiraz", "Mashhad",
		"Tabriz", "Karaj", "Kermanshah",
	}
	return cities[rand.Intn(len(cities))]
}

func randomTime() time.Time {
	now := time.Now()
	randomDuration := time.Duration(rand.Intn(30*24)) * time.Hour
	return now.Add(-randomDuration)
}

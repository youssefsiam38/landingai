package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/youssefsiam38/landingai"
)

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("LANDINGAI_API_KEY")
	if apiKey == "" {
		log.Fatal("LANDINGAI_API_KEY environment variable is required")
	}

	// Example 1: Basic usage - Parse from file
	fmt.Println("Example 1: Parse from file")
	basicExample(apiKey)

	// Example 2: Parse from URL
	fmt.Println("\nExample 2: Parse from URL")
	urlExample(apiKey)

	// Example 3: Parse with specific model
	fmt.Println("\nExample 3: Parse with specific model")
	modelExample(apiKey)

	// Example 4: Parse with page splitting
	fmt.Println("\nExample 4: Parse with page splitting")
	splitExample(apiKey)

	// Example 5: EU region
	fmt.Println("\nExample 5: Parse with EU region")
	euRegionExample(apiKey)

	// Example 6: Error handling
	fmt.Println("\nExample 6: Error handling")
	errorHandlingExample(apiKey)
}

func basicExample(apiKey string) {
	// Create a new client
	client := landingai.NewClient(apiKey)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Parse a document
	result, err := client.Parse(ctx).
		WithFile("document.pdf").
		Do()

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Parsed %d pages in %d ms\n", result.Metadata.PageCount, result.Metadata.DurationMs)
	fmt.Printf("Credit usage: %.2f\n", result.Metadata.CreditUsage)
	fmt.Printf("Number of chunks: %d\n", len(result.Chunks))
	fmt.Printf("Markdown preview: %s...\n", result.Markdown[:min(100, len(result.Markdown))])
}

func urlExample(apiKey string) {
	client := landingai.NewClient(apiKey)
	ctx := context.Background()

	// Parse from URL
	result, err := client.Parse(ctx).
		WithURL("https://example.com/document.pdf").
		Do()

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully parsed document from URL\n")
	fmt.Printf("Number of chunks: %d\n", len(result.Chunks))
}

func modelExample(apiKey string) {
	client := landingai.NewClient(apiKey)
	ctx := context.Background()

	// Parse with specific model version
	result, err := client.Parse(ctx).
		WithFile("document.pdf").
		WithModel("dpt-2-latest").
		Do()

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Parsed with model: %s\n", *result.Metadata.Version)
	fmt.Printf("Number of chunks: %d\n", len(result.Chunks))
}

func splitExample(apiKey string) {
	client := landingai.NewClient(apiKey)
	ctx := context.Background()

	// Parse with page splitting
	result, err := client.Parse(ctx).
		WithFile("document.pdf").
		WithPageSplit().
		Do()

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Number of splits: %d\n", len(result.Splits))
	for i, split := range result.Splits {
		fmt.Printf("  Split %d: pages %v, %d chunks\n", i, split.Pages, len(split.Chunks))
	}
}

func euRegionExample(apiKey string) {
	// Create client with EU region
	client := landingai.NewClient(apiKey, landingai.WithRegion(landingai.RegionEU))
	ctx := context.Background()

	result, err := client.Parse(ctx).
		WithFile("document.pdf").
		Do()

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully parsed with EU region\n")
	fmt.Printf("Number of chunks: %d\n", len(result.Chunks))
}

func errorHandlingExample(apiKey string) {
	client := landingai.NewClient(apiKey)
	ctx := context.Background()

	// Intentionally cause an error (missing file)
	_, err := client.Parse(ctx).
		WithFile("nonexistent.pdf").
		Do()

	if err != nil {
		// Check if it's an API error
		if apiErr, ok := err.(*landingai.APIError); ok {
			fmt.Printf("API Error (Status %d): %s\n", apiErr.StatusCode, apiErr.Message)

			// Check specific error types
			if apiErr.IsUnauthorized() {
				fmt.Println("  -> Invalid API key")
			} else if apiErr.IsPaymentRequired() {
				fmt.Println("  -> Insufficient credits")
			} else if apiErr.IsRateLimited() {
				fmt.Println("  -> Rate limit exceeded")
			}
		} else if valErr, ok := err.(*landingai.ValidationErrors); ok {
			fmt.Printf("Validation Error: %v\n", valErr)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

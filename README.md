# Landing AI Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/youssefsiam38/landingai.svg)](https://pkg.go.dev/github.com/youssefsiam38/landingai)
[![Go Report Card](https://goreportcard.com/badge/github.com/youssefsiam38/landingai)](https://goreportcard.com/report/github.com/youssefsiam38/landingai)

The official Go SDK for [Landing AI](https://landing.ai/)'s Agentic Document Extraction (ADE) API. Parse documents and spreadsheets into structured data with state-of-the-art AI models.

## Features

- 🚀 **Simple & Intuitive API** - Clean, idiomatic Go with fluent builder pattern
- 📄 **Document Parsing** - Extract structured data from PDFs, images, and spreadsheets
- 🌍 **Multi-Region Support** - US and EU regions available
- 🎯 **Type-Safe** - Full type definitions for all API responses
- ⚡ **Flexible** - Parse from files or URLs
- 🛡️ **Error Handling** - Comprehensive error types and validation
- 🔧 **Configurable** - Custom HTTP clients, timeouts, and more

## Installation

```bash
go get github.com/youssefsiam38/landingai
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/youssefsiam38/landingai"
)

func main() {
    // Create a client with your API key
    client := landingai.NewClient("your-api-key")

    // Parse a document
    result, err := client.Parse(context.Background()).
        WithFile("invoice.pdf").
        Do()

    if err != nil {
        log.Fatal(err)
    }

    // Access the parsed data
    fmt.Printf("Parsed %d pages\n", result.Metadata.PageCount)
    fmt.Printf("Found %d chunks\n", len(result.Chunks))
    fmt.Printf("Markdown:\n%s\n", result.Markdown)
}
```

## Authentication

Get your API key from the [Landing AI Platform](https://va.landing.ai/settings/api-key):

- **US Region**: https://va.landing.ai/settings/api-key
- **EU Region**: https://va.eu-west-1.landing.ai/settings/api-key

```go
// US Region (default)
client := landingai.NewClient("your-api-key")

// EU Region
client := landingai.NewClient("your-api-key",
    landingai.WithRegion(landingai.RegionEU))
```

## Usage Examples

### Parse from File

```go
result, err := client.Parse(ctx).
    WithFile("document.pdf").
    Do()
```

### Parse from URL

```go
result, err := client.Parse(ctx).
    WithURL("https://example.com/document.pdf").
    Do()
```

### Parse with Specific Model

Landing AI offers multiple parsing models:
- `dpt-2-latest` - Latest DPT-2 model (recommended)
- `dpt-2-20250919` - Specific DPT-2 snapshot
- `dpt-1-latest` - DPT-1 model
- `DPT-2-mini-latest` - Lightweight model for simple documents

```go
result, err := client.Parse(ctx).
    WithFile("document.pdf").
    WithModel("dpt-2-latest").
    Do()
```

### Parse with Page Splitting

Split documents into page-level sections:

```go
result, err := client.Parse(ctx).
    WithFile("document.pdf").
    WithPageSplit().
    Do()

// Access splits
for _, split := range result.Splits {
    fmt.Printf("Pages %v: %d chunks\n", split.Pages, len(split.Chunks))
}
```

### Parse with File Data (In-Memory)

```go
data, err := os.ReadFile("document.pdf")
if err != nil {
    log.Fatal(err)
}

result, err := client.Parse(ctx).
    WithFileData(data, "document.pdf").
    Do()
```

## Configuration

### Custom HTTP Client

```go
httpClient := &http.Client{
    Timeout: 10 * time.Minute,
    Transport: &http.Transport{
        MaxIdleConns: 10,
    },
}

client := landingai.NewClient("your-api-key",
    landingai.WithHTTPClient(httpClient))
```

### Custom Timeout

```go
client := landingai.NewClient("your-api-key",
    landingai.WithTimeout(5 * time.Minute))
```

### Custom Base URL

```go
client := landingai.NewClient("your-api-key",
    landingai.WithBaseURL("https://custom-endpoint.com"))
```

## Response Structure

The `ParseResponse` contains rich structured data:

```go
type ParseResponse struct {
    Markdown  string                            // Full document as Markdown
    Chunks    []ParseChunk                      // Extracted content chunks
    Splits    []ParseSplit                      // Document splits (if enabled)
    Grounding map[string]ParseResponseGrounding // Bounding box coordinates
    Metadata  ParseMetadata                     // Parsing metadata
}
```

### Chunks

Each chunk represents a discrete element from the document:

```go
type ParseChunk struct {
    Markdown  string         // Chunk content as Markdown
    Type      string         // text, table, figure, logo, etc.
    ID        string         // Unique chunk identifier
    Grounding ParseGrounding // Location in document
}
```

**Chunk Types:**
- `text` - Paragraphs, headings, lists, form fields
- `table` - Tables and structured data
- `marginalia` - Headers, footers, page numbers
- `figure` - Images, graphs, diagrams
- `logo` - Company logos (DPT-2 only)
- `card` - ID cards, licenses (DPT-2 only)
- `attestation` - Signatures, stamps, seals (DPT-2 only)
- `scan_code` - QR codes, barcodes (DPT-2 only)

### Grounding (Bounding Boxes)

Each chunk includes its location in the original document:

```go
type ParseGrounding struct {
    Box  ParseGroundingBox // Relative coordinates (0-1)
    Page int               // Zero-indexed page number
}

type ParseGroundingBox struct {
    Left   float64 // Left coordinate (0-1)
    Top    float64 // Top coordinate (0-1)
    Right  float64 // Right coordinate (0-1)
    Bottom float64 // Bottom coordinate (0-1)
}
```

### Metadata

```go
type ParseMetadata struct {
    Filename    string   // Original filename
    PageCount   int      // Number of pages
    DurationMs  int      // Processing time in milliseconds
    CreditUsage float64  // Credits consumed
    JobID       string   // Unique job identifier
    Version     *string  // Model version used
    FailedPages []int    // Pages that failed (if any)
}
```

## Error Handling

The SDK provides comprehensive error handling:

```go
result, err := client.Parse(ctx).
    WithFile("document.pdf").
    Do()

if err != nil {
    // Check for API errors
    if apiErr, ok := err.(*landingai.APIError); ok {
        switch {
        case apiErr.IsUnauthorized():
            fmt.Println("Invalid API key")
        case apiErr.IsPaymentRequired():
            fmt.Println("Insufficient credits")
        case apiErr.IsRateLimited():
            fmt.Println("Rate limit exceeded - please retry")
        case apiErr.IsBadRequest():
            fmt.Println("Invalid request parameters")
        case apiErr.IsValidationError():
            fmt.Println("Input validation failed")
        case apiErr.IsServerError():
            fmt.Println("Server error - please retry")
        case apiErr.IsTimeout():
            fmt.Println("Request timeout")
        }

        fmt.Printf("Status: %d\n", apiErr.StatusCode)
        fmt.Printf("Message: %s\n", apiErr.Message)
        fmt.Printf("Detail: %v\n", apiErr.Detail)
    }

    // Check for validation errors
    if valErr, ok := err.(*landingai.ValidationErrors); ok {
        for _, detail := range valErr.Detail {
            fmt.Printf("Validation error: %s\n", detail.Message)
        }
    }
}
```

### Common Error Status Codes

- `400` - Bad Request (invalid parameters)
- `401` - Unauthorized (invalid API key)
- `402` - Payment Required (insufficient credits)
- `422` - Unprocessable Entity (validation error)
- `429` - Too Many Requests (rate limited)
- `500` - Internal Server Error
- `504` - Gateway Timeout

## Supported File Types

### Documents
- PDF (`.pdf`)
- Images: PNG, JPEG, JPG, WEBP, BMP, TIFF

### Spreadsheets
- Excel (`.xlsx`, `.xls`)
- CSV (`.csv`)
- TSV (`.tsv`)
- Google Sheets (via URL)

For the complete list, see [Landing AI Documentation](https://docs.landing.ai/ade/ade-file-types).

## Models

Landing AI offers several Document Pre-Trained Transformer (DPT) models:

### DPT-2 (Recommended)
The latest model with advanced features:
- Agentic table captioning
- Refined figure captioning
- Smarter layout detection
- Extended chunk types (logos, cards, attestations, scan codes)

**Versions:**
- `dpt-2-latest` - Always uses the newest snapshot
- `dpt-2-20250919` - September 19, 2025 snapshot
- `dpt-2-20251103` - November 3, 2025 snapshot

### DPT-1
The original model with basic parsing capabilities.

**Versions:**
- `dpt-1-latest` - Latest snapshot
- `dpt-1-20250615` - June 15, 2025 snapshot

### DPT-2 Mini
Lightweight model optimized for simple, digitally-native documents.

**Versions:**
- `DPT-2-mini-latest` - Latest snapshot
- `DPT-2-mini-20251003` - October 3, 2025 snapshot

**Note:** If you don't specify a model, the API uses `dpt-2-latest` by default.

## Rate Limits

The API has the following limits:
- Maximum page count varies by account
- Rate limiting applies to prevent abuse
- Credit usage depends on document complexity and model used

For current limits, see [Landing AI Rate Limits](https://docs.landing.ai/ade/ade-rate-limits).

## Best Practices

1. **Use Context for Cancellation**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
   defer cancel()

   result, err := client.Parse(ctx).WithFile("doc.pdf").Do()
   ```

2. **Handle Partial Failures**
   ```go
   if len(result.Metadata.FailedPages) > 0 {
       fmt.Printf("Warning: %d pages failed\n", len(result.Metadata.FailedPages))
   }
   ```

3. **Reuse Client Instances**
   ```go
   // Create once, reuse many times
   client := landingai.NewClient(apiKey)

   for _, file := range files {
       result, err := client.Parse(ctx).WithFile(file).Do()
       // ...
   }
   ```

4. **Use Specific Model Snapshots for Production**
   ```go
   // Consistent results over time
   result, err := client.Parse(ctx).
       WithFile("doc.pdf").
       WithModel("dpt-2-20250919").
       Do()
   ```

5. **Handle Rate Limits Gracefully**
   ```go
   if apiErr.IsRateLimited() {
       time.Sleep(time.Second * 5)
       // Retry with exponential backoff
   }
   ```

## Examples

See the [examples](./examples) directory for complete working examples:

- `parse_example.go` - Comprehensive examples demonstrating all features

Run the example:

```bash
export LANDINGAI_API_KEY="your-api-key"
go run examples/parse_example.go
```

## Documentation

- [Landing AI ADE Documentation](https://docs.landing.ai/ade)
- [API Reference](https://docs.landing.ai/api-reference/tools/ade-parse)
- [Chunk Types](https://docs.landing.ai/ade/ade-chunk-types)
- [Parsing Models](https://docs.landing.ai/ade/ade-parse-models)

## Requirements

- Go 1.24.0 or higher

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This SDK is distributed under the MIT License. See LICENSE file for more information.

## Support

- **Documentation**: https://docs.landing.ai/
- **Issues**: https://github.com/youssefsiam38/landingai/issues
- **Email**: support@landing.ai

## Version

Current version: v0.1.0

This version supports the ADE Parse API (`/v1/ade/parse`). Additional APIs will be added in future releases.

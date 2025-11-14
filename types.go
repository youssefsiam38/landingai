package landingai

// ChunkType represents the type of content chunk extracted from a document
type ChunkType string

const (
	ChunkTypeText        ChunkType = "text"
	ChunkTypeTable       ChunkType = "table"
	ChunkTypeMarginalia  ChunkType = "marginalia"
	ChunkTypeFigure      ChunkType = "figure"
	ChunkTypeLogo        ChunkType = "logo"
	ChunkTypeCard        ChunkType = "card"
	ChunkTypeAttestation ChunkType = "attestation"
	ChunkTypeScanCode    ChunkType = "scan_code"
)

// SplitType represents the type of document splitting
type SplitType string

const (
	SplitTypePage SplitType = "page"
)

// GroundingType represents the type of grounding information
type GroundingType string

const (
	GroundingTypeChunkLogo        GroundingType = "chunkLogo"
	GroundingTypeChunkCard        GroundingType = "chunkCard"
	GroundingTypeChunkAttestation GroundingType = "chunkAttestation"
	GroundingTypeChunkScanCode    GroundingType = "chunkScanCode"
	GroundingTypeChunkForm        GroundingType = "chunkForm"
	GroundingTypeChunkTable       GroundingType = "chunkTable"
	GroundingTypeChunkFigure      GroundingType = "chunkFigure"
	GroundingTypeChunkText        GroundingType = "chunkText"
	GroundingTypeChunkMarginalia  GroundingType = "chunkMarginalia"
	GroundingTypeChunkTitle       GroundingType = "chunkTitle"
	GroundingTypeChunkPageHeader  GroundingType = "chunkPageHeader"
	GroundingTypeChunkPageFooter  GroundingType = "chunkPageFooter"
	GroundingTypeChunkPageNumber  GroundingType = "chunkPageNumber"
	GroundingTypeChunkKeyValue    GroundingType = "chunkKeyValue"
	GroundingTypeTable            GroundingType = "table"
	GroundingTypeTableCell        GroundingType = "tableCell"
)

// ParseGroundingBox represents a bounding box in relative coordinates (0 to 1)
type ParseGroundingBox struct {
	Left   float64 `json:"left"`
	Top    float64 `json:"top"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}

// ParseGrounding represents the location of a chunk within the original document
type ParseGrounding struct {
	Box  ParseGroundingBox `json:"box"`
	Page int               `json:"page"`
}

// ParseChunk represents an extracted chunk from the document
type ParseChunk struct {
	Markdown  string         `json:"markdown"`
	Type      string         `json:"type"`
	ID        string         `json:"id"`
	Grounding ParseGrounding `json:"grounding"`
}

// ParseSplit represents a split section of the document
type ParseSplit struct {
	Class      string   `json:"class"`
	Identifier string   `json:"identifier"`
	Pages      []int    `json:"pages"`
	Markdown   string   `json:"markdown"`
	Chunks     []string `json:"chunks"`
}

// ParseResponseGrounding represents grounding information with type
type ParseResponseGrounding struct {
	Box  ParseGroundingBox `json:"box"`
	Page int               `json:"page"`
	Type GroundingType     `json:"type"`
}

// ParseMetadata contains metadata about the parsing operation
type ParseMetadata struct {
	Filename    string  `json:"filename"`
	OrgID       *string `json:"org_id"`
	PageCount   int     `json:"page_count"`
	DurationMs  int     `json:"duration_ms"`
	CreditUsage float64 `json:"credit_usage"`
	JobID       string  `json:"job_id"`
	Version     *string `json:"version"`
	FailedPages []int   `json:"failed_pages,omitempty"`
}

// ParseResponse represents the complete response from the Parse API
type ParseResponse struct {
	Markdown  string                            `json:"markdown"`
	Chunks    []ParseChunk                      `json:"chunks"`
	Splits    []ParseSplit                      `json:"splits"`
	Grounding map[string]ParseResponseGrounding `json:"grounding"`
	Metadata  ParseMetadata                     `json:"metadata"`
}

// ParseRequest represents a request to parse a document
type ParseRequest struct {
	Model       *string    `json:"model,omitempty"`
	Document    []byte     `json:"-"` // Handled as multipart file upload
	DocumentURL *string    `json:"document_url,omitempty"`
	Split       *SplitType `json:"split,omitempty"`
}

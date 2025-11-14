package landingai

// Region represents an API region
type Region string

const (
	// RegionUS represents the US region (default)
	RegionUS Region = "us"
	// RegionEU represents the EU region
	RegionEU Region = "eu"
)

const (
	baseURLUS = "https://api.va.landing.ai"
	baseURLEU = "https://api.va.eu-west-1.landing.ai"
)

// BaseURL returns the base URL for the given region
func (r Region) BaseURL() string {
	switch r {
	case RegionEU:
		return baseURLEU
	case RegionUS:
		return baseURLUS
	default:
		return baseURLUS
	}
}

// String returns the string representation of the region
func (r Region) String() string {
	return string(r)
}

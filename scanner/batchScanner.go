package scanner

type BatchScanner struct {
	BareDir string
}

type Scannable interface {
	Path() string
	Branch() string
}

func (b BatchScanner) Run(locations []Scannable) []ScanResponse {
	scannerChannel := make(chan ScanResponse)
	scanResponses := []ScanResponse{}

	for _, lo := range locations {
		scanner := NewScanner(lo, b.BareDir, scannerChannel)
		go scanner.ScanLastModified()
	}

	for i := 0; i < len(locations); i++ {
		scanResponses = append(scanResponses, <-scannerChannel)
	}

	return scanResponses
}

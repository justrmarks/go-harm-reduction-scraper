package db

type Entry struct {
	SamplePhoto    []byte         `json:"samplePhoto"`
	SampleName     string         `json:"sampleName"`
	SubstanceRatio map[string]int `json:"substanceRatio"`
	DatePublished  string         `json:"datePublished"`
	DateTested     string         `json:"dateTested"`
	Location       string         `json:"location"`
	SampleSize     string         `json:"sampleSize"`
	DataSource     string         `json:"dataSource"`
}

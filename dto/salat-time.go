package dto

type SalatTimeResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
}

type Timings struct {
	Fajr       string `json:"Fajr"`
	Sunrise    string `json:"Sunrise"`
	Dhuhr      string `json:"Dhuhr"`
	Asr        string `json:"Asr"`
	Sunset     string `json:"Sunset"`
	Maghrib    string `json:"Maghrib"`
	Isha       string `json:"Isha"`
	Midnight   string `json:"Midnight"`
	FirstThird string `json:"Firstthird"`
	LastThird  string `json:"Lastthird"`
}

type Date struct {
	Readable  string         `json:"readable"`
	Timestamp string         `json:"timestamp"`
	Hijri     map[string]any `json:"hijri"`
	Gregorian map[string]any `json:"gregorian"`
}

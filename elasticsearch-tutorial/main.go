package main

// https://www.freecodecamp.org/news/go-elasticsearch/
// alternative: https://outcrawl.com/go-elastic-search-service
type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func main() {
	IndexDocument()
	QueryDocuments()
}

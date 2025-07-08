package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var USERAGENTS []string = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.7.5",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.9",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.3",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/24.0.2.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.2.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 115Browser/5.1.6",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
}

func Agent() string {
	if len(USERAGENTS) == 0 {
		return ""
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random index within the bounds of the slice.
	randomIndex := r.Intn(len(USERAGENTS))

	// Return the string at the random index.
	return USERAGENTS[randomIndex]

}

type Course struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func ReadCoursesFromJSONFile(filePath string) ([]Course, error) {
	// Read the content of the file into a byte slice
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	var courses []Course
	// Unmarshal the byte slice into the Go struct slice
	err = json.Unmarshal(jsonData, &courses)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from file '%s': %w", filePath, err)
	}

	return courses, nil
}

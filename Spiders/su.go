package main

import (
	utils "GoSpiders/Utils" // Assuming this path is correct for your project structure
	"encoding/json"
	"fmt"
	"log" // Added for better error handling
	"os"

	"github.com/gocolly/colly/v2"
)

var BASEURLS []string = []string{
	"https://www.standyou.com/study-abroad/",
	"https://www.standyou.com/country/",
	"https://www.standyou.com/universities/",
	"https://www.standyou.com/course-program/",
}

// CourseCountries stores course name with list of countries
type StudyAbraod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func Courses() {
	// Create collector for the homepage
	c := colly.NewCollector(
		colly.AllowedDomains("www.standyou.com"),
		colly.UserAgent(utils.Agent()),
	)

	var courses []StudyAbraod

	// OnHTML for the <ul> element with a specific class
	c.OnHTML("ul.selectCourses ", func(e *colly.HTMLElement) {
		fmt.Printf("Found UL with class: %s\n", e.Attr("class"))

		// Iterate over <li> elements within this <ul>
		// CORRECTED CSS SELECTOR: combine classes with no spaces
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			sa := StudyAbraod{} // Create a new StudyAbraod for each li
			// Access attributes or text content of the <li>
			sa.Name = el.Text
			// The 'href' attribute is typically on an <a> tag *inside* the <li>
			// You need to find the <a> tag within the current <li> and get its href.
			sa.URL = el.ChildAttr("a", "href")

			fmt.Printf("  Found LI: Name: '%s', URL: '%s'\n", sa.Name, sa.URL)

			courses = append(courses, sa)
			fmt.Println(sa)
		})
	})

	// Add an error handler for debugging
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error visiting %s: %v\n", r.Request.URL.String(), err)
	})

	// Start scraping
	err := c.Visit(BASEURLS[0])
	if err != nil {
		fmt.Println("Error visiting:", err)
		return
	}

	// Save output to JSON file
	file, err := json.MarshalIndent(courses, "", " ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	err = os.WriteFile("courses.json", file, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println("✅ Scraping completed. Output saved to courses.json")
}

type Countries struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	c := colly.NewCollector(
		// colly.AllowedDomains("www.standyou.com"),
		colly.UserAgent(utils.Agent()),
	)
	countries := []Countries{}
	jsonD, _ := utils.ReadCoursesFromJSONFile("D:/Projects/go/GoSpiders/Spiders/courses.json")
	for i, courseLinks := range jsonD {

		fmt.Printf("Visiting\n: %d", i+1)
		c.OnHTML("ul.selectCourses", func(e *colly.HTMLElement) {
			fmt.Printf("Found UL with class: %s\n", e.Attr("class"))

			// Iterate over <li> elements within this <ul>
			// CORRECTED CSS SELECTOR: combine classes with no spaces
			e.ForEach("li", func(_ int, el *colly.HTMLElement) {
				country := Countries{} // Create a new StudyAbraod for each li
				// Access attributes or text content of the <li>
				country.Name = el.Text
				// The 'href' attribute is typically on an <a> tag *inside* the <li>
				// You need to find the <a> tag within the current <li> and get its href.
				country.URL = el.ChildAttr("a", "href")

				fmt.Printf("  Found LI: Name: '%s', URL: '%s'\n", country.Name, country.URL)

				countries = append(countries, country)
				fmt.Println(country)
			})
		})

		// Add an error handler for debugging
		c.OnError(func(r *colly.Response, err error) {
			log.Printf("Error visiting %s: %v\n", r.Request.URL.String(), err)
		})
		c.Visit(courseLinks.URL)
		file, err := json.MarshalIndent(countries, "", " ")
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}
		err = os.WriteFile("Countries.json", file, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	PartOfSpeech string
	Meanings     []string
}

func main() {
	// Step 1: Read the CSV file
	file, err := os.Open("dictionary.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Step 2: Store data in a map
	dictionary := make(map[string]Entry)
	for _, record := range records {
		if len(record) < 3 {
			continue // Skip malformed rows
		}
		word := strings.ToTitle(record[0]) // Capitalize the first letter
		partOfSpeech := record[1]
		meaning := record[2]

		// If the word already exists, append to its meanings
		if entry, found := dictionary[word]; found {
			entry.Meanings = append(entry.Meanings, meaning)
			dictionary[word] = entry
		} else {
			dictionary[word] = Entry{PartOfSpeech: partOfSpeech, Meanings: []string{meaning}}
		}
	}

	// Step 3: Search for a keyword
	var keyword string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&keyword)

	// Capitalize the entered keyword
	keyword = strings.ToTitle(keyword)

	if entry, found := dictionary[keyword]; found {
		fmt.Printf("Word: %s\nPart of Speech: %s\nMeanings:\n", keyword, entry.PartOfSpeech)
		for _, meaning := range entry.Meanings {
			fmt.Printf(" - %s\n", meaning)
		}
	} else {
		fmt.Println("Word not found.")
	}
}

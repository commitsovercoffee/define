package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"strings"
)

type Meaning struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Meanings     []string `json:"meanings"`
}

type Entry struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

var partOfSpeechDescriptions = map[string]string{
	"n.":     "Noun",
	"v.":     "Verb",
	"v. t.":  "Transitive Verb",
	"v. i.":  "Intransitive Verb",
	"a.":     "Adjective",
	"adv.":   "Adverb",
	"n. pl.": "Plural Noun",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-dictionary <word>")
		return
	}

	word := os.Args[1]

	// Load dictionary from CSV
	entries, err := loadDictionary("dictionary.csv")
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return
	}

	// Find and print the meanings of the word
	entry := findEntry(entries, word)
	if entry != nil {
		jsonResponse, err := json.MarshalIndent(entry, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			return
		}
		fmt.Println(string(jsonResponse))
	} else {
		fmt.Printf("'%s' not found in the dictionary.\n", word)
	}
}

// loadDictionary reads the CSV file and returns a slice of Entry structs
func loadDictionary(filename string) ([]Entry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	entriesMap := make(map[string]map[string][]string)

	for _, record := range records {
		if len(record) == 3 {
			word := record[0]
			partOfSpeech := record[1]
			meaning := formatMeaning(record[2])

			if _, exists := entriesMap[word]; !exists {
				entriesMap[word] = make(map[string][]string)
			}
			entriesMap[word][partOfSpeech] = append(entriesMap[word][partOfSpeech], meaning)
		}
	}

	var entries []Entry
	for word, parts := range entriesMap {
		meanings := []Meaning{}
		for part, meaningsList := range parts {
			descriptivePart := partOfSpeechDescriptions[part]
			meanings = append(meanings, Meaning{
				PartOfSpeech: descriptivePart,
				Meanings:     meaningsList,
			})
		}
		entries = append(entries, Entry{
			Word:     word,
			Meanings: meanings,
		})
	}

	return entries, nil
}

// formatMeaning cleans up the meaning string
func formatMeaning(meaning string) string {
	// Decode HTML entities and handle other formatting
	decoded := html.UnescapeString(meaning)
	return strings.ReplaceAll(decoded, "&", "and")
}

// findEntry searches for a word in the dictionary entries and returns a structured Entry
func findEntry(entries []Entry, word string) *Entry {
	for _, entry := range entries {
		if strings.EqualFold(entry.Word, word) {
			return &entry
		}
	}
	return nil
}

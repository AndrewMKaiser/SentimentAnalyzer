package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)





func buildSocialSentimentTable(filename string) map[string]float32 {

	socialSentFile, err := os.ReadFile(filename) // Reads file into a byte slice
	check(err) // Checks validity of filename
	socialSentString := string(socialSentFile) // Converts byte slice into a string
	socialSentCsvReader := csv.NewReader(strings.NewReader(socialSentString)) // Creates a csv reader (standard library)
	socialSentTable := make(map[string]float32) // Creates a table mapping each word to its sentiment score

	socialSentCsvReader.Read() // Skips first record in the table ("word,Sentiment scores,,https://nlp.stanford.edu/projects/socialsent/")

	for {

		record, err := socialSentCsvReader.Read() // Reads in each record in the table

		if err == io.EOF { // Returns socialSentTable once the end-of-file has been reached
			return socialSentTable
		}

		check(err) // Checks for errors when reading in the data

		scoreFloat, _ := strconv.ParseFloat(record[1], 32) // Converts each score from a float to a string

		socialSentTable[record[0]] = float32(scoreFloat) // Maps the score to the word and stores it in the table
	}


}

func getSocialSentimentScore(filename string) float32 {
	var avgScore float32
	reviewFile, err := os.ReadFile(filename) // reviewFile is a byte slice of the file data
	check(err) // Checks for filename validity
	reviewText := removeNonAlphabetical(string(reviewFile)) // Stores the file data as a string containing only words
	reviewTextReader := strings.NewReader(reviewText)
	
	return avgScore
}

func removeNonAlphabetical(input string) string {
	var builder strings.Builder // Initializes string builder

	for _, r := range input { // Scans each character in the string
		if unicode.IsLetter(r) || r == ' ' { // Checks if the character is a letter or a whitespace
			builder.WriteRune(r) // If so, write it to the builder
		} // Skips all non-letter, non-whitespace characters
	}

	return builder.String() // Returns the string containing only words
}

func check(e error) { // From https://gobyexample.com/reading-files
    if e != nil {
        log.Fatalf("Error found %q\n", e)
    }
}

/*
func buildSocialSentimentScores(table map[string]float32, tableSize int) []float32 {

	scoreSlice := make([]float32, tableSize) // Creates a slice with length tableSize (setting length = capacity = tableSize for performance)

	for _, score := range table { // Iterates through all the scores in the table
		scoreSlice = append(scoreSlice, score) // Appends the scores to scoreSlice
	}

	return scoreSlice
}
*/

func main() {
	filename := "socialsent.csv"
	
	SocialSentimentScores := buildSocialSentimentTable(filename)

	fmt.Print(SocialSentimentScores)
	

}

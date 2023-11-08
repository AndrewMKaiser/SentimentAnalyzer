package main

// packages
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
	if err != nil { // Checks validity of filename
		fmt.Println("You've entered an invalid filename. Ensure the file is in the directory of the executable and that the filename was typed correctly.")
		os.Exit(1)
	}
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

func getSocialSentimentScore(filename string, socialSentTable *map[string]float32) (float32, int) {

	var finalScore float32 // Declares variable to store the overall score of the file

	reviewFile, err := os.ReadFile(filename) // reviewFile is a byte slice of the file data
	check(err) // Checks for filename validity
	reviewText := removeNonAlphabetical(string(reviewFile)) // Stores the file data as a string containing only words and whitespace
	reviewText = strings.ReplaceAll(reviewText, "\n", " ")
	reviewTextSlice := strings.Split(strings.ToLower(reviewText), " ") // Splits reviewText into a slice of lowercase words (lowercase because socialsent.csv contains only lowercase words)
	
	fmt.Println()
	fmt.Println("\033[1m[word: current_score, accumulated_score]\033[0m") // Format message: bolded using ANSI escape chars
	for _, word := range reviewTextSlice { // Iterates through each word in the file
		sentScore, inMap := (*socialSentTable)[word] // sentScore is the value in the table corresponding to key 'word,' and inMap is a boolean value (true if word is in the table, false if not)
		if inMap {
			finalScore += sentScore // If the word is in the table, extract its sentiment score and add it to finalScore
			fmt.Printf("%s: %.2f, %.2f\n", word, sentScore, finalScore) // Prints out the word, current score, and accumulated score in correct format
		}
		// Otherwise, move to the next word
	}
	
	starRating := getStarRating(finalScore) // Stores the star rating

	fmt.Println()
	fmt.Printf("\033[1m%s score: %.2f\n\033[0m", filename, finalScore) // Score message: bolded using ANSI escape chars
	fmt.Printf("\033[1m%s star rating: %d\n\033[0m", filename, starRating) // Stars message: bolded using ANSI escape chars

	return finalScore, starRating // Return the result of the sum of all sentiment scores in the file
}

func getStarRating(ratingNum float32) int {

	switch {
	case ratingNum >= 5:
		return 5
	case ratingNum >= 1:
		return 4
	case ratingNum >= -1:
		return 3
	case ratingNum >= -5:
		return 2
	case ratingNum < -5:
		return 1
	default:
		fmt.Println("Invalid rating")
		return 0
	}
}

func removeNonAlphabetical(input string) string {

	var builder strings.Builder // Initializes string builder

	for _, r := range input { // Scans each character in the string
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' || r == '\n' { // Checks if the character is a letter, a number, a whitespace, or a newline
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

func main() {

	var reviewFilename string

	args := os.Args // Command-line args

	switch len(args) {
	case 1:
		reviewFilename = "review.txt" // Default review filename
	case 2:
		reviewFilename = args[1] // User-entered review filename
	default:
		fmt.Println("Usage is: \"./socialsent <review_filename>\"") // Reminds user of proper usage
		os.Exit(1) // Exits
	}
	
	tableFilename := "socialsent.csv"
	SocialSentimentScores := buildSocialSentimentTable(tableFilename) // Builds the sentiment table
	getSocialSentimentScore(reviewFilename, &SocialSentimentScores) // Retrieves and prints sentiment data from review file

}

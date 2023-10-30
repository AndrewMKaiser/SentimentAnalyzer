# Homework 4 - CSCI 3415

- **Name:** *Andrew Kaiser*
- **Student ID:** *110431932*
- **Class:** *CSCI 3414 - Principles of Programming Languages*
- **HW#:** *4*
- **Due Date:** *November 7, 2023*

## Q1) Leftmost Derivation and Parse Tree

PDF file is in submission.

## Q2) Social Sentiment Scores

### How to run

1. [Navigate to this REPL.](https://replit.com/join/teqpaondka-andrewkaiser5) The source files are also included in the submission in case the link does not work. The code is working in REPL.
2. In the shell, type: `go build socialsent.go`
3. After the binary is compiled, you may run it using the following format: `./socialsent <review_filename>` for a specific filename, or `./socialsent` to run it on the default `review.txt` file.

### **Social Sentiment Report**

### Abstract

Go is a strongly and statically typed, compiled language, popularly used in backend web development, and for sys-admins and DevOps engineers. Its performant and procedural nature make it a great language to learn for programmers experienced in similar languages. Go's simple, elegant syntax makes the language an easy pick-up for developers used to syntactically C-like languages like Java or Python. The language is noted for its unique approach to concurrency through the use of `goroutines` and `channels`. Its standard package library is vast and allows developers to access everything from low-level syscalls to parsing arbitrary data formats like CSV files. Overall, Go is an expansive, performant, and simple language great for a multitude of applications.

---

### Approach

To formulate an overall sentiment score for a given file, I followed a straightforward approach that utilized Go's `map` data type (dictionary ADT) to assign each word in the file to a sentiment score. I then summed the scores of each word to calculate the overall score, and rated it using a star ranking system. Below are the steps taken to implement this approach:

1. **Reading and Parsing Data from CSV**

    I made use of Go's standard libraries like `os` and `encoding/csv` to read files and parse CSV data.
  
2. **String Conversion**

    I parsed strings from the review file using the `strings` package, converted them to appropriate data types (like converting sentiment scores stored as strings to floats), and then stored them in a map (dictionary ADT).

3. **Sentiment Analysis**

    For the sentiment analysis, the given review text was stripped by removing non-letter characters from the text in the review file. Each word in the file is then looked up in the sentiment table, and if it exists in the sentiment table, its score is added to the final sentiment score.

4. **Rating System**

    Based on the accumulated sentiment score, I designed a simple system to determine a star rating for the review using Go's multi-way selection `switch` statement.

---

### Unique Language Features

1. **Builder Types**

    I utilized Go's strings.Builder to construct one new string, rune by rune, which only consisted of letters, whitespaces, and newlines from the review file. Go's Builder type avoids creating multiple temporary string objects, which enhances performance and just feels natural from a programming perspective.

### Features I Liked

1. **Explicit Error Handling**

    Unlike languages such as C++, Go returns errors as normal return values (in contrast to blocking exceptions), which makes error handling straightforward and explicit. No need for any verbose `try... catch...` statements.

2. **Simplicity**

    The simplicity of Go makes it a joy to work with. Features like garbage collection mean I don't have to stress about memory issues.

3. **Rich Standard Library**

    The wide array of functionalities available in Go's standard library, from C-style string conversions to CSV file handling, is flexible and well-documented. Incorporating these packages into this assignment required little effort because of the great developer community and documentation of each package.

4. **Typing**

    Go's typing methods allow for inferred and explicit, static typing. When I needed to declare a variable before using it, I could easily type it explicitly using the syntax `var <var_name> <type>` or I could initialize a value upon declaration using type inference with the simpler syntax `<var_name> := <typed_expr>`. Complex datatypes like slices and maps are easy to understand with syntax like `<map_name> = make(map[<key_type>]<valuetype>)`.

5. **gofmt**

    `gofmt` is a command that formats your Go code to make it more consistent and readable. This way, style remains consistent across all my Go projects.

---

### Other Special Notes About the Language

- **Package Management**

    Go's `mod` files make dependency tracking a breeze by maintaining a `go.mod` file containing the version of Go the code runs on, along with all external dependencies used for the project.

- **Learning Resources**

    One of the strengths of Go is its community. The language has a vast number of tutorials to implement different programs and services, giving learners the ability to learn the intricacies of the language through applicable projects.

---

### Conclusions

Creating a sentiment analysis program in Go has demonstrated its outstanding versatility, efficiency, and simplicity. The language's many features allowing me to quickly and efficiently parse and analyze data using readable syntax make it a great choice for future projects in the domain of data analysis. Go's expansive standard library and its associated simplicity allowed me to complete my task in far less time than it would have in similar procedural languages like C. Tools like `gofmt` and the project's associated `go.mod` file meant that I could spend less time worrying about code formatting and documentation, and more time actually writing code. Go's tools and packages had excellent documentation which drastically improved my research time compared to languages with difficult-to-understand documentation like many UNIX-based C libraries. Overall, my experience with Go has been fantastic, and I fully intend to use it in future projects, and in other school assignments.

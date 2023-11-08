# Go Social Sentiment Analyzer

A command-line program that parses text and uses social sentiment data gathered from Stanford researchers to determine the overall reception of the text. The program will assign social sentiment scores to each word in a given text file, total the score of each word in the file, and give an overall star rating from 1 to 5 describing the positive or negative connotation of the text as a whole.

## Why use this program?

### Business Intelligence
- **Product Feedback:** Utilize sentiment analysis to understand customer reactions to products for enhancement and innovation.
- **Market Research:** Leverage social sentiment to comprehend consumer feelings towards market trends for strategic planning.

### Brand Management
- **Reputation Monitoring:** Monitor brand sentiment to manage reputation and address negative feedback swiftly.
- **Crisis Management:** Utilize sentiment analysis for real-time public reaction tracking during crises.

### Academic Research
- **Social Science Studies:** Explore human behavior and cultural patterns through sentiment analysis.
- **Psychological Analysis:** Assess community mood and stress for public health implications.

## How to install and run

1. [Install Go, version 1.21 or higher](https://go.dev/doc/install)
2. [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
3. Navigate to your desired directory and run `git clone https://github.com/AndrewMKaiser/GoSocialSentimentAnalyzer.git` to clone the project to your system.
4. In the project directory, run `go build socialsent.go` to build the `socialsent` executable.
5. After the executable is compiled, you may run it using the following format: 
    - **Mac/Linux** `./socialsent <review_filename>` for a specific filename, or `./socialsent` to run it on the default `review.txt`
    - **Windows** `.\socialsent <review_filename>` or `.\socialsent` to run it on the default `review.txt`.

### Citation

William L. Hamilton, Kevin Clark, Jure Leskovec, and Dan Jurafsky. [Inducing Domain-Specific Sentiment Lexicons from Unlabeled Corpora](https://arxiv.org/abs/1606.02820). ArXiv preprint (arxiv:1606.02820). 2016.
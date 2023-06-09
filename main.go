package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func shuffleSlice(slice [][]string) [][]string {

	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func newRandomSeed() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

// call -filename -shuffle -timer
func main() {

	quizFilename := "quiz.csv"
	randomize := false
	timerSeconds := 30

	for i, arg := range os.Args {
		switch arg {
		case "-f":
			quizFilename = os.Args[i+1]
		case "-s":
			randomize = true
		case "-t":
			timerSeconds, _ = strconv.Atoi(os.Args[i+1])
		case "-help":
			fmt.Println("-f <filepath> to specify a different .csv file to populate the quiz.")
			fmt.Println("-s to specify that the quiz question order be randomized.")
			fmt.Println("-t <time in seconds> to specify a non-default (30 seconds) timer for the quiz.")
			os.Exit(0)
		}
	}

	quiz := NewQuiz()
	quiz.populateQuestionsandAnswers(quizFilename)
	if randomize {
		quiz.shuffleQuestionOrder()
	}
	quiz.setTotalQuestionsCount()
	fmt.Println("Press enter when ready to begin!")
	fmt.Scanln()
	quiz.startTimer(timerSeconds)
	quiz.runQuiz()
	quiz.printFormattedScore()

}

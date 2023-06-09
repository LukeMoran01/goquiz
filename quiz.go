package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Quiz struct {
	questionsAnswers [][]string
	totalQuestions   int
	correctAnswers   int
}

func NewQuiz() Quiz {
	return Quiz{}
}

func (quiz *Quiz) populateQuestionsandAnswers(quizFilename string) {
	quizFile, _ := os.Open(quizFilename)
	csvReader := csv.NewReader(quizFile)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		quiz.questionsAnswers = append(quiz.questionsAnswers, record)
	}
}

func (quiz *Quiz) setTotalQuestionsCount() {
	quiz.totalQuestions = len(quiz.questionsAnswers)
}

func (quiz *Quiz) incrementCorrectAnswers() {
	quiz.correctAnswers++
}

func (quiz *Quiz) printFormattedScore() {
	fmt.Printf("%v/%v correct answers.\n", quiz.correctAnswers, quiz.totalQuestions)
}

func (quiz *Quiz) runQuiz() {
	for _, question := range quiz.questionsAnswers {
		var userAnswer string
		fmt.Println(question[0])
		fmt.Scanln(&userAnswer)
		if userAnswer == question[1] {
			quiz.incrementCorrectAnswers()
		}
	}
}

func (quiz *Quiz) startTimer(timeInSeconds int) {
	go func(quiz *Quiz) {
		timer := time.NewTimer(time.Duration(timeInSeconds) * time.Second)
		<-timer.C
		quiz.printFormattedScore()
		fmt.Println("Time limit reached. Exiting...")
		os.Exit(0)
	}(quiz)
}

func (quiz *Quiz) shuffleQuestionOrder() {
	newRandomSeed()
	shuffleSlice(quiz.questionsAnswers)
}

package main

import "fmt"

// I need a counter for the next element

// I need a structure to represent a question
// animal name, question, yes result, no result 

// I need a function to ask a question, and based on the answer navigate to the next question

// I need a function to add a question to the tree

// I need a function to save the tree to a file

// I need a function to load the tree from a file

// I need a function to play the game

// I need a function to reset the game

// I need a function to exit the game

type question struct {
	Id int 
	Name string
	Question string
	Yes int
	No int
}

var nextQuestion int = 1

var allQuestions map[int]question

func incrimentNextQuestion() int {
	nextQuestion++
	return nextQuestion
}

func initialQuestionList() {
	q1 := question{Id: 1, Name: "Elephant", Question: "Is it a mamal?", Yes: 0, No: 2}
	q2 := question{Id: 2, Name: "Shark"}

	var questions = map[int]question{1: q1, 2: q2}
	allQuestions = questions
}

func finalGuess(question question) {
	fmt.Println("Is it a", question.Name)
	var guess string
	fmt.Scanln(&guess)
	if guess == "yes" {
		fmt.Println("Woo Hoo, I win")
	} else {
		fmt.Println("Congratulations, you win")
	}
}

func gameRound(startingQuestion int) {
	var quest = allQuestions[startingQuestion]
	if quest.Question == "" {
		finalGuess(quest)
		return
	}
	fmt.Println(quest.Question)
	var guess string
	fmt.Scanln(&guess)
	if guess == "yes" {
		if quest.Yes == 0 {
			finalGuess(quest)
		} else {
			gameRound(quest.Yes)
		}
	} else {
		if quest.No == 0 {
			finalGuess(quest)
		} else {
			gameRound(quest.No)
		}
	}
}

func main() {
	initialQuestionList()
	gameRound(1)
//	fmt.Println(questions)
//	fmt.Println(guess)
}
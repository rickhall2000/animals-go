package main

import "fmt"

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

// I got this wrong, first I should ask a question, but have no animal
// then each answer should be a different animal
// so mamal -> 1 elephant or 2 shark

func initialQuestionList() {
	q1 := question{Id: 1, Name: "Elephant", Question: "Is it a mamal?", Yes: 0, No: 2}
	q2 := question{Id: incrimentNextQuestion(), Name: "Shark"}
    

	var questions = map[int]question{1: q1, 2: q2}
	allQuestions = questions
}

func growGame(quest question){
	var answer string
	fmt.Println("What is the name of the animal?")
	fmt.Scanln(&answer)

	var newQuestion question
	newQuestion.Name = answer
	newQuestion.Id = incrimentNextQuestion()

	var newQuestionText string
	fmt.Println("What is a question that would distinguish between a %s and a %s", quest.Name, newQuestionText)
	fmt.Scanln(&newQuestionText)

	quest.Question = newQuestionText

	var newAnswer string 
	fmt.Println("What is the right answer for a %s?", answer)


	if newAnswer == "yes" {
		quest.Yes = newQuestion.Id
	} else {
		quest.No = newQuestion.Id
	}

	allQuestions[newQuestion.Id] = newQuestion
}

func finalGuess(question question) {
	fmt.Println("Is it a", question.Name)
	var guess string
	fmt.Scanln(&guess)
	if guess == "yes" {
		fmt.Println("Woo Hoo, I win")
	} else {
		fmt.Println("Congratulations, you win")
		growGame(question)
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
	fmt.Println(allQuestions)
//	fmt.Println(questions)
//	fmt.Println(guess)
}
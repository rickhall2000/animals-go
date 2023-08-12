package main

import "fmt"
import "bufio"
import "os"
import "strings"

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
	q1 := question{Id: 1, Question: "Is it a mamal?", Yes: 2, No: 3} 
	q2 := question{Id: incrimentNextQuestion(), Name: "Elephant"}
	q3 := question{Id: incrimentNextQuestion(), Name: "Shark"}
    
	var questions = map[int]question{1: q1, 2: q2, 3: q3}
	allQuestions = questions
}

func growGame(quest question){
	
	var oldLeaf question 
	oldLeaf.Name = quest.Name
	oldLeaf.Id = incrimentNextQuestion()

	reader := bufio.NewReader(os.Stdin)

	var answer string
	fmt.Println("What is the name of the animal?")
	text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
	answer = text

	var newLeaf question
	newLeaf.Name = answer
	newLeaf.Id = incrimentNextQuestion()

	var newQuestionText string
	fmt.Printf("What is a question that would distinguish between a %s and a %s?\n", oldLeaf.Name, newLeaf.Name)
	text, _ = reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
	newQuestionText = text
	
	quest.Question = newQuestionText
	quest.Name = ""

	var newAnswer string 
	fmt.Printf("What is the right answer for a %s?\n", answer)
	fmt.Scanln(&newAnswer)
	

	if newAnswer == "yes" {
		quest.Yes = newLeaf.Id
		quest.No = oldLeaf.Id
	} else {
		quest.No = newLeaf.Id
		quest.Yes = oldLeaf.Id
	}

	allQuestions[newLeaf.Id] = newLeaf
	allQuestions[oldLeaf.Id] = oldLeaf
	allQuestions[quest.Id] = quest
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
	for {
		gameRound(1)
		fmt.Println(allQuestions)
		fmt.Println("Play again?")
		var ans string
		fmt.Scanln(&ans)
		if ans != "yes" {
			break;
		}
	} 
	fmt.Println("Good bye")
}
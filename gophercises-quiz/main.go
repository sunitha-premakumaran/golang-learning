package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem Structure for quiz
type Problem struct {
	q string
	a string
}

func main() {
	array := readFile("./problem.csv")
	problems := generateProblem(array)
	score := beginQuiz(problems)
	fmt.Printf("Your Total score %d out of %d\n", score, len(problems))
	os.Exit(1)
}

func readFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed with err")
	}
	reader := csv.NewReader(file)
	lines, _ := reader.ReadAll()
	return lines
}

func beginQuiz(problems []Problem) int {
	var score int
	timer := time.NewTimer(time.Duration(30) * time.Second)
	answerChannel := make(chan string)
	for i, problem := range problems {
		fmt.Printf("%d.What is %s\n", i+1, problem.q)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()

		select {
		case <-timer.C:
			return score
		case answer := <-answerChannel:
			if strings.TrimSpace(answer) != strings.TrimSpace(problem.a) {
				fmt.Printf("Incorrect!.Expected %s Got %s\n", problem.a, answer)
			} else {
				score++
			}
		}
	}

	return score
}

func generateProblem(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, item := range lines {
		problems[i].q = item[0]
		problems[i].a = item[1]
	}
	return problems
}

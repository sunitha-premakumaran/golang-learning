package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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
	fmt.Printf("Your Total score %d\n", score)
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
	for _, problem := range problems {
		fmt.Printf("What is %s\n", problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if strings.TrimSpace(answer) != strings.TrimSpace(problem.a) {
			fmt.Printf("Incorrect!!!.Expected %s Got %s\n", problem.a, answer)
		} else {
			score++
			fmt.Println("Correct!!!!!")
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

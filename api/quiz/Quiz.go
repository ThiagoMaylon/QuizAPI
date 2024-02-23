package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Questions struct {
	Question     string   `json:"question"`
	Answers      []string `json:"answers"`
	Right_answer int      `json:"right_answer"`
}

type Topics struct {
	Title     string      `json:"title"`
	Content   string      `json:"contente"`
	Questions []Questions `json:"questions"`
}
type Quiz struct {
	Topics []Topics `json:"topics"`
}

func GetQuiz(filename string) (Quiz, error) {
	file, err := os.Open(filename)

	if err != nil {
		return Quiz{}, err
	}

	defer file.Close()

	var (
		quiz         Quiz
		currentTopic *Topics
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Topico: ") {
			currentTopic = &Topics{
				Title: line[len("Topico: "):],
			}
		} else if strings.HasPrefix(line, "Conteudo: ") {
			if currentTopic != nil {
				content := line[len("Conteudo: "):]
				currentTopic.Content += content
			}
		} else if strings.HasPrefix(line, "Pergunta: ") {

			if currentTopic != nil {
				ques := Questions{}
				ques.Question = line[len("Pergunta: "):]

				for i := 0; i < 3; i++ {
					scanner.Scan()
					resp := scanner.Text()
					ques.Answers = append(ques.Answers, resp)
				}

				scanner.Scan()
				respCorreta, err := strconv.Atoi(scanner.Text())

				if err != nil {
					return Quiz{}, fmt.Errorf("erro ao converter resposta correta para número: %v", err)
				}
				ques.Right_answer = respCorreta - 1
				currentTopic.Questions = append(currentTopic.Questions, ques)
				quiz.Topics = append(quiz.Topics, *currentTopic)

				fmt.Println(quiz)

			}
		}

	}

	if err := scanner.Err(); err != nil {
		return Quiz{}, err
	}

	return quiz, nil
}

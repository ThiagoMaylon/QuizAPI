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
	Tip          string   `json:"tip"`
}

type Topics struct {
	Id        int         `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
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

	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Topico: ") {
			if currentTopic != nil {
				quiz.Topics = append(quiz.Topics, *currentTopic)
			}
			currentTopic = &Topics{
				Title: line[len("Topico: "):],
				Id:    id,
			}
			id++
		} else if strings.HasPrefix(line, "Conteudo: ") {
			if currentTopic != nil {
				content := line[len("Conteudo: "):]
				currentTopic.Content += content
			}
		} else if strings.HasPrefix(line, "Pergunta: ") {
			if currentTopic != nil {
				var ques Questions
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
			}
		} else if strings.HasPrefix(line, "Dica: ") {
			if currentTopic != nil && len(currentTopic.Questions) > 0 {
				tip := line[len("Dica: "):]
				currentTopic.Questions[len(currentTopic.Questions)-1].Tip = tip
			}
		}
		
	}

	if currentTopic != nil {
		quiz.Topics = append(quiz.Topics, *currentTopic)
	}

	if err := scanner.Err(); err != nil {
		return Quiz{}, err
	}

	return quiz, nil
}

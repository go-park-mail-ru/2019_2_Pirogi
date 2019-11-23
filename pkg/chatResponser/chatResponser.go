package chatResponser

import "regexp"

type Patterns struct {
	QuestionPattern *regexp.Regexp
	Answer          string
}

var Answers = map[string]Patterns{
	"greeting": {
		QuestionPattern: regexp.MustCompile(`(?i)прив\w*|здравствуй\w*|добр\w*`),
		Answer:          "Здравствуйте! Чем могу помочь?",
	},
	"introduction": {
		QuestionPattern: regexp.MustCompile(`(?i)я \w*|я - \w*|я-\w*|меня зовут \w*`),
		Answer:          "Очень приятно, я - Ваш ассистент по cinsear. Чем могу помочь?",
	},
	"howareyou": {
		QuestionPattern: regexp.MustCompile(`(?i)как дела`),
		Answer:          "Великолепно. Чем могу помочь?",
	},
	"goodbye": {
		QuestionPattern: regexp.MustCompile(`(?i)пока|до свидания|всего доброго`),
		Answer:          "Всего доброго!",
	},
}

// Возвращает ответ и флаг, отвечающий за то, что ответ исчерпывающий, т.е. не требует оператора
func Answer(question string) (string, bool) {
	for k := range Answers {
		if Answers[k].QuestionPattern.MatchString(question) {
			return Answers[k].Answer, true
		}
	}
	if regexp.MustCompile(`(?i).*?`).MatchString(question) {
		return "Ваш вопрос в обработке. Подождите, пока с Вами свяжется наш специалист.", false
	}
	return "Я вас немного не понимаю. Чем могу помочь?", true
}

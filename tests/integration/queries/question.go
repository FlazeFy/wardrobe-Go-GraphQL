package queries

const CreateQuestionMutation = `
	mutation CreateQuestion($input: CreateQuestionInput!) {
		createQuestion(input: $input) {
			id
			question
			answer
			is_show
			createdAt
		}
	}
`

type CreateQuestionData struct {
	CreateQuestion struct {
		ID        string  `json:"id"`
		Question  string  `json:"question"`
		Answer    *string `json:"answer"`
		IsShow    bool    `json:"is_show"`
		CreatedAt string  `json:"createdAt"`
	} `json:"createQuestion"`
}

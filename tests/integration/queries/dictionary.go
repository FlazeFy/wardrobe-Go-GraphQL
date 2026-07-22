package queries

const CreateDictionaryMutation = `
	mutation CreateDictionary($input: CreateDictionaryInput!) {
		createDictionary(input: $input) {
			id
			dictionaryType
			dictionaryName
			createdAt
		}
	}
`

type CreateDictionaryData struct {
	CreateDictionary struct {
		ID             string `json:"id"`
		DictionaryType string `json:"dictionaryType"`
		DictionaryName string `json:"dictionaryName"`
		CreatedAt      string `json:"createdAt"`
	} `json:"createDictionary"`
}

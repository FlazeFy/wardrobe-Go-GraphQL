package integration

import (
	"strings"
	"testing"

	"wardrobe-graphql/tests/integration/client"
	"wardrobe-graphql/tests/integration/queries"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateQuestion_ValidData(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateQuestionMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"question": "lorem ipsum?",
			},
		},
	})
	require.NoError(t, err)

	// Validate empty error message
	require.False(t, resp.HasErrors(), resp.FirstErrorMessage())

	var data queries.CreateQuestionData
	require.NoError(t, resp.Unmarshal(&data))

	// Validate response
	assert.NotEmpty(t, data.CreateQuestion.ID)
	assert.Equal(t, "lorem ipsum?", data.CreateQuestion.Question)
	assert.Nil(t, data.CreateQuestion.Answer)
	assert.False(t, data.CreateQuestion.IsShow)
	assert.NotEmpty(t, data.CreateQuestion.CreatedAt)
}

func TestCreateQuestion_QuestionTooShort(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateQuestionMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"question": "short",
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, strings.ToLower(resp.FirstErrorMessage()), "question min length is 10 characters")
}

func TestCreateQuestion_QuestionTooLong(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateQuestionMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"question": strings.Repeat("a", 501),
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, strings.ToLower(resp.FirstErrorMessage()), "question max length is 500 characters")
}

func TestCreateQuestion_EmptyQuestion(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateQuestionMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"question": "",
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, strings.ToLower(resp.FirstErrorMessage()), "question is required")
}

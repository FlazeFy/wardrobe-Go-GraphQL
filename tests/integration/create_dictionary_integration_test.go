package integration

import (
	"strings"
	"testing"

	"wardrobe-graphql/tests/integration/client"
	"wardrobe-graphql/tests/integration/queries"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDictionary_ValidData(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateDictionaryMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"dictionaryType": "clothes_type",
				"dictionaryName": "shirt",
			},
		},
	})
	require.NoError(t, err)

	// Validate empty error message
	require.False(t, resp.HasErrors(), resp.FirstErrorMessage())

	var data queries.CreateDictionaryData
	require.NoError(t, resp.Unmarshal(&data))

	// Validate response
	assert.NotEmpty(t, data.CreateDictionary.ID)
	assert.Equal(t, "clothes_type", data.CreateDictionary.DictionaryType)
	assert.Equal(t, "shirt", data.CreateDictionary.DictionaryName)
	assert.NotEmpty(t, data.CreateDictionary.CreatedAt)
}

func TestCreateDictionary_DictionaryTooShort(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateDictionaryMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"dictionaryType": "clothes_type",
				"dictionaryName": "sh",
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, resp.FirstErrorMessage(), "dictionaryName min length is 3 characters")
}

func TestCreateDictionary_DictionaryTooLong(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateDictionaryMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"dictionaryType": strings.Repeat("a", 37),
				"dictionaryName": "shirt",
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, resp.FirstErrorMessage(), "dictionaryType max length is 36 characters")
}

func TestCreateDictionary_EmptyDictionary(t *testing.T) {
	c := client.New()

	// Exec
	resp, err := c.Do(client.Request{
		Query: queries.CreateDictionaryMutation,
		Variables: map[string]interface{}{
			// Test Data
			"input": map[string]interface{}{
				"dictionaryType": "clothes_type",
				"dictionaryName": "",
			},
		},
	})
	require.NoError(t, err)

	// Validate error message exist
	require.True(t, resp.HasErrors())

	// Validate response
	assert.Contains(t, resp.FirstErrorMessage(), "dictionaryName is required")
}

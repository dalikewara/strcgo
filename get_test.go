package strcgo_test

import (
	"github.com/dalikewara/strcgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetTagValueFromField_OK tests GetTagValueFromField function with "OK" scenario
func TestGetTagValueFromField_OK(t *testing.T) {
	type s struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	tagValue := strcgo.GetTagValueFromField(s{}, "Name", "json")

	assert.NotEmpty(t, tagValue)
	assert.Equal(t, "name", tagValue)
}

// TestGetTagValueFromField_OK tests GetTagValueFromField function with "InvalidName" scenario
func TestGetTagValueFromField_InvalidName(t *testing.T) {
	t.Run("InvalidName_Field", func(t *testing.T) {
		type s struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		}

		tagValue := strcgo.GetTagValueFromField(s{}, "Name2", "json")

		assert.Empty(t, tagValue)
	})

	t.Run("InvalidName_Tag", func(t *testing.T) {
		type s struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		}

		tagValue := strcgo.GetTagValueFromField(s{}, "Name", "json2")

		assert.Empty(t, tagValue)
	})
}

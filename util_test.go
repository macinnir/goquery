package goquery_test

import (
	"testing"

	query "github.com/macinnir/goquery"
	"github.com/stretchr/testify/assert"
)

func TestEscapeString(t *testing.T) {

	result := query.EscapeString("I'm a string")
	assert.Equal(t, `I\'m a string`, result)

	result = query.EscapeString(`I"m a string`)
	assert.Equal(t, `I\"m a string`, result)

}

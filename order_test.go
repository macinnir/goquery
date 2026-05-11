package goquery_test

import (
	"testing"

	query "github.com/macinnir/goquery"
	"github.com/stretchr/testify/assert"
)

func TestOrderDirFromString(t *testing.T) {
	assert.Equal(t, query.OrderDirASC, query.OrderDirFromString("asc"))
	assert.Equal(t, query.OrderDirDESC, query.OrderDirFromString("desc"))
	assert.Equal(t, query.OrderDirASC, query.OrderDirFromString("foo"))
}

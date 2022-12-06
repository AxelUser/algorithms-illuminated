package mwis_test

import (
	"testing"

	"github.com/AxelUser/algorithms-illuminated/dynamic-programming/mwis"
	"github.com/stretchr/testify/assert"
)

func TestMaxSetWeight(t *testing.T) {
	input := []int{3, 2, 1, 6, 4, 5}
	expected := 14

	actual, _ := mwis.MaxSetWeight(input)
	assert.Equal(t, expected, actual)
}

func TestSetWithMaxWeight(t *testing.T) {
	input := []int{3, 2, 1, 6, 4, 5}
	expected := []int{0, 3, 5}

	actual, _ := mwis.SetWithMaxWeight(input)
	assert.ElementsMatch(t, expected, actual)
}

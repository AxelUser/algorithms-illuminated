package knapsack_test

import (
	"testing"

	"github.com/AxelUser/algorithms-illuminated/dynamic-programming/knapsack"
	"github.com/stretchr/testify/assert"
)

func TestMaxSum(t *testing.T) {
	assert := assert.New(t)

	values := []int{3, 2, 4, 4}
	sizes := []int{4, 3, 2, 3}
	capacity := 6
	expected := 8

	actual, _ := knapsack.MaxSum(values, sizes, capacity)
	assert.Equal(expected, actual)
}

func TestSetWithMaxSum(t *testing.T) {
	assert := assert.New(t)

	values := []int{3, 2, 4, 4}
	sizes := []int{4, 3, 2, 3}
	capacity := 6
	expected := []int{2, 3}

	actual, _ := knapsack.SetWithMaxSum(values, sizes, capacity)
	assert.ElementsMatch(expected, actual)
}

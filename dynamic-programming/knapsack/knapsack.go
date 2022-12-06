package knapsack

import (
	"errors"

	"github.com/AxelUser/algorithms-illuminated/common"
)

func MaxSum(values []int, sizes []int, capacity int) (int, error) {
	dp, err := compute(values, sizes, capacity)
	if err != nil {
		return 0, err
	}

	return dp[len(dp)-1][capacity], nil
}

func SetWithMaxSum(values []int, sizes []int, capacity int) ([]int, error) {
	dp, err := compute(values, sizes, capacity)
	if err != nil {
		return nil, err
	}

	var set []int
	remCap := capacity
	for i := len(values); i >= 1; i-- {
		si := sizes[i-1]
		if sizes[i-1] <= remCap && dp[i-1][remCap-si]+values[i-1] >= dp[i-1][remCap-si] {
			set = append(set, i-1)
			remCap -= si
		}
	}

	return set, nil
}

func compute(values []int, sizes []int, capacity int) ([][]int, error) {
	if len(values) != len(sizes) {
		return nil, errors.New("values and sizes have different length")
	}

	if len(values) == 0 {
		return nil, errors.New("should contain at least one element")
	}

	n := len(values)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for c := 0; c <= capacity; c++ {
			if sizes[i-1] > c {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = common.Max(dp[i-1][c], dp[i-1][c-sizes[i-1]]+values[i-1])
			}
		}
	}

	return dp, nil
}

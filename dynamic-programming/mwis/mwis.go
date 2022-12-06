package mwis

import (
	"errors"

	"github.com/AxelUser/algorithms-illuminated/common"
)

func MaxSetWeight(input []int) (int, error) {
	dp, err := compute(input)
	if err != nil {
		return 0, err
	}

	return dp[len(dp)-1], nil
}

func SetWithMaxWeight(input []int) ([]int, error) {
	dp, err := compute(input)
	if err != nil {
		return nil, err
	}

	i := len(input)
	set := make([]int, 0)

	for i >= 2 {
		if dp[i-1] >= dp[i-2]+input[i-1] {
			i--
		} else {
			set = append(set, i-1)
			i -= 2
		}
	}

	if i == 1 {
		set = append(set, i-1)
	}

	return set, nil
}

func compute(input []int) ([]int, error) {
	if len(input) == 0 {
		return nil, errors.New("graph should contain at least one element")
	}

	dp := make([]int, len(input)+1)
	dp[1] = input[0]
	for i := 2; i <= len(input); i++ {
		dp[i] = common.Max(dp[i-1], dp[i-2]+input[i-1])
	}

	return dp, nil
}

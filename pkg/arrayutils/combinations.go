package arrayutils

import "errors"

func GenerateCombinations[T any](collection []T, n int) ([][]T, error) {
	if n > len(collection) || n <= 0 {
		return nil, errors.New("bad size of combination")
	}

	result := [][]T{}
	combination := []T{}

	var helper func(start int)
	helper = func(start int) {
		if len(combination) == n {
			combCopy := make([]T, n)
			copy(combCopy, combination)
			result = append(result, combCopy)
			return
		}

		for i := start; i < len(collection); i++ {
			combination = append(combination, collection[i])
			helper(i + 1)
			combination = combination[:len(combination)-1]
		}
	}

	helper(0)
	return result, nil
}

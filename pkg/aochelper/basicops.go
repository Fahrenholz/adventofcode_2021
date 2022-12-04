package aochelper

func SumScore[T any](vals []T, score func(elem T) int) int {
	var sum int

	for i, _ := range vals {
		sum += score(vals[i])
	}

	return sum
}

func Sum(vals []int) int {
	return SumScore(vals, func(elem int) int { return elem })
}

func Filter[T any](sl []T, cond func(i int) bool) []T {
	var res []T

	for i, _ := range sl {
		if cond(i) {
			res = append(res, sl[i])
		}
	}

	return res
}

func FilterMap[K comparable, T any](vals map[K]T, cond func(key K) bool) map[K]T {
	res := make(map[K]T)

	for k, _ := range vals {
		if cond(k) {
			res[k] = vals[k]
		}
	}

	return res
}

func Map[T, U any](vals []T, f func(T) U) []U {
	var res []U

	for _, v := range vals {
		res = append(res, f(v))
	}

	return res
}

func MapMap[K comparable, T, U any](vals map[K]T, f func(T) U) map[K]U {
	res := make(map[K]U)

	for i, v := range vals {
		res[i] = f(v)
	}

	return res
}

func Reduce[T, U any](vals []T, init U, f func(U, T) U) U {
	res := init
	for _, v := range vals {
		res = f(res, v)
	}

	return res
}

func ReduceMap[K comparable, T, U any](vals map[K]T, init U, f func(U, T) U) U {
	res := init
	for _, v := range vals {
		res = f(res, v)
	}

	return res
}

func SliceToMap[T comparable](vals []T) map[T]bool {
	res := make(map[T]bool)

	for _, v := range vals {
		res[v] = true
	}

	return res
}

func MapToSlice[K comparable, T any](vals map[K]T) []T {
	var res []T

	for _, v := range vals {
		res = append(res, v)
	}

	return res
}

func MaxScore[T any](vals []T, score func(elem T) int) T {
	var max int
	var maxElem T

	for i, _ := range vals {
		curr := score(vals[i])

		if curr > max {
			max = curr
			maxElem = vals[i]
		}
	}

	return maxElem
}

func Max(vals []int) int {
	return MaxScore(vals, func(elem int) int { return elem })
}

func MinScore[T any](vals []T, score func(elem T) int) T {
	var min int
	var minElem T

	for i, _ := range vals {
		curr := score(vals[i])

		if curr < min {
			min = curr
			minElem = vals[i]
		}
	}

	return minElem
}

func Min(vals []int) int {
	return MinScore(vals, func(elem int) int { return elem })
}

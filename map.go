package evaluation2

func Map(f func(int) bool, arr []int) []bool {
	var result []bool
	for i := 0; i < len(arr); i++ {
		if f(arr[i]) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

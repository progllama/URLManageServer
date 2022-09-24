package utility

func Merge(m ...map[string]interface{}) map[string]interface{} {
	ans := EmptyMap()

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func EmptyMap() map[string]interface{} {
	return make(map[string]interface{}, 0)
}

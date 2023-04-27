package process

func IF(condition bool, result1, result2 any) any {
	if condition {
		return result1
	}
	return result2
}

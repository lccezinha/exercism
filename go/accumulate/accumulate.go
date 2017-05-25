package accumulate

const testVersion = 1

func Accumulate(values []string, converter func(string) string) []string {
	result := make([]string, len(values))

	for index, value := range values {
		result[index] = converter(value)
	}

	return result
}

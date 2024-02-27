package iteration

func Repeat(input string, count int) (output string) {
	for i := 0; i < count; i++ {
		output += input
	}
	return
}

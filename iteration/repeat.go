package iteration

const repeatCount = 4

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

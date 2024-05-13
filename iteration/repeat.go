package iteration

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < 4; i++ {
		repeated += character
	}
	return repeated
}

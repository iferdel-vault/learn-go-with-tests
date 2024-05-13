package iteration

// no longer needed since strings.Repeat function do the job.
func Repeat(character string, repeatCount int) string {
	repeated := Repeat("a", 4)
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

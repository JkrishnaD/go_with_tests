package itretion

var repeatCount = 4

func Repeat(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

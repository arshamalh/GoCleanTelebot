package validation

func PanicIfEmpty(s string, msg string) {
	if s == "" {
		panic(msg)
	}
}

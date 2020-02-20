package utils

// DieIf throw error if found
func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}

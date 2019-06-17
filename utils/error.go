package utils

// CheckError : if error, panic it
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

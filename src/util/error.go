package util

func Panic(e error) {
	if e != nil {
		panic(e)
	}
}

func PanicMsg(e error, msg string) {
	if e != nil {
		panic(msg)
	}
}
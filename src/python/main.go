package main

//go build -buildmode=c-shared -o hello.so main.go

import (
	"C"
)

//export Hello
func Hello() *C.char {
	return C.CString("python import go mode test...")
}

//export Cstr
func Cstr(s *C.char) *C.char {
	gostr := C.GoString(s)
	return C.CString(gostr)
}

func main()  {
	
}
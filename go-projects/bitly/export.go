package main

import "C"

//export getLink
func getLink(url *C.char) {
	msg := C.GoString(url)

	exportLink(msg)
}

func main() {}

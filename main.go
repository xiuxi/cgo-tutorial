package main

// The comments below are important.
// They're actually C code that gets used by
// our program when it's being compiled with cgo.
// Additionally, we must have a separate line for
// `import "C"` otherwise our C directives won't be
// considered.

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// C.CString returns a char pointer
	// in C, which if you remember is the
	// way to create a string in C. This
	// constructor will call malloc to create
	// the space needed for our string so we must
	// free the memory when we're finished with the
	// string.
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	// C integer primative... guessing
	// this uses the word size of the machine
	// compiling the code. Does cgo allow us to
	// specify integers as unsigned/16bit/32bit?
	year := C.int(2020)

	// We can access any C struct by using the
	// C.struct_ syntax then we can initialize the
	// struct just as we would with any Go struct
	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	// Here we allocate the buffer that will
	// store the greeting. We're not checking for
	// buffer overflow here... Might fix that later
	// as an exercise.
	// ***CGo note: if malloc fails, the program will
	// crash. There is no way to catch this error, this
	// is by design.***
	// Calling C.malloc returns a Go unsafe.Pointer
	// which can be cast to any type.
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	// In order to call the greeting function in
	// our C program, we must cast our unsafe.Pointer
	// to a *C.char.
	size := C.greet(&g, (*C.char)(ptr))

	// C.GoBytes takes in a C buffer and the size
	// of the data in the buffer and returns a []byte.
	// We can then cast this to a string and print it
	// using normal Go syntax!
	// ***CGo note: when we call C.GoBytes, new memory
	// is allocated for our []byte so we can free our
	// C pointer without fear of losing our data.
	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}

// Package cracklib provides a Golang binding for cracklib
// https://github.com/cracklib/cracklib
package cracklib

// #cgo LDFLAGS: -lcrack
// #include <stdlib.h>
// #include <crack.h>
import "C"
import "unsafe"

// FascistCheck checks a potential password for guessability
// It returns an error message and a boolean value
// The error message will be "" if ok is true
func FascistCheck(pw, path string) (message string, ok bool) {
	pathptr := C.CString(path)
	pwptr := C.CString(pw)
	defer C.free(unsafe.Pointer(pathptr))
	defer C.free(unsafe.Pointer(pwptr))

	v := C.FascistCheck(pwptr, pathptr)
	message = C.GoString(v)
	if message != "" {
		return message, false
	}
	return "", true
}

// FascistCheckDefault checks a potential password for guessability
// It returns an error message and a boolean value
// The error message will be "" if ok is true
func FascistCheckDefault(pw string) (message string, ok bool) {
	pathptr := C.GetDefaultCracklibDict()

	pwptr := C.CString(pw)
	defer C.free(unsafe.Pointer(pwptr))

	v := C.FascistCheck(pwptr, pathptr)
	message = C.GoString(v)
	if message != "" {
		return message, false
	}
	return "", true
}

// extern const char *FascistCheck(const char *pw, const char *dictpath);
//
// /* This function returns the compiled in value for DEFAULT_CRACKLIB_DICT.
//  */
// extern const char *GetDefaultCracklibDict(void);

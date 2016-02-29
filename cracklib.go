// Package cracklib provides a Golang binding for cracklib
// https://github.com/cracklib/cracklib
package cracklib

// #cgo LDFLAGS: -lcrack
// #include <crack.h>
import "C"

// FascistCheck checks a potential password for guessability
// It returns an error message and a boolean value
// The error message will be "" if ok is true
func FascistCheck(pw string) (message string, ok bool) {
	path := C.GetDefaultCracklibDict()
	v := C.FascistCheck(C.CString(pw), path)
	message = C.GoString(v)
	if message != "" {
		return message, false
	}
	return "", true
}

// FascistCheckUser executes tests against an arbitrary user
// It returns an error message and a boolean value
// The error message will be "" if ok is true
func FascistCheckUser(pw string, user string) (message string, ok bool) {
	path := C.GetDefaultCracklibDict()
	v := C.FascistCheckUser(C.CString(pw), path, C.CString(user), nil)
	message = C.GoString(v)
	if message != "" {
		return message, false
	}
	return "", true
}

// extern const char *FascistCheck(const char *pw, const char *dictpath);
// extern const char *FascistCheckUser(const char *pw, const char *dictpath,
// 				    const char *user, const char *gecos);
//
// /* This function returns the compiled in value for DEFAULT_CRACKLIB_DICT.
//  */
// extern const char *GetDefaultCracklibDict(void);

// The wslcheck package provides a simple function that checks whether the
// current Linux environment is in fact the Windows Subsystem for Linux (WSL).
package wslcheck

// Check returns true if the Linux version string matches that of a
// WSL distribution. Otherwise (including when the host OS is not Linux
// at all), it returns false.
//
// If an error occurs when attempting to get the version string, Check will
// return false and the error value will be populated (not-nil).
func Check() (bool, error) {
	return check()
}

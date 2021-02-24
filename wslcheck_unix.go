package wslcheck

import (
	"io/ioutil"
	"runtime"
)

// Check if the current kernel's version string is one that matches WSL
func check() (bool, error) {
	runt := runtime.GOOS

	if runt == "linux" {
		verBytes, err := ioutil.ReadFile(ProcReleasePath)
		if err != nil {
			return false, err
		}
		return CheckVer(verBytes), nil
	}

	return false, nil
}

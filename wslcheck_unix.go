package wslcheck

import (
	"io/ioutil"
	"runtime"
	"strings"
)

const (
	linuxReleaseProc = "/proc/sys/kernel/osrelease"
)

func check() (bool, error) {
	runt := runtime.GOOS

	if runt == "linux" {
		verBytes, err := ioutil.ReadFile(linuxReleaseProc)
		if err != nil {
			return false, err
		}
		if strings.HasSuffix(string(verBytes), "microsoft-standard\n") {
			return true, nil
		}
	}

	return false, nil
}

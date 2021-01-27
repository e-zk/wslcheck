package wslcheck

import (
	"io/ioutil"
	"regexp"
	"runtime"
)

const (
	linuxReleaseProc = "/proc/sys/kernel/osrelease"
)

func check() (bool, error) {
	runt := runtime.GOOS
	r := regexp.MustCompile(`.*microsoft-standard.*`)

	if runt == "linux" {
		verBytes, err := ioutil.ReadFile(linuxReleaseProc)
		if err != nil {
			return false, err
		}
		//if strings.HasSuffix(string(verBytes), "microsoft-standard\n") {
		//	return true, nil
		//}
		if r.Match(verBytes) {
			return true, nil
		}
	}

	return false, nil
}

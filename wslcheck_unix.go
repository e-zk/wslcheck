package wslcheck

import (
	"io/ioutil"
	"regexp"
	"runtime"
)

func check() (bool, error) {
	runt := runtime.GOOS
	r := regexp.MustCompile(VersionRegexp)

	if runt == "linux" {
		verBytes, err := ioutil.ReadFile(ProcReleasePath)
		if err != nil {
			return false, err
		}
		if r.Match(verBytes) {
			return true, nil
		}
	}

	return false, nil
}

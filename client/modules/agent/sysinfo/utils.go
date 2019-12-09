package sysinfo

import (
	"io/ioutil"
	"strings"
)

func FileToString(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

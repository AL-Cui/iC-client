package sysinfo

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

//type OS struct {
//	Name     string
//	Version  string
//	CodeName string
//}

var reOSSeparator = regexp.MustCompile(":[\t ]+")

func (c *Collector) getOSInfo() {
	output, err := exec.Command("lsb_release", "-a").Output()
	if err != nil {
		return
	}
	r := bytes.NewReader(output)
	bf := bufio.NewScanner(r)
	for bf.Scan() {
		if col := reOSSeparator.Split(bf.Text(), 2); col != nil {
			switch col[0] {
			case "Distributor ID":
				c.Info.OS.Name = strings.TrimSpace(col[1])
			case "Release":
				c.Info.OS.Version = strings.TrimSpace(col[1])
			case "Codename":
				c.Info.OS.CodeName = strings.TrimSpace(col[1])
			}
		}
	}
}

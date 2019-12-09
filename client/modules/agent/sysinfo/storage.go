package sysinfo

import (
	"bufio"
	"bytes"
	pb "iCenter-client/models/proto"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//type Storage struct {
//	FileSystem string
//	Mount      string
//	Total      int64
//	Used       int64
//	Free       int64
//}

var (
	reStorageSeparator = regexp.MustCompile("[\t ]+")
	reRemoteStorage    = regexp.MustCompile(`\S+:\S+`)
)

func (c *Collector) getStorageInfo() {
	c.Info.Storages = map[string][]*pb.Storage{}
	output, err := exec.Command("df", "-P", "-x", "tmpfs", "-x", "devtmpfs").Output()
	if err != nil {
		return
	}
	r := bytes.NewReader(output)
	b := bufio.NewScanner(r)
	localStorages := []*pb.Storage{}
	remoteStorages := []*pb.Storage{}
	// count to remove headers at first column. a
	for b.Scan() {

		if col := reStorageSeparator.Split(b.Text(), 6); col != nil {
			if col[0] == "Filesystem" || strings.Contains(col[5], "docker") {
				continue
			}
			total, _ := strconv.ParseInt(col[1], 10, 64)
			used, _ := strconv.ParseInt(col[2], 10, 64)
			free, _ := strconv.ParseInt(col[3], 10, 64)
			storage := &pb.Storage{
				FileSystem: col[0],
				Mount:      col[5],
				Total:      total,
				Used:       used,
				Free:       free,
			}
			if reRemoteStorage.MatchString(storage.FileSystem) {
				remoteStorages = append(remoteStorages, storage)
			} else {
				localStorages = append(localStorages, storage)
			}
		}
	}
	c.Info.Storages["Remote"] = remoteStorages
	c.Info.Storages["Local"] = localStorages
}

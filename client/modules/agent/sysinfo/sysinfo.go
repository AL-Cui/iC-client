package sysinfo

import (
	"fmt"
	"iCenter-client/conf"
	pb "iCenter-client/models/proto"
	"iCenter-client/utils/logs"
	"net"
	"time"
)

type SysInfo struct {
	OS       pb.OS
	CPU      pb.CPU
	GPU      pb.GPU
	Kernel   pb.Kernel
	Memory   pb.Memory
	Networks []*pb.Network
	Storages map[string][]*pb.Storage
}

var collector *Collector

type Collector struct {
	Info      SysInfo
	Interval  time.Duration
	mgmtIPNet net.IPNet
	hostname  string
}

func GetSysInfoCollector(hostname string) *Collector {
	if collector != nil {
		return collector
	}
	mgmtIPs, err := net.LookupHost(conf.MgmtHost())
	if err != nil {
		panic(fmt.Sprintf("can't resolve mgmt host address: %v", err))
	}
	if len(mgmtIPs) == 0 {
		panic("can't find mgmt host IP address")
	}
	ipNet := net.IPNet{
		IP:   net.ParseIP(mgmtIPs[0]),
		Mask: net.IPMask(net.ParseIP(conf.MgmtSubnetMask())),
	}
	collector = &Collector{
		mgmtIPNet: ipNet,
		hostname:  hostname,
		Interval:  time.Second * time.Duration(conf.MonitorInterval()),
	}
	err = collector.run()
	//TODO
	return collector
}

func (c *Collector) GetSysInfo() (*SysInfo, error) {
	c.getOSInfo()
	c.getCPUInfo()
	c.getGPUInfo()
	c.getKernelInfo()
	c.getMemoryInfo()
	err := c.getNetworkInfo()
	if err != nil {
		return nil, err
	}
	c.getStorageInfo()
	return &c.Info, nil
}

func (c *Collector) run() error {
	_, err := c.GetSysInfo()
	if err != nil {
		logs.Error("collector can't get system info %v", err)
		return err
	}
	return nil
}

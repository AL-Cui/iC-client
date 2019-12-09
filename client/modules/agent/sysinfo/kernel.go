package sysinfo

//type Kernel struct {
//	OSRelease string
//	OSType    string
//	Version   string
//}

func (c *Collector) getKernelInfo() {
	c.Info.Kernel.OSRelease = FileToString("/proc/sys/kernel/osrelease")
	c.Info.Kernel.OSType = FileToString("/proc/sys/kernel/ostype")
	c.Info.Kernel.Version = FileToString("/proc/sys/kernel/version")
}

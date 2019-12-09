package routers

import (
	"strings"

	pb "iCenter-client/models/proto"
	"iCenter-client/modules/agent/sysinfo"
	"iCenter-client/utils/logs"

	"golang.org/x/net/context"
)

type Server struct{}

// GetHardwareInfo get computer cpu info from local
func (s *Server) GetHardwareInfo(ctx context.Context, in *pb.HardwareRequest) (*pb.HardwareResponse, error) {
	logs.Debug("received: %s", in.Name)
	collector := sysinfo.GetSysInfoCollector("")
	storageMap := make(map[string]*pb.StorageArray)
	remoteStorageArray := pb.StorageArray{Storages: collector.Info.Storages["Remote"]}
	localStorageArray := pb.StorageArray{Storages: collector.Info.Storages["Local"]}
	storageMap[""] = &remoteStorageArray
	storageMap[""] = &localStorageArray
	//any, err := ptypes.MarshalAny(collector.Info.Storages)
	//if err != nil {
	//
	//}
	//cpuInfo := pb.CPU{Core: *proto.Int(collector.Info.CPU.Core)}
	cpuInfo := &pb.HardwareResponse{
		CPUInfo: &collector.Info.CPU,
		GPUInfos: &collector.Info.GPU,
		OSInfo: &collector.Info.OS,
		KernelInfo: &collector.Info.Kernel,
		MemoryInfo: &collector.Info.Memory,
		NetworkInfos: collector.Info.Networks,
		StorageInfo: storageMap,

	}
	return cpuInfo, nil
}

func (s *Server) Upper(ctx context.Context, in *pb.UpperRequest) (*pb.UpperReply, error) {
	logs.Debug("Received: %s", in.Name)
	return &pb.UpperReply{Message: strings.ToUpper(in.Name)}, nil
}

func (s *Server) Test() (*sysinfo.Collector, error) {
	sys := sysinfo.GetSysInfoCollector("")
	return sys, nil
}

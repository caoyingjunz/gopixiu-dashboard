package cmd

import (
	"context"
	"os/exec"

	"k8s.io/klog"

	"grpc-pixiu/options"
	pixiupb "grpc-pixiu/pixiu"
)

var CreateClusterService = &createClusterService{}

type createClusterService struct {
}

func (p *createClusterService) Check(ctx context.Context, clusterInfo *pixiupb.ClusterRequest) (*pixiupb.ClusterResponse, error) {
	var err error
	startTime := options.GetStartTime()

	// 检查系统内是否安装kubez-ansible
	checkCmd := exec.Command(options.CheckKubezCommand)
	err = checkCmd.Run()
	if err != nil {
		installCmd := exec.Command("/bin/bash", "-c", options.InstallKubezCommand)
		err := installCmd.Run()
		if err == nil {
			multinodeCheckCmd := exec.Command(options.MultinodeCheckCmd)
			if err = multinodeCheckCmd.Run(); err != nil {
				klog.Errorf("Multinode configuration fail:", err)
			}
		}
		klog.Errorf("precondition fail:", err)
	}
	endTime := options.GetEndTime()
	return &pixiupb.ClusterResponse{
		ResponseInfo: "precondition successful:",
		StartTime:    startTime,
		EndTime:      endTime,
	}, err
}

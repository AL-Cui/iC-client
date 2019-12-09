package pod

import (
	"iCenter-client/models/mdef"
)
// PodManager represents the pod manager.
type PodManager struct {

}

// NewManager initials PodManager.
func NewManager() (*PodManager, error) {
	return &PodManager{}, nil
}

// ListPods list podInfo from namespace.
func (p *PodManager)ListPods() (*[]mdef.PodInfo, error){
	return nil, nil
}
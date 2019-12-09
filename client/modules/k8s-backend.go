package modules

import (
	"iCenter-client/def"
	"iCenter-client/modules/pod"
	"iCenter-client/utils/errors"
)

// KubernetesServer the KubernetesServer backend
var KubernetesServer *Backend

// Backend the backend struct
type Backend struct {
	PodManager *pod.PodManager
	inited    bool
}

func NewBackend() (*Backend, error) {
	if KubernetesServer != nil && KubernetesServer.inited {
		return KubernetesServer, nil
	}
	podManager,err := pod.NewManager()
	if err != nil {
		return nil, errors.Wrap(err, def.ErrPodModule,"init pod module failed")
	}
	
	backend := &Backend{
		PodManager: podManager,
		inited: true,
	}
	return backend,nil
}
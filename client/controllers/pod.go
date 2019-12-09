package controllers

import (
	"iCenter-client/def"
	"iCenter-client/modules"
)

type PodController struct {
	BaseController
}

func (p *PodController) nestPrepare() {

}

func (p *PodController) ListPodsFromNamespace() {
	pods, err := modules.KubernetesServer.PodManager.ListPods()
	if err != nil {
		p.APIAbort(def.ListPodsFromNamespace, err)
	}
	p.Data["json"] = pods
	p.ServeJSON()
	
}
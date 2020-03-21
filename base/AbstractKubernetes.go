package base

import (
	"encoding/json"
	"../bean"
	"../logger"
)

var (
	TYPE_SERVICE    = "Service"
	TYPE_DEPLOYMENT = "Deployment"
	TYPE_POD        = "Pod"
	TYPE_JOB        = "Job"
	TYPE_CONFIG_MAP = "ConfigMap"
)

type AbstractKubernetes struct {
}

func NewAbstractKubernetes() *AbstractKubernetes {
	h := &AbstractKubernetes{}
	return h
}

func (*AbstractKubernetes) GetResourceTypes() (string) {
	service := bean.BaseBean{ApiVersion: "v1", Kind: TYPE_SERVICE}
	deployment := bean.BaseBean{ApiVersion: "v1", Kind: TYPE_DEPLOYMENT}
	Pod := bean.BaseBean{ApiVersion: "v1", Kind: TYPE_POD}
	Job := bean.BaseBean{ApiVersion: "v1", Kind: TYPE_JOB}
	configMap := bean.BaseBean{ApiVersion: "v1", Kind: TYPE_CONFIG_MAP}

	resourceTypes := bean.ResourceTypes{Service: service, Deployment: deployment, Pod: Pod, Job: Job, ConfigMap: configMap}

	jsonString, _ := json.Marshal(resourceTypes)
	logger.Info(string(jsonString))
	return string(jsonString)
}

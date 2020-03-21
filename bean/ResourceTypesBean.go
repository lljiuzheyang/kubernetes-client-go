package bean

type ResourceTypes struct {
	Service    BaseBean `json:"service"`
	Deployment BaseBean `json:"deployment"`
	Pod        BaseBean `json:"pod"`
	Job        BaseBean `json:"job"`
	ConfigMap  BaseBean `json:"config_map"`
}


package bean

type BaseBean struct {
	ApiVersion string `json:"api_version"`
	Kind       string `json:"kind"`
}

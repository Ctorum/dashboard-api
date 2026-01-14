package infra

type BaseResponseEnveloper struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

package m2mmodels

type TokenModel struct {
	ClusterId string `json:"clusterId" validate:"required,min=1,ne=' '" `
	Token     string `json:"token" validate:""`
}

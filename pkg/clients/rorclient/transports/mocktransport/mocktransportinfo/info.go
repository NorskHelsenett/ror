package mocktransportinfo

import "context"

type V1Client struct {
}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) GetVersion(ctx context.Context) (string, error) {
	return "1.1.1", nil
}

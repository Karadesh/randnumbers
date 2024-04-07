package randomNumbers

import (
	"context"
	api "grpcserverclient/pkg/api/randomNumbers"
)

type RandomNumbersServer struct{}

func (c *RandomNumbersServer) Generate(ctx context.Context, req *api.GenRequest) (*GenResponse, error) {
	return

}

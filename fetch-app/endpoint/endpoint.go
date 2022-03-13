package endpoint

import (
	"context"

	"github.com/khakim88/sample-auth/fetch-app/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetResource endpoint.Endpoint
}

func MakeEndpoints(svc service.ResourceService) Endpoints {
	return Endpoints{
		GetResource: MakeGetResource(svc),
	}
}

func MakeGetResource(svc service.ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		return svc.ValidatePromotionService(ctx)
	}
}

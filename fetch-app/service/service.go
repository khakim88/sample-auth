package service

import (
	"context"

	"github.com/khakim88/sample-auth/fetch-app/model"
)

type ResourceService interface {
	GetResourceService(ctx context.Context) (resp *model.GetResourceResponse, err error)
}

func GetResourceService(ctx context.Context) (resp *model.GetResourceResponse, err error) {

	return &model.GetResourceResponse{}, nil
}

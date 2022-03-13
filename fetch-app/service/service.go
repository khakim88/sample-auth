package service

import (
	"context"

	"github.com/khakim88/sample-auth/fetch-app/model"
	"github.com/khakim88/sample-auth/fetch-app/repository"
)

type resourceService struct {
	repo repository.DBReaderWriter
}

func NewResourceService(repo repository.DBReaderWriter) resourceService {
	return resourceService{
		repo: repo,
	}
}

type ResourceService interface {
	GetResourceService(ctx context.Context) (resp *model.GetResourceResponse, err error)
}

package app_version

import (
	"mono-base/internal/entities"
	"mono-base/internal/repositories"

	"github.com/gin-gonic/gin"
)

type GetListAppVersionInput struct {
}

type GetListAppVersionOutput struct {
	AppVersion []entities.AppVersionData
}

type GetListAppVersionUseCase interface {
	Execute(ctx *gin.Context, input GetListAppVersionInput) (*GetListAppVersionOutput, error)
}

type getListAppVersionUseCase struct {
	appVersionRepo repositories.AppVersionRepository
}

func NewGetListAppVersionUseCase(appVersionRepo repositories.AppVersionRepository) GetListAppVersionUseCase {
	return &getListAppVersionUseCase{
		appVersionRepo: appVersionRepo,
	}
}

func (u *getListAppVersionUseCase) Execute(ctx *gin.Context, input GetListAppVersionInput) (*GetListAppVersionOutput, error) {
	output, err := u.appVersionRepo.GetListAppVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &GetListAppVersionOutput{
		AppVersion: output,
	}, nil
}

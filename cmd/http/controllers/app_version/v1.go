package app_version

import (
	"github.com/gin-gonic/gin"
	"mono-base/cmd/http/rest"
	"mono-base/internal/usecases/app_version"
	"mono-base/pkg/http"
	"mono-base/pkg/utils"
)

var _ Controller = (*ControllerV1)(nil)

type ControllerV1 struct {
	getListAppVersionUseCase app_version.GetListAppVersionUseCase
}

func NewStoryControllerV1(
	getListAppVersionUseCase app_version.GetListAppVersionUseCase) *ControllerV1 {
	return &ControllerV1{getListAppVersionUseCase: getListAppVersionUseCase}
}

func (s *ControllerV1) GetListVersion(ctx *gin.Context) {

	// Execute use case
	output, err := s.getListAppVersionUseCase.Execute(ctx, app_version.GetListAppVersionInput{})

	if err != nil {
		rest.HandleError(ctx, err)
		return
	}

	var result []ListAppVersionResponse
	err = utils.MappingInterface(output.AppVersion, &result)
	if err != nil {
		return
	}

	response := result

	http.SuccessResponse(ctx, response)
}

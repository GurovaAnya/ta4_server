package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/project"
	"ta4/mod/repository"
)

type GetProjectHandler struct {
	repo repository.PostgresRepository
}

func GetProjectHandlerImpl() project.GetProjectHandler {
	return GetProjectHandler{repo: repository.PostgresRepository{}}
}

func (h GetProjectHandler) Handle(params project.GetProjectParams) middleware.Responder {
	var response = h.repo.GetProject(params.HTTPRequest.Header.Get("apiKey"))

	var responseVal = &models.Project{
		ID:         strfmt.UUID(response.Id.String()),
		Name:       &response.Name,
		WebhookURL: &response.WebhookUrl,
	}
	return project.NewGetProjectOK().WithPayload(responseVal)
}

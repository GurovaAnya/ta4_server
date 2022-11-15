package project

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"ta4/mod/internal/app"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/project"
	"ta4/mod/repository"
)

type GetProject struct {
	repo repository.ProjectRepository
}

func GetProjectHandlerImpl(a *app.Application) GetProject {
	return GetProject{
		repo: repository.NewProjectRepository(a.Store),
	}
}

func (h *GetProject) Handle(params project.GetProjectParams) middleware.Responder {
	var response = h.repo.GetProject(params.HTTPRequest.Header.Get("apiKey"))

	var responseVal = &models.Project{
		ID:         strfmt.UUID(response.Id.String()),
		Name:       &response.Name,
		WebhookURL: &response.WebhookUrl,
	}
	return project.NewGetProjectOK().WithPayload(responseVal)
}

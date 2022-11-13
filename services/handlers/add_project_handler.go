package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/project"
	"ta4/mod/repository"
)

type AddProjectHandler struct {
	repo repository.PostgresRepository
}

func AddProjectHandlerImpl() project.AddProjectHandler {
	return AddProjectHandler{repo: repository.PostgresRepository{}}
}

func (h AddProjectHandler) Handle(params project.AddProjectParams) middleware.Responder {
	var projectRequest = params.Body
	var projectForSave = &repository.ProjectModel{
		Id:         uuid.New(),
		WebhookUrl: *projectRequest.WebhookURL,
		Name:       *projectRequest.Name,
		ApiKey:     "someApiKey",
	}
	var result = h.repo.CreateProject(*projectForSave)
	var responseVal = &models.ProjectResponse{
		ID:     strfmt.UUID(result.Id.String()),
		APIKey: result.ApiKey,
	}
	return project.NewAddProjectCreated().WithPayload(responseVal)
}

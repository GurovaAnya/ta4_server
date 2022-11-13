package event

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/event"
	"ta4/mod/model"
	"ta4/mod/repository"
)

type AddEventHandler struct {
	eventRepository   repository.EventRepository
	projectRepository repository.ProjectRepository
}

func AddEventHandlerImpl() event.PostEventHandler {
	return AddEventHandler{
		eventRepository:   repository.EventRepository{},
		projectRepository: repository.ProjectRepository{},
	}
}

func (h AddEventHandler) Handle(params event.PostEventParams) middleware.Responder {
	var project = h.projectRepository.GetProject(params.HTTPRequest.Header.Get("apiKey"))

	var eventRequest = params.Body
	var eventForSave = &model.EventModel{
		Id:          uuid.New(),
		Sku:         eventRequest.Sku,
		Description: eventRequest.Description,
		ProjectId:   project.Id,
	}

	var result = h.eventRepository.CreateEvent(*eventForSave)
	var responseVal = &models.Event{
		Description: result.Description,
		ID:          strfmt.UUID(result.Id.String()),
		Sku:         result.Sku,
	}
	return event.NewPostEventCreated().WithPayload(responseVal)
}

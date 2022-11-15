package trigger

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"ta4/mod/internal/app"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/trigger"
	"ta4/mod/model"
	"ta4/mod/repository"
)

type AddTriggerHandler struct {
	triggerRepository repository.TriggerRepository
	projectRepository repository.ProjectRepository
}

func AddTriggerHandlerImpl(a *app.Application) trigger.PostTriggerHandler {
	return AddTriggerHandler{
		triggerRepository: repository.NewTriggerRepository(a.Store),
		projectRepository: repository.NewProjectRepository(a.Store),
	}
}

func (h AddTriggerHandler) Handle(params trigger.PostTriggerParams) middleware.Responder {
	var project = h.projectRepository.GetProject(params.HTTPRequest.Header.Get("apiKey"))

	var triggerRequest = params.Body

	var id = uuid.New()
	var triggerForSave = &model.TriggerModel{
		Id:            id,
		Sku:           triggerRequest.Sku,
		Description:   triggerRequest.Description,
		AchievementId: ConvertStrFrmToUuid(triggerRequest.AchievementID),
		ProjectId:     project.Id,
	}
	var result = h.triggerRepository.CreateTrigger(*triggerForSave)

	var responseVal = &models.Trigger{
		AchievementID: ConvertToStrFrm(result.AchievementId),
		Events:        nil,
		ID:            "",
		Sku:           "",
	}

	for _, event := range triggerRequest.Events {
		var eventForSave = &model.TriggerToEventModel{
			Id:         uuid.New(),
			EventId:    ConvertStrFrmToUuid(event.EventID),
			TriggerId:  id,
			ProjectId:  project.Id,
			EventCount: event.Count,
		}

		var eventResult = h.triggerRepository.CreateTriggerToEventRelation(*eventForSave)

		responseVal.Events = append(responseVal.Events, &models.EventCount{
			Count:   eventResult.EventCount,
			EventID: ConvertToStrFrm(eventResult.EventId),
		})
	}
	return trigger.NewPostTriggerCreated().WithPayload(responseVal)
}

func ConvertToStrFrm(id uuid.UUID) strfmt.UUID {
	return strfmt.UUID(id.String())
}

func ConvertStrFrmToUuid(id strfmt.UUID) uuid.UUID {
	result, _ := uuid.Parse(id.String())
	return result
}

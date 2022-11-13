package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"ta4/mod/internal/server/models"
	"ta4/mod/internal/server/restapi/operations/achievement"
	"ta4/mod/repository"
)

type AddAchievementHandler struct {
	achievementRepository repository.AchievementRepository
	projectRepository     repository.ProjectRepository
}

func AddAchievementHandlerImpl() achievement.PostAchievementHandler {
	return AddAchievementHandler{
		achievementRepository: repository.AchievementRepository{},
		projectRepository:     repository.ProjectRepository{},
	}
}

func (h AddAchievementHandler) Handle(params achievement.PostAchievementParams) middleware.Responder {
	var achievementRequest = params.Body
	var project = h.projectRepository.GetProject(params.HTTPRequest.Header.Get("apiKey"))

	var achievementForSave = &repository.AchievementModel{
		Id:          uuid.New(),
		Sku:         achievementRequest.Sku,
		Description: achievementRequest.Description,
		ImageUrl:    achievementRequest.Image,
		ProjectId:   project.Id,
	}
	var result = h.achievementRepository.CreateAchievement(*achievementForSave)
	var responseVal = &models.Achievement{
		Description: result.Description,
		ID:          strfmt.UUID(result.Id.String()),
		Image:       result.ImageUrl,
		Sku:         result.Sku,
	}
	return achievement.NewPostAchievementCreated().WithPayload(responseVal)
}

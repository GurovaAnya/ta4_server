package main

import (
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"ta4/mod/config"
	"ta4/mod/internal/app"
	"ta4/mod/internal/server/restapi"
	"ta4/mod/internal/server/restapi/operations"
	"ta4/mod/internal/server/restapi/operations/achievement"
	"ta4/mod/internal/server/restapi/operations/event"
	"ta4/mod/internal/server/restapi/operations/project"
	"ta4/mod/internal/server/restapi/operations/trigger"
	achievementImpl "ta4/mod/services/handlers/achievement"
	eventImpl "ta4/mod/services/handlers/event"
	projectImpl "ta4/mod/services/handlers/project"
	triggerImpl "ta4/mod/services/handlers/trigger"
)

func main() {
	flag.Parse()
	c, err := config.Init()
	if err != nil {
		print("Failed init configuration. Error: %s", err)
		//logger.Logger.Errorf("Failed init configuration. Error: %s", err)
		os.Exit(1)
	}

	st, err := initDBConnection(&c)
	if err != nil {
		print("Failed init db. Error: %s", err)
		//logger.Logger.Errorf("Failed init db. Error: %s", err)
	}

	application := app.NewApplication(st, &c)

	// init handlers
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		print(err)
		//logger.Logger.Fatalln(err)
	}

	api := operations.NewSwaggerT4AAPI(swaggerSpec)

	server := restapi.NewServer(api)
	server.SetHandler(setupAPI(api, &application))
	server.EnabledListeners = []string{"http"}

	defer func() {
		if err := server.Shutdown(); err != nil {
			print("Failed serve shutdown. Error: %s", err)
			//logger.Logger.Errorf("Failed serve shutdown. Error: %s", err)
		}
	}()

	server.ConfigureAPI()
	server.Port = c.ServerPort
	server.Host = c.ServerHost
	if err := server.Serve(); err != nil {
		print("Failed serve requests. Error: %s", err)
		//logger.Logger.Errorf("Failed serve requests. Error: %s", err)
	}
}

func initDBConnection(c *config.Config) (*gorm.DB, error) {
	dbDSN := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s",
		c.PostgresUser,
		c.PostgresDbName,
		c.PostgresPassword,
	)
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		return nil, fmt.Errorf("failed init db connection. Error: %s", err)
	}
	return db, nil
}

func setupAPI(api *operations.SwaggerT4AAPI, application *app.Application) http.Handler {
	addProjectHandler := projectImpl.AddProjectHandlerImpl(application)
	api.ProjectAddProjectHandler = project.AddProjectHandlerFunc(addProjectHandler.Handle)

	getProjectHandler := projectImpl.GetProjectHandlerImpl(application)
	api.ProjectGetProjectHandler = project.GetProjectHandlerFunc(getProjectHandler.Handle)

	addAchievementHandler := achievementImpl.AddAchievementHandlerImpl(application)
	api.AchievementPostAchievementHandler = achievement.PostAchievementHandlerFunc(addAchievementHandler.Handle)

	addEventHandler := eventImpl.AddEventHandlerImpl(application)
	api.EventPostEventHandler = event.PostEventHandlerFunc(addEventHandler.Handle)

	addTriggerHandler := triggerImpl.AddTriggerHandlerImpl(application)
	api.TriggerPostTriggerHandler = trigger.PostTriggerHandlerFunc(addTriggerHandler.Handle)

	return api.Serve(nil)
}

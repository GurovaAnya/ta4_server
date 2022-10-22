package main

import (
	"flag"
	"github.com/go-openapi/loads"
	"ta4/mod/internal/app"
	"ta4/mod/internal/server/restapi"
	"ta4/mod/internal/server/restapi/operations"
)

func main() {
	flag.Parse()

	_ = app.NewApplication()

	// init handlers
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		print(err)
		//logger.Logger.Fatalln(err)
	}

	api := operations.NewSwaggerT4AAPI(swaggerSpec)

	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}

	defer func() {
		if err := server.Shutdown(); err != nil {
			print("Failed serve shutdown. Error: %s", err)
			//logger.Logger.Errorf("Failed serve shutdown. Error: %s", err)
		}
	}()

	server.ConfigureAPI()
	server.Port = 1025        //c.ServerPort
	server.Host = "localhost" //c.ServerHost
	if err := server.Serve(); err != nil {
		print("Failed serve requests. Error: %s", err)
		//logger.Logger.Errorf("Failed serve requests. Error: %s", err)
	}
}

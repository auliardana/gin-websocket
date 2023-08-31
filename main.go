package main

import (

	"GO-RESTAPI-GIN/routes"
	"GO-RESTAPI-GIN/databases"
	"github.com/gin-contrib/cors"
	// "github.com/swaggo/gin-swagger"
	ws "GO-RESTAPI-GIN/websocket"
    
	
)
// @title Swagger Data Mahasiswa
//	@version		1.0
// @termsOfService  http://swagger.io/terms/

//	@contact.email	auliardana@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html


//	@host		localhost:8000

//	@BasePath	/api/v1


func main()  {
	r:=routes.SetupRouter()
	databases.ConnectDatabase()
	r.Use(cors.Default())
	go ws.StartWebSocketServer()


	r.Run(":8000")
}
package routes

import(
	"GO-RESTAPI-GIN/controllers/mahasiswacontroller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_"GO-RESTAPI-GIN/docs"

	ws "GO-RESTAPI-GIN/websocket"


)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// Grup rute untuk mahasiswa
		mahasiswaGroup := v1.Group("/mahasiswas")
		{
			mahasiswaGroup.GET("", mahasiswacontroller.Index)
			mahasiswaGroup.GET("/:id", mahasiswacontroller.Show)
			mahasiswaGroup.POST("", mahasiswacontroller.Create)
			mahasiswaGroup.PUT("/:id", mahasiswacontroller.Update)
			mahasiswaGroup.DELETE("/:id", mahasiswacontroller.Delete)
		}

	}
	
	// WebSocket endpoint
	r.GET("/ws", ws.HandleWebSocket)

	// Endpoint untuk dokumentasi Swaggo
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
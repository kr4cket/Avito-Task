package handlers

import (
	"avitoTask/pkg/services"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "avitoTask/docs"
)

type Handler struct {
	services *services.Service
}

func New(service *services.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	service := router.Group("/service")
	{
		service.GET("/download/:id", h.downloadFile)
	}

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.GET("/all", h.getAllSegments)
			segments.GET("/get/:id", h.getSegmentById)
			segments.POST("/add", h.createSegment)
			segments.DELETE("/delete/:segmentName", h.deleteSegmentByName)
		}

		users := api.Group("/users")
		{
			users.GET("/active-segments/:id", h.getActiveUserSegments)
			users.POST("/history/:id", h.getHistory)
			users.POST("/change-segments/:id", h.changeUserSegments)
			users.POST("/expired-segments/:id", h.addExpiredSegment)
			users.POST("/add", h.addUser)
		}

	}

	return router
}

package main

import (
	"log"
	"net/http"
	"routefinder_golang/config"
	"routefinder_golang/handler"
	"routefinder_golang/repository"
	"routefinder_golang/service"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.AllowAll())

	db, err := config.ConnectToDb()
	if err != nil {
		log.Println("DB Connection Failed..")
		log.Println(err)
		return
	}

	mountainRepository := repository.NewMountainRepository(db)
	routeRepository := repository.NewRouteRepository(db)
	waypointsRepository := repository.NewWaypointsRepository(db)

	mountainService := service.NewMountainService(mountainRepository)
	routeService := service.NewRouteService(routeRepository)
	waypointsService := service.NewWaypointsService(waypointsRepository)

	mountainHandler := handler.NewMountainHandler(mountainService)
	routeHandler := handler.NewRouteHandler(routeService)
	waypointHandler := handler.NewWaypointsHandler(waypointsService)

	apiPublic := router.Group("/api")

	routeMountain := apiPublic.Group("/mountain")
	{
		routeMountain.GET("", mountainHandler.GetAllMountain)
		routeMountain.GET("/:id", mountainHandler.FindMountain)
		routeMountain.POST("", mountainHandler.CreateMountain)
		routeMountain.PUT("/:id", mountainHandler.EditMountain)
		routeMountain.DELETE("/:id", mountainHandler.DeleteMountain)
	}

	routeRoute := apiPublic.Group("/route")
	{
		routeRoute.GET("", routeHandler.GetAllRoute)
		routeRoute.GET("/:id", routeHandler.FindRoute)
		routeRoute.POST("", routeHandler.CreateRoute)
		routeRoute.PUT("/:id", routeHandler.EditRoute)
		routeRoute.DELETE("/:id", routeHandler.DeleteRoute)
	}

	routeWaypoints := apiPublic.Group("/waypoints")
	{
		routeWaypoints.GET("", waypointHandler.GetAllWaypoints)
		routeWaypoints.GET("/:id", waypointHandler.FindWaypoints)
		routeWaypoints.POST("", waypointHandler.CreateWaypoints)
		routeWaypoints.PUT("/:id", waypointHandler.EditWaypoints)
		routeWaypoints.DELETE("/:id", waypointHandler.DeleteWaypoints)
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message": "oke",
				"code":    "200",
				"status":  "oke",
				"data":    "",
			})

	})

	router.Run("localhost:8090")
	// go router.Run(":8090")
}

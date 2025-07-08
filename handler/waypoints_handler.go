package handler

import (
	"log"
	"net/http"
	"routefinder_golang/entity"
	"routefinder_golang/service"
	"routefinder_golang/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type waypointsHandler struct {
	waypointsService service.WaypointsService
}

func NewWaypointsHandler(service service.WaypointsService) *waypointsHandler {
	return &waypointsHandler{waypointsService: service}
}

func (handler *waypointsHandler) GetAllWaypoints(c *gin.Context) {
	result, err := handler.waypointsService.GetAllWaypoints()
	log.Println(result)
	if err != nil {
		util.APIListResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil, 0)
		return
	}
	util.APIListResponse(c, "Retrieve data success!", http.StatusOK, "ok", result, uint(len(result)))
}

func (handler *waypointsHandler) FindWaypoints(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	result, err := handler.waypointsService.FindWaypoints(id)
	log.Println(result)
	if err != nil {
		util.APIResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", result)
}

func (handler *waypointsHandler) CreateWaypoints(c *gin.Context) {
	var input entity.Waypoints
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	err = handler.waypointsService.CreateWaypoints(input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}

func (handler *waypointsHandler) EditWaypoints(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	var input entity.Waypoints
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err = handler.waypointsService.EditWaypoints(id, input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

func (handler *waypointsHandler) DeleteWaypoints(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err := handler.waypointsService.DeleteWaypoints(id)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

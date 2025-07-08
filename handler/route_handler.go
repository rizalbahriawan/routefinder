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

type routeHandler struct {
	routeService service.RouteService
}

func NewRouteHandler(service service.RouteService) *routeHandler {
	return &routeHandler{routeService: service}
}

func (handler *routeHandler) GetAllRoute(c *gin.Context) {
	result, err := handler.routeService.GetAllRoute()
	log.Println(result)
	if err != nil {
		util.APIListResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil, 0)
		return
	}
	util.APIListResponse(c, "Retrieve data success!", http.StatusOK, "ok", result, uint(len(result)))
}

func (handler *routeHandler) FindRoute(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	result, err := handler.routeService.FindRoute(id)
	log.Println(result)
	if err != nil {
		util.APIResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", result)
}

func (handler *routeHandler) CreateRoute(c *gin.Context) {
	var input entity.Route
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	err = handler.routeService.CreateRoute(input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}

func (handler *routeHandler) EditRoute(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	var input entity.Route
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err = handler.routeService.EditRoute(id, input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

func (handler *routeHandler) DeleteRoute(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err := handler.routeService.DeleteRoute(id)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

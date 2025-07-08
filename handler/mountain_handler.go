package handler

import (
	"log"
	"net/http"
	"routefinder_golang/request"
	"routefinder_golang/service"
	"routefinder_golang/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type mountainHandler struct {
	mountainService service.MountainService
}

func NewMountainHandler(service service.MountainService) *mountainHandler {
	return &mountainHandler{mountainService: service}
}

func (handler *mountainHandler) GetAllMountain(c *gin.Context) {
	result, err := handler.mountainService.GetAllMountain()
	log.Println(result)
	if err != nil {
		util.APIListResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil, 0)
		return
	}
	util.APIListResponse(c, "Retrieve data success!", http.StatusOK, "ok", result, uint(len(result)))
}

func (handler *mountainHandler) FindMountain(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	result, err := handler.mountainService.FindMountain(id)
	log.Println(result)
	if err != nil {
		util.APIResponse(c, "Data tidak ditemukan", http.StatusBadRequest, "Data tidak ditemukan", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", result)
}

func (handler *mountainHandler) CreateMountain(c *gin.Context) {
	var input request.MountainRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	err = handler.mountainService.CreateMountain(input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}

func (handler *mountainHandler) EditMountain(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	var input request.MountainRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err = handler.mountainService.EditMountain(id, input)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

func (handler *mountainHandler) DeleteMountain(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err := handler.mountainService.DeleteMountain(id)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}

	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"grupo-siete-go/internal/paciente"
	"net/http"
	"strconv"
)

type PacienteHandler struct {
	service paciente.Service
}

func NewHandler(service paciente.Service) *PacienteHandler {
	return &PacienteHandler{service: service}
}

func (h *PacienteHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	paciente, err := h.service.GetByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, paciente)
}
func (h *PacienteHandler) Save(ctx *gin.Context) {
	var pacienteInput paciente.Paciente
	err := ctx.ShouldBind(&pacienteInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	paciente, err := h.service.Save(pacienteInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, paciente)
}

func (h *PacienteHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var pacienteInput paciente.Paciente
	err = ctx.ShouldBind(&pacienteInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	paciente, err := h.service.Update(ID, pacienteInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, paciente)
}
func (h *PacienteHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.Delete(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Paciente deleted")
}
package handler

import (
	"grupo-siete-go/internal/paciente"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	service paciente.Service
}

func NewPacienteHandler(service paciente.Service) *PacienteHandler {
	return &PacienteHandler{service: service}
}

// GetPacienteByID godoc
// @Summary      Gets a paciente by id
// @Description  Gets a paciente by id from the repository
// @Tags         paciente
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} paciente.Paciente
// @Router       /pacientes/{id} [get]
func (h *PacienteHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	paciente, err := h.service.GetByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, paciente)
}

// Save Paciente godoc
// @Summary      Saves a paciente
// @Description  Saves a paciente into the repository
// @Tags         paciente
// @Produce      json
// @Param        paciente body request
// @Success      200 {object} paciente.Paciente
// @Router       /pacientes [post]
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

// Replace Paciente godoc
// @Summary      Replaces an paciente
// @Description  Replaces an existing paciente from the repository
// @Tags         paciente
// @Produce      json
// @Param        paciente body request
// @Success      200 {object} paciente.Paciente
// @Router       /pacientes [put]

// Update Paciente godoc
// @Summary      Udpates an paciente
// @Description  Updates an existing paciente from the repository with one o more features
// @Tags         paciente
// @Produce      json
// @Param        id path string true "ID" - paciente body request
// @Success      200 {object} paciente.Paciente
// @Router       /pacientes [patch]
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

// Delete Paciente godoc
// @Summary      Deletes an paciente
// @Description  Deletes an existing paciente from the repository
// @Tags         paciente
// @Produce      string
// @Param        id path string true "ID"
// @Success      200 {object} web.response
// @Router       /pacientes [delete]
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

package handler

import (
	modelo "grupo-siete-go/internal/turno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TurnoCreator interface {
	Save(turno modelo.Turno) (modelo.Turno, error)
}

type TurnoUpdater interface {
	Update(id int, turno modelo.Turno) (modelo.Turno, error)
	Replace(turno modelo.Turno) (modelo.Turno, error)
}

type TurnoGetter interface {
	GetByID(id int) (modelo.Turno, error)
	GetByDNI(dni string) ([]modelo.Turno, error)
}

type TurnoEliminator interface {
	Delete(id int) (string, error)
}

type TurnoHandler struct {
	turnoCreator    TurnoCreator
	turnoGetter     TurnoGetter
	TurnoUpdater    TurnoUpdater
	TurnoEliminator TurnoEliminator
}

func NewTurnoHandler(creator TurnoCreator, getter TurnoGetter, updater TurnoUpdater, eliminator TurnoEliminator) *TurnoHandler {
	return &TurnoHandler{turnoCreator: creator, turnoGetter: getter, TurnoUpdater: updater, TurnoEliminator: eliminator}
}

// Save Turno godoc
// @Summary      Saves a turno
// @Description  Saves a turno into the repository
// @Tags         turno
// @Produce      json
// @Param        turno body turno.Turno true "Saves a turno"
// @Success      200 {object} turno.Turno
// @Router       /turnos [post]
func (th *TurnoHandler) Save(ctx *gin.Context) {
	// obtengo el turno del contexto
	var turno modelo.Turno
	if err := ctx.ShouldBind(&turno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// guardo el turno
	savedTurno, err := th.turnoCreator.Save(turno)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusCreated, savedTurno)
}

// GetTurnoByID godoc
// @Summary      Gets a turno by id
// @Description  Gets a turno by id from the repository
// @Tags         turno
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} turno.Turno
// @Router       /turnos/{id} [get]
func (th *TurnoHandler) GetByID(ctx *gin.Context) {
	// obtengo el id del turno a buscar
	strID := ctx.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// busco el turno
	turno, err := th.turnoGetter.GetByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusOK, turno)
}

// Replace Turno godoc
// @Summary      Replaces a turno
// @Description  Replaces an existing turno from the repository
// @Tags         turno
// @Produce      json
// @Param        turno body turno.Turno true "Replaces a turno"
// @Success      200 {object} turno.Turno
// @Router       /turnos [put]
func (th *TurnoHandler) Replace(ctx *gin.Context) {
	// obtengo el turno
	var turno modelo.Turno
	if err := ctx.ShouldBind(&turno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTurno, err := th.TurnoUpdater.Replace(turno)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusOK, updatedTurno)
}

// Update Turno godoc
// @Summary      Udpates an turno
// @Description  Updates an existing turno from the repository with one o more features
// @Tags         turno
// @Produce      json
// @Param        id path string true "ID" - turno body request
// @Success      200 {object} turno.Turno
// @Router       /turnos [patch]
func (th *TurnoHandler) Update(ctx *gin.Context) {
	// obtengo el id
	strID := ctx.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// obtengo el turno
	var turno modelo.Turno
	if err := ctx.ShouldBind(&turno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTurno, err := th.TurnoUpdater.Update(ID, turno)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusOK, updatedTurno)
}

// Delete Turno godoc
// @Summary      Deletes an turno
// @Description  Deletes an existing turno from the repository
// @Tags         turno
// @Produce      json
// @Param        id path string true "ID"
// @Success      200
// @Router       /turnos [delete]
func (th *TurnoHandler) Delete(ctx *gin.Context) {
	// obtengo el id
	strID := ctx.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// elimino el turno
	res, err := th.TurnoEliminator.Delete(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusOK, res)
}

// GetTurnoByDNI godoc
// @Summary      Gets a turno by paciente DNI
// @Description  Gets a turno by paciente dni from the repository
// @Tags         turno
// @Produce      json
// @Param        dni query string true "DNI"
// @Success      200 {object} turno.Turno
// @Router       /turnos [get]
func (th *TurnoHandler) GetByDNI(ctx *gin.Context) {
	// obtengo el dni
	dni := ctx.Query("dni")

	// traigo los turnos
	turnos, err := th.turnoGetter.GetByDNI(dni)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// OK
	ctx.JSON(http.StatusOK, turnos)
}

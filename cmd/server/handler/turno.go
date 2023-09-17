package handler

import (
	"github.com/gin-gonic/gin"
	modelo "grupo-siete-go/internal/turno"
	"net/http"
	"strconv"
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
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

type TurnoGetter interface {
	GetByID(id int) (modelo.Turno, error)
}

type TurnoHandler struct {
	turnoCreator TurnoCreator
	turnoGetter  TurnoGetter
}

func NewTurnoHandler(creator TurnoCreator, getter TurnoGetter) *TurnoHandler {
	return &TurnoHandler{turnoCreator: creator, turnoGetter: getter}
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

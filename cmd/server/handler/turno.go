package handler

import (
	"github.com/gin-gonic/gin"
	modelo "grupo-siete-go/internal/turno"
	"net/http"
)

type TurnoCreator interface {
	Save(turno modelo.Turno) (modelo.Turno, error)
}

type TurnoHandler struct {
	turnoCreator TurnoCreator
}

func NewTurnoHandler(creator TurnoCreator) *TurnoHandler {
	return &TurnoHandler{turnoCreator: creator}
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

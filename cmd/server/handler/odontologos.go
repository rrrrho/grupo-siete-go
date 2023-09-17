package handler

import (
	"grupo-siete-go/internal/odontologo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OdontologoGetter interface {
	GetByID(id int) odontologo.Odontologo
}

type OdontologoCreator interface {
	ModifyById(id int)
}

type OdontologosHandler struct {
	odontologoGetter  OdontologoGetter
	odontologoCreator OdontologoCreator
}

func NewOdontologoHandler(getter OdontologoGetter, creator OdontologoCreator) *OdontologosHandler {
	return &OdontologosHandler{
		odontologoGetter:  getter,
		odontologoCreator: creator,
	}
}

func (odontologoHandler *OdontologosHandler) GetOdontologoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	odontologo, err := ph.productsGetter.GetByID(id)

}

// FALTA TERMINAR

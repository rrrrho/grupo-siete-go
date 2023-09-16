package handler

import (
	"grupo-siete-go/internal/odontologo"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type OdontologoGetter interface {
	GetByID(id int) (odontologo.Odontologo)
}

type OdontologoCreator interface {
	ModifyById(id int, )
}

// --------------
// FALTA TERMINAR
// --------------

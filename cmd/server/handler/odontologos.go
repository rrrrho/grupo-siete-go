package handler

import (
	"grupo-siete-go/internal/odontologo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OdontologoGetter interface {
	GetByID(id int) (odontologo.Odontologo, error)
}

type OdontologoCreator interface {
	Save(odontologo odontologo.Odontologo) (odontologo.Odontologo, error)
}

type OdontologoUpdater interface {
	Update(id int, odontologo odontologo.Odontologo) (odontologo.Odontologo, error)
	Replace(odontologo odontologo.Odontologo) (odontologo.Odontologo, error)
}

type OdontologoEliminator interface {
	Delete(id int) (string, error)
}

type OdontologosHandler struct {
	odontologoGetter     OdontologoGetter
	odontologoCreator    OdontologoCreator
	odontologoUpdater    OdontologoUpdater
	odontologoEliminator OdontologoEliminator
}

func NewOdontologoHandler(getter OdontologoGetter, creator OdontologoCreator, updater OdontologoUpdater, eliminator OdontologoEliminator) *OdontologosHandler {
	return &OdontologosHandler{
		odontologoGetter:     getter,
		odontologoCreator:    creator,
		odontologoUpdater:    updater,
		odontologoEliminator: eliminator,
	}
}

// GetOdontologoByID godoc
// @Summary      Gets a odontologo by id
// @Description  Gets a odontologo by id from the repository
// @Tags         odontologos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} odontolog.Odontologo
// @Router       /odontologos/{id} [get]
func (h *OdontologosHandler) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	odontologo, err := h.odontologoGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Odontologo not found"})
		return
	}
	ctx.JSON(http.StatusOK, odontologo)
}

// Save Odontologo godoc
// @Summary      Saves a odontologo
// @Description  Saves a odontologo into the repository
// @Tags         odontologos
// @Produce      json
// @Param        odontologo body request
// @Success      200 {object} odontolog.Odontologo
// @Router       /odontologos [post]
func (h *OdontologosHandler) Save(ctx *gin.Context) {
	// obtengo el odontologo del contexto
	var odontologo odontologo.Odontologo
	if err := ctx.ShouldBind(&odontologo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// guardo el odontologo
	savedOdontologo, err := h.odontologoCreator.Save(odontologo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// OK
	ctx.JSON(http.StatusCreated, savedOdontologo)
}

// Replace Odontologo godoc
// @Summary      Replaces an odontologo
// @Description  Replaces an existing odontologo from the repository
// @Tags         odontologos
// @Produce      json
// @Param        odontologo body request
// @Success      200 {object} odontolog.Odontologo
// @Router       /odontologos [put]
func (h *OdontologosHandler) Replace(ctx *gin.Context) {
	// obtengo el odontologo
	var odontologo odontologo.Odontologo
	if err := ctx.ShouldBind(&odontologo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOdontologo, err := h.odontologoUpdater.Replace(odontologo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// OK
	ctx.JSON(http.StatusOK, updatedOdontologo)
}

// Update Odontologo godoc
// @Summary      Udpates an odontologo
// @Description  Updates an existing odontologo from the repository with one o more features
// @Tags         odontologos
// @Produce      json
// @Param        id path string true "ID" - odontologo body request
// @Success      200 {object} odontolog.Odontologo
// @Router       /odontologos [patch]
func (h *OdontologosHandler) Update(ctx *gin.Context) {
	// obtengo el id
	strID := ctx.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// obtengo el odontologo
	var odontologo odontologo.Odontologo
	if err := ctx.ShouldBind(&odontologo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOdontologo, err := h.odontologoUpdater.Update(ID, odontologo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// OK
	ctx.JSON(http.StatusOK, updatedOdontologo)
}

// Delete Odontologo godoc
// @Summary      Deletes an odontologo
// @Description  Deletes an existing odontologo from the repository
// @Tags         odontologos
// @Produce      string
// @Param        id path string true "ID"
// @Success      200 {object} web.response
// @Router       /odontologos [delete]
func (h *OdontologosHandler) Delete(ctx *gin.Context) {
	// obtengo el id
	strID := ctx.Param("id")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// elimino el odontologo
	res, err := h.odontologoEliminator.Delete(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// OK
	ctx.JSON(http.StatusOK, res)
}

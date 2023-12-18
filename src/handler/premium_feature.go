package handler

import (
	"net/http"
	"strconv"
	"DatingApp/src/filter"
	"DatingApp/src/formatter"
	"DatingApp/src/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags PremiumFeature
// @Security ApiKeyAuth
// @Param paging query filter.Paging[filter.PremiumFeatureFilter] false "paging"
// @Param filter query filter.PremiumFeatureFilter false "filter"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /premium-feature/ [GET]
func (h *handler) GetPremiumFeature(ctx *gin.Context) {
	var filter filter.Paging[filter.PremiumFeatureFilter]
	filter.SetDefault()

	if err := h.BindParams(ctx, &filter); err != nil {
		response := models.APIResponse("Get PremiumFeature Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	premiumfeatures, count, err := h.service.PremiumFeature.Get(ctx, filter)
	if err != nil {
		response := models.APIResponse("Get PremiumFeature Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	paginatedItems := formatter.PaginatedItems{}
	paginatedItems.Format(filter.Page, float64(len(premiumfeatures)), float64(count), float64(filter.Take), premiumfeatures)

	response := models.APIResponse("Get PremiumFeature Success", http.StatusOK, "Success", paginatedItems, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags PremiumFeature
// @Security ApiKeyAuth
// @Param models body models.PremiumFeatureInput true "models"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /premium-feature/ [POST]
func (h *handler) CreatePremiumFeature(ctx *gin.Context) {
	var input models.Query[models.PremiumFeatureInput]

	if err := ctx.ShouldBindJSON(&input.Model); err != nil {
		response := models.APIResponse("Create PremiumFeature Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.PremiumFeature.Create(ctx, input); err != nil {
		response := models.APIResponse("Create PremiumFeature Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Create PremiumFeature Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags PremiumFeature
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Param models body models.PremiumFeatureInput true "models"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /premium-feature/{id} [PUT]
func (h *handler) UpdatePremiumFeature(ctx *gin.Context) {
	var input models.Query[models.PremiumFeatureInput]

	if err := ctx.ShouldBindJSON(&input.Model); err != nil {
		response := models.APIResponse("Update PremiumFeature Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Update PremiumFeature Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.PremiumFeature.Update(ctx, input, id); err != nil {
		response := models.APIResponse("Update PremiumFeature Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Update PremiumFeature Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags PremiumFeature
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /premium-feature/{id} [DELETE]
func (h *handler) DeletePremiumFeature(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Delete PremiumFeature Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.PremiumFeature.Delete(ctx, id); err != nil {
		response := models.APIResponse("Delete PremiumFeature Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Delete PremiumFeature Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

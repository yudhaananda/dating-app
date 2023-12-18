package handler

import (
	"DatingApp/src/filter"
	"DatingApp/src/formatter"
	"DatingApp/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags UserActivity
// @Security ApiKeyAuth
// @Param paging query filter.Paging[filter.UserActivityFilter] false "paging"
// @Param filter query filter.UserActivityFilter false "filter"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user-activity/ [GET]
func (h *handler) GetUserActivity(ctx *gin.Context) {
	var filter filter.Paging[filter.UserActivityFilter]
	filter.SetDefault()

	if err := h.BindParams(ctx, &filter); err != nil {
		response := models.APIResponse("Get UserActivity Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	useractivitys, count, err := h.service.UserActivity.Get(ctx, filter)
	if err != nil {
		response := models.APIResponse("Get UserActivity Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	paginatedItems := formatter.PaginatedItems{}
	paginatedItems.Format(filter.Page, float64(len(useractivitys)), float64(count), float64(filter.Take), useractivitys)

	response := models.APIResponse("Get UserActivity Success", http.StatusOK, "Success", paginatedItems, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags UserActivity
// @Security ApiKeyAuth
// @Param activity path string true "activity"
// @Param targetUserId body models.UserActivityInputJson true "passed or liked userId"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user-activity/{activity} [POST]
func (h *handler) CreateUserActivity(ctx *gin.Context) {
	activity := ctx.Param("activity")
	var targetUserId models.UserActivityInputJson
	var input models.Query[models.UserActivityInput]

	if err := ctx.ShouldBindJSON(&targetUserId); err != nil {
		response := models.APIResponse("Create UserActivity Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if activity == "pass" {
		input.Model.PassedUserId = targetUserId.TargetUserId
		input.Model.UserId = int(ctx.Value(models.UserKey).(models.User).Id)
	} else if activity == "like" {
		input.Model.LikedUserId = targetUserId.TargetUserId
		input.Model.UserId = int(ctx.Value(models.UserKey).(models.User).Id)
	}

	if err := h.service.UserActivity.Create(ctx, input); err != nil {
		response := models.APIResponse("Create UserActivity Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Create UserActivity Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags UserActivity
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Param models body models.UserActivityInput true "models"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user-activity/{id} [PUT]
func (h *handler) UpdateUserActivity(ctx *gin.Context) {
	var input models.Query[models.UserActivityInput]

	if err := ctx.ShouldBindJSON(&input.Model); err != nil {
		response := models.APIResponse("Update UserActivity Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Update UserActivity Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.UserActivity.Update(ctx, input, id); err != nil {
		response := models.APIResponse("Update UserActivity Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Update UserActivity Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags UserActivity
// @Security ApiKeyAuth
// @Param id path integer true "id"
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /user-activity/{id} [DELETE]
func (h *handler) DeleteUserActivity(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := models.APIResponse("Delete UserActivity Failed", http.StatusUnprocessableEntity, "Failed", nil, err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := h.service.UserActivity.Delete(ctx, id); err != nil {
		response := models.APIResponse("Delete UserActivity Failed", http.StatusInternalServerError, "Failed", nil, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.APIResponse("Delete UserActivity Success", http.StatusOK, "Success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

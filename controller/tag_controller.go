package controller

import (
	"ginFramework/data/request"
	"ginFramework/data/response"
	"ginFramework/helper"
	"ginFramework/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsContoller struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsContoller {
	return &TagsContoller{
		tagsService: service,
	}
}

// CreateTags godoc
// @Summary    Create tags
// @Description Save tags data in Db.
// @Param        Tags body request.CreateTagsRequest true "Create tags"
// @Produce      application/json
// @Tags         tags
// @Success      200 {object} response.Response{}
// @Router       /tags [port]
func (controller *TagsContoller) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags godoc
// @Summary   Update tags
// @Description Update tags data.
// @Param       tagId path string true "update tags by id"
// @Param       tags body request.CreateTagsRequest true "Update tags"
// @Tags      tags
// @Produce   application/json
// @Success   200 {object} response.Response{}
// @Router    /tags/{tagId} [patch]

func (controller *TagsContoller) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeleteTags godoc
// @Summary  Delete tags
// @Description  Remove tags data by id.
// @Produce   application/json
// @Tags    tags
// @Success  200 {object} response.Response{}
// @Router /tag/{tagID} [delete]
func (controller *TagsContoller) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIdTags  godoc
// @Summary      Get Single tags by id.
// @Param        tagId path string true "update tags by id"
// @Description  Return the tahs shoes tagId value mathes id.
// @Produce      application/json
// @Tags         tags
// @Success 200  {object} response.Response{}
// @Router       /tag/{tagId} [get]
func (controller *TagsContoller) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags   godoc
// @Summary Get All tags.
// @Description Return list of tags.
// @Tags    tags
// @Sucess   200 {object} response.Response{}
// @Router   /tags [get]
func (controller *TagsContoller) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

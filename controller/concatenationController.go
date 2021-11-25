package controller

import (
	"net/http"
	"video_api/common"
	"video_api/model"
	"video_api/response"
	"video_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func HandleConcatenationRequeset(ctx *gin.Context) {
	var obj model.RcvObj

	if err := ctx.ShouldBindJSON(&obj); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := common.GetDB()
	//parse timeline sequence
	txtPath, err := utils.ParseTimelineSequense(obj.Timeline, db)
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}

	videoPath, err := utils.ConcatenateViedoFiles(txtPath)
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}

	respString, err := utils.UploadVideoFile(obj.Project, videoPath)
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	response.Success(ctx, gin.H{"onLineUrl": "http://" + viper.GetString("server.address") + ":" + viper.GetString("server.port") + videoPath[1:], "response": respString}, "rendering succeed")

}

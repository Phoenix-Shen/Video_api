package controller

import (
	"fmt"
	"net/http"
	"video_api/common"
	"video_api/dto"
	"video_api/model"
	"video_api/response"
	"video_api/utils"

	"github.com/gin-gonic/gin"
)

func HandleEffectRequest(ctx *gin.Context) {
	//get file and effect
	file, err := ctx.FormFile("file")
	effect := ctx.PostForm("effect")

	//if cant get file, return error
	if err != nil {
		response.Fail(ctx, nil, "ctx.Formfile error")
		return
	}
	//mapping to objects
	videoeft := model.VideoEffect{
		File:      file,
		Effect:    effect,
		LocalPath: fmt.Sprintf("./static/videos/%s", utils.GetUniqueVideoName(file.Filename)),
	}

	//save video

	ctx.SaveUploadedFile(file, videoeft.LocalPath)

	//process the video
	if videoeft.Effect == "STYLE" {
		videoPath, err := utils.HandleTransformStyleCall(videoeft.LocalPath)
		if err != nil {
			response.Fail(ctx, nil, err.Error())
			return
		}
		videoeft.PathAfterProcess = videoPath
		err = utils.GetFirstFrame(videoPath)
		if err != nil {
			response.Fail(ctx, nil, err.Error())
			return
		}
		videoeft.CoverPath = fmt.Sprintf("./static/images/%s.jpg", utils.GetFileNamewithoutExt(videoPath))
	} else if videoeft.Effect == "WATERMARK" {
		videoPath, err := utils.HandleWatermarkCall(videoeft.LocalPath)
		if err != nil {
			response.Fail(ctx, nil, err.Error())
			return
		}
		videoeft.PathAfterProcess = videoPath
		//get the first frame of video and will be used as cover
		err = utils.GetFirstFrame(videoPath)
		if err != nil {
			response.Fail(ctx, nil, err.Error())
			return
		}
		videoeft.CoverPath = fmt.Sprintf("./static/images/%s.jpg", utils.GetFileNamewithoutExt(videoPath))
	} else if videoeft.Effect == "DEMO" {
		videoeft.LocalPath = "./static/videos/flower.mp4"
		videoeft.CoverPath = "./static/images/flower.png"
		videoeft.PathAfterProcess = "./static/videos/flower.mp4"
		videoeft.Effect = "STYLE"
	} else {
		err = utils.GetFirstFrame(videoeft.LocalPath)
		if err != nil {
			response.Fail(ctx, nil, err.Error())
			return
		}
		videoeft.CoverPath = fmt.Sprintf("./static/images/%s.jpg", utils.GetFileNamewithoutExt(videoeft.LocalPath))
		videoeft.PathAfterProcess = videoeft.LocalPath
		videoeft.Effect = "none"
	}

	//save the object to database
	db := common.GetDB()
	db.Create(&videoeft)
	//response
	response.Response(ctx, http.StatusOK, 200, gin.H{"obj": dto.ToVideoEffectDto(videoeft)}, "operation succeed")
}

package dto

import (
	"video_api/model"

	"github.com/spf13/viper"
)

type VideoEffectDto struct {
	IsOnline  bool
	IsFx      bool
	Effect    string
	OnLineId  int
	OnLineUrl string
	CoverUrl  string
}

func ToVideoEffectDto(videoEffect model.VideoEffect) VideoEffectDto {
	returnAddress := "http://" + viper.GetString("server.address") + ":" + viper.GetString("server.port")
	coveUrl := ""
	if videoEffect.CoverPath == "none" {
		coveUrl = "none"
	} else {
		coveUrl = returnAddress + videoEffect.CoverPath[1:]
	}
	return VideoEffectDto{
		IsOnline:  true,
		IsFx:      true,
		Effect:    videoEffect.Effect,
		OnLineId:  int(videoEffect.ID),
		OnLineUrl: returnAddress + videoEffect.PathAfterProcess[1:],
		CoverUrl:  coveUrl,
	}
}

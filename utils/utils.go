package utils

import (
	"errors"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
	"video_api/model"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/spf13/viper"
)

//get unique video name
func GetUniqueVideoName(originalFileName string) string {
	//get unique id using ksuid
	id := ksuid.New()
	//parse file extension
	fileExt := path.Ext(originalFileName)
	return id.String() + fileExt
}

//read in config
func InitConifg() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

//get random string
func RandomString(length int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	rand.Seed(time.Now().Unix())
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//get filename with out extension
func GetFileNamewithoutExt(fileFullName string) string {
	fileName := path.Base(fileFullName)
	fileExt := path.Ext(fileName)
	return strings.TrimSuffix(fileName, fileExt)
}

//parse time line sequence
func ParseTimelineSequense(timelines []model.Timeline, db *gorm.DB) (string, error) {
	//db search record
	var localPaths string = ""

	for _, timeline := range timelines {
		var vef model.VideoEffect
		vef.ID = uint(timeline.OnLineId)
		db.Where("ID = ?", vef.ID).First(&vef)
		localPaths += "file " + path.Base(vef.LocalPath) + "\n"
		if vef.LocalPath == "" {
			return "", errors.New("can't find the id in one of the timelines, pls check the data")
		}
	}
	//write paths to txt
	uniqueName := "./static/videos/" + GetUniqueVideoName("a.txt")
	/*pwd, err := os.Getwd()
	if err != nil {

		return "", err
	}
	joinedName := filepath.Join(pwd, uniqueName)*/
	file, err := os.OpenFile(uniqueName, os.O_CREATE, 0777)
	if err != nil {

		return "", err
	}

	_, err = file.Write([]byte(localPaths))
	if err != nil {
		return "", err
	}
	defer file.Close()
	return uniqueName, nil

}

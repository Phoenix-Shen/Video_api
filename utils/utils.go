package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
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
		uintID, err := strconv.Atoi(timeline.OnLineId)
		vef.ID = uint(uintID)
		if err != nil {
			return "", err
		}
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

//upload video file to the server
func UploadVideoFile(videoName string, videoNetLocation string) (string, error) {
	requestBody := fmt.Sprintf(`{
	"controller": "1612779773437",
	"userId": "1612779773437",
	"videoName": "%s",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxNjEyNzc5NzczNDM3IiwiaWF0IjoxNjEyNzgwMDE0fQ.J27ujArwYmr2b7Muv2wI3FEs1YbXO8Ce2llju6dMzjo",
	"url": "%s"
	}`, videoName, "http://"+viper.GetString("server.address")+":"+viper.GetString("server.port")+videoNetLocation[1:])

	var jsonStr = []byte(requestBody)

	url := "https://qcmt57.fn.thelarkcloud.com/createVideo"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

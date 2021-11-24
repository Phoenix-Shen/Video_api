package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"video_api/common"
	"video_api/router"
	"video_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//read in configuration
	utils.InitConifg()
	//init logs function
	gin.ForceConsoleColor()
	t := time.Now()
	f, _ := os.Create(fmt.Sprintf("./logs/%d-%d-%d %d-%d-%d.log", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
	gin.DefaultWriter = io.MultiWriter(f, os.Stderr)
	//init database connection
	common.DB = common.InitDB()
	defer common.DB.Close()
	//init routers
	r := gin.Default()
	r = router.InitRouter(r)
	port := viper.GetString("server.port")
	//run application
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}

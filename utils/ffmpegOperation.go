package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func SplitVideoIntoFrames(videoPath string) (string, error) {
	//create New unique folder
	folderName := "dataCache_" + RandomString(10)
	os.Mkdir(folderName, os.ModePerm)
	//execute command
	cmd := exec.Command("./utils/ffmpeg/bin/ffmpeg.exe", "-i", videoPath, fmt.Sprintf("./%s/frames_%%05d.jpg", folderName))
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	//error handling
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
		return "", err
	}
	//打印命令行的标准输出
	fmt.Println(cmd.Stdout.(*bytes.Buffer).String())

	return folderName, nil
}

//this operation requires python environment and PyTorch
func TransformStyle(folderName string) error {
	//open directory
	f, err := os.Open(folderName)
	if err != nil {
		return err
	}
	files, err := f.ReadDir(-1)
	f.Close()
	if err != nil {
		return err
	}
	//Enumerate files and process it
	for _, file := range files {
		fileFullPath := "./" + folderName + "/" + file.Name()
		fileFullPathOut := "./" + folderName + "/transformed_" + file.Name()
		//execute python shell
		cmd := exec.Command("python", "./utils/neural-style/main.py", "stylize", "--use_gpu=False", fmt.Sprintf("--content_path=%s", fileFullPath), fmt.Sprintf("--result_path=%s", fileFullPathOut))
		cmd.Stdout = &bytes.Buffer{}
		cmd.Stderr = &bytes.Buffer{}
		err := cmd.Run()
		//error handling
		if err != nil {
			//打印程序中的错误以及命令行标准错误中的输出
			fmt.Println(err)
			fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
			return err
		}
		//打印命令行的标准输出
		fmt.Println(cmd.Stdout.(*bytes.Buffer).String())

	}

	return nil
}

func CompositeVideo(folderName string) (string, error) {
	//execute command
	uniqueVideoName := fmt.Sprintf("./static/videos/%s", GetUniqueVideoName("1.mp4"))
	cmd := exec.Command("./utils/ffmpeg/bin/ffmpeg.exe", "-i", fmt.Sprintf("./%s/transformed_frames_%%05d.jpg", folderName), uniqueVideoName)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	//error handling
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
		return "", err
	}
	//打印命令行的标准输出
	fmt.Println(cmd.Stdout.(*bytes.Buffer).String())
	//delete cache folder
	os.RemoveAll(folderName)
	return uniqueVideoName, nil
}

func HandleTransformStyleCall(videoPath string) (string, error) {
	//split video into frame
	returnFolderPath, err := SplitVideoIntoFrames(videoPath)
	if err != nil {
		return "", err
	}

	// and then transform style frame by frame
	err = TransformStyle(returnFolderPath)
	if err != nil {
		return "", err
	}

	videoPath, err = CompositeVideo(returnFolderPath)
	if err != nil {
		return "", err
	}
	return videoPath, nil
}

func HandleWatermarkCall(videoPath string) (string, error) {
	uniqueVideoName := fmt.Sprintf("./static/videos/%s", GetUniqueVideoName("1.mp4"))
	cmd := exec.Command("./utils/ffmpeg/bin/ffmpeg.exe", "-i", videoPath, "-i", "./static/images/laugh.png", "-filter_complex", "overlay=W-w", uniqueVideoName)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	//error handling
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
		return "", err
	}
	//打印命令行的标准输出
	fmt.Println(cmd.Stdout.(*bytes.Buffer).String())
	return uniqueVideoName, nil
}

//clip the first frame of the video
func GetFirstFrame(videoPath string) error {
	fileNamewithoutExt := GetFileNamewithoutExt(videoPath)
	filter := "select=eq(pict_type\\,I)"
	//ffmpeg.exe  -i C:\users\ssk\desktop\result.mp4 -vf "select=eq(pict_type\,I)" -vframes 1 -vsync vfr -qscale:v 2 -f image2 ./%08d.jpg
	cmd := exec.Command("./utils/ffmpeg/bin/ffmpeg.exe", "-i", videoPath, "-vf", filter, "-vframes", "1", "-vsync", "vfr", "-qscale:v", "2", "-f", "image2", fmt.Sprintf("./static/images/%s.jpg", fileNamewithoutExt))

	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	//error handling
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
		return err
	}
	//打印命令行的标准输出
	fmt.Println(cmd.Stdout.(*bytes.Buffer).String())
	return nil
}

//contatenate videos accroding to the text file
func ConcatenateViedoFiles(txtPath string) (string, error) {
	// ffmpeg -f concat -i filelist.txt -c copy output.mkv
	//./utils/ffmpeg/bin/ffmpeg.exe -f concat -safe 0 -i ./static/videos/21LitFPQjdVli8SJyJkTtVh79xC.txt -c copy a.mp4
	uniqueVideoName := fmt.Sprintf("./static/videos/%s", GetUniqueVideoName("1.mp4"))
	cmd := exec.Command("./utils/ffmpeg/bin/ffmpeg.exe", "-f", "concat", "-safe", "0", "-i", txtPath, "-c", "copy", uniqueVideoName)

	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	//error handling
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(cmd.Stderr.(*bytes.Buffer).String())
		return "", err
	}
	//打印命令行的标准输出
	fmt.Println(cmd.Stdout.(*bytes.Buffer).String())
	//delete cached text folder
	os.Remove(txtPath)
	return uniqueVideoName, nil
}

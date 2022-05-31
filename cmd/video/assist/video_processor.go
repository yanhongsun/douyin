package assist

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"gocv.io/x/gocv"
)

//将视频数据保存
func SaveVideo(videoPath string, video []byte) error {
	ioutil.WriteFile(videoPath, video, 0666)
	//TODO异常处理
	return nil
}

//从视频数据中提取封面保存路径
func GetCover(videoCoverPath string, videoPath string) (err error) {
	var imag image.Image
	var er error
	imag, er = GetVideoMoment(videoPath, 0.1)
	if er != nil {
		fmt.Printf("type  :%T\n", imag)
	}
	outFile, err := os.Create(videoCoverPath)
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, imag)
	if err != nil {
		panic(err)
	}
	err = b.Flush()
	if err != nil {
		panic(err)
	}
	//TODO异常处理
	return nil
}

func GetVideoMoment(filePath string, time float64) (i image.Image, err error) {
	//load video
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return i, err
	}

	frames := vc.Get(gocv.VideoCaptureFrameCount)
	fps := vc.Get(gocv.VideoCaptureFPS)
	duration := frames / fps

	frames = (time / duration) * frames

	// Set Video frames
	vc.Set(gocv.VideoCapturePosFrames, frames)

	img := gocv.NewMat()

	vc.Read(&img)

	imageObject, err := img.ToImage()
	if err != nil {
		return i, err
	}
	return imageObject, err
}

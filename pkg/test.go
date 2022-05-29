package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"gocv.io/x/gocv"
)

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

func Save() {
	filepath := "/home/sun/Linux/zhangyunzhe/douyin/public/bear.mp4"
	var imag image.Image
	var er error
	imag, er = GetVideoMoment(filepath, 0.1)
	if er != nil {
		fmt.Printf("type  :%T\n", imag)
	}
	fmt.Printf("type  :%T\n", imag)
	//img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	outFile, err := os.Create("/home/sun/Linux/zhangyunzhe/douyin/public/gopher2.png")
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
}
func main() {

	filedata, err := ioutil.ReadFile("/home/sun/Linux/zhangyunzhe/douyin/public/bear.mp4")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("%T\n", filedata)
	ioutil.WriteFile("/home/sun/Linux/zhangyunzhe/douyin/public/bear1.mp4", filedata, 0666)
}

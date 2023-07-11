package ffmpeg

import (
	"context"
	"fmt"
	"giligili/common/cmd"
	"strings"
)

// VideoSlice 视频切片
func VideoSlice(input, output, resolution string) string {
	split := strings.Split(resolution, "x")
	instruction := ""
	if split[0] == "1920" {
		instruction = fmt.Sprintf("/Users/mac/Desktop/project/giligili-microservice/giligili-microservice-server/common/ffmpeg/convert_1.sh %s %s", input, output)
	} else {
		instruction = fmt.Sprintf("/Users/mac/Desktop/project/giligili-microservice/giligili-microservice-server/common/ffmpeg/convert_2.sh %s %s", input, output)
	}
	result := make(chan string, 1)
	err := make(chan error, 1)
	ctc, cancelFunc := context.WithCancel(context.TODO())
	go cmd.Cmd(ctc, result, err, instruction)
	defer cancelFunc()
	for {
		select {
		case e := <-err:
			return e.Error()
		case r := <-result:
			return r
		}
	}
}

// GetVideoInformation 获取视频信息
func GetVideoInformation(input string) string {
	instruction := fmt.Sprintf("/Users/mac/Desktop/project/giligili-microservice/giligili-microservice-server/common/ffmpeg/get_video_information.sh %s", input)
	result := make(chan string, 1)
	err := make(chan error, 1)
	ctc, cancelFunc := context.WithCancel(context.TODO())
	go cmd.Cmd(ctc, result, err, instruction)
	defer cancelFunc()
	for {
		select {
		case e := <-err:
			return e.Error()
		case r := <-result:
			return r
		}
	}
}

// GetVideoDuration 获取视频时长
func GetVideoDuration(input string) string {
	instruction := fmt.Sprintf("/Users/mac/Desktop/project/giligili-microservice/giligili-microservice-server/common/ffmpeg/get_video_duration.sh %s", input)
	result := make(chan string, 1)
	err := make(chan error, 1)
	ctc, cancelFunc := context.WithCancel(context.TODO())
	go cmd.Cmd(ctc, result, err, instruction)
	defer cancelFunc()
	for {
		select {
		case e := <-err:
			return e.Error()
		case r := <-result:
			return r
		}
	}
}

// GetVideoResolution 获取视频分辨率
func GetVideoResolution(input string) string {
	instruction := fmt.Sprintf("/Users/mac/Desktop/project/giligili-microservice/giligili-microservice-server/common/ffmpeg/get_video_resolution.sh %s", input)
	result := make(chan string, 1)
	err := make(chan error, 1)
	ctc, cancelFunc := context.WithCancel(context.TODO())
	go cmd.Cmd(ctc, result, err, instruction)
	defer cancelFunc()
	for {
		select {
		case e := <-err:
			return e.Error()
		case r := <-result:
			return r
		}
	}
}

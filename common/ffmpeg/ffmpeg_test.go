package ffmpeg

import (
	"strconv"
	"strings"
	"testing"
)

func TestFfmpeg_VideoSlice(t *testing.T) {
	information := VideoSlice("/Users/mac/Downloads/英雄联盟：双城之战/mp4/英雄联盟：双城之战_01.mp4", "/Users/mac/Downloads/英雄联盟：双城之战/hls/02/", GetVideoResolution("/Users/mac/Downloads/英雄联盟：双城之战/mp4/英雄联盟：双城之战_01.mp4"))
	t.Log(information)
}

func TestFfmpeg_GetVideoInformation(t *testing.T) {
	information := GetVideoInformation("/Users/mac/Downloads/英雄联盟：双城之战/mp4/英雄联盟：双城之战_01.mp4")
	t.Log(information)
}

func TestFfmpeg_GetVideoDuration(t *testing.T) {
	duration := GetVideoDuration("/Users/mac/Downloads/英雄联盟：双城之战/mp4/英雄联盟：双城之战_01.mp4")
	t.Log(duration)
	float, err := strconv.ParseFloat(strings.Split(duration, "\n")[0], 64)
	if err != nil {
		t.Error(err)
	}
	t.Log(float)
}

func TestFfmpeg_GetVideoResolution(t *testing.T) {
	resolution := GetVideoResolution("/Users/mac/Downloads/英雄联盟：双城之战/mp4/英雄联盟：双城之战_01.mp4")
	t.Log(resolution)
}

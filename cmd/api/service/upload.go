package service

import (
	"douyin/shared/constant"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"path/filepath"
	"time"
)

// Upload Process the data acquired from request of publish action
func Upload(userID int64, c *app.RequestContext) (string, string, error) {
	playURL, saveFile, err := UploadVideo(userID, c)
	if err != nil {
		return "", "", err
	}
	coverURL, err := UploadCover(userID, saveFile)
	if err != nil {
		return playURL, "", nil
	}
	return playURL, coverURL, nil
}

func UploadVideo(userID int64, c *app.RequestContext) (string, string, error) {
	data, err := c.FormFile("data")
	if err != nil {
		return "", "", err
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), filename)
	saveFile := filepath.Join("../../public/video/", finalName)
	playURL := filepath.Join(constant.ServerAddress, "/video/", finalName)
	playURL = "http://" + playURL
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		return "", "", err
	}
	return playURL, saveFile, nil
}

func UploadCover(userID int64, playURL string) (string, error) {
	// todo code of get the video cover, temporarily use a fixed cover instead
	coverURL := "https://w.wallhaven.cc/full/d6/wallhaven-d6p1qo.jpg"
	return coverURL, nil
}

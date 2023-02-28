package service

import (
	"bytes"
	"context"
	"douyin/cmd/api/config"
	"douyin/shared/constant"
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

type UploadInfo struct {
	coverTmpPath string
	videoTmpPath string
	coverURL     string
	videoURL     string
}

type Upload struct {
	minioClient *minio.Client
	minioConfig *config.MinioConfig
}

func NewUpload(minioClient *minio.Client, minioConfig *config.MinioConfig) *Upload {
	return &Upload{
		minioClient: minioClient,
		minioConfig: minioConfig,
	}
}

func (s *Upload) UploadVideo(fh *multipart.FileHeader) (playURL, coverURL string, err error) {
	suffix := path.Ext(fh.Filename)
	sf, err := snowflake.NewSnowflake(constant.SnowFlakeDataCenterId, constant.MinioSnowFlakeWorkerId)
	if err != nil {
		return "", "", err
	}
	id := strconv.FormatInt(sf.NextVal(), 10)
	uploadPath := time.Now().Format("2006/01/02/") + id
	info := &UploadInfo{
		videoTmpPath: "./tmp/video/" + id + suffix,
		coverTmpPath: "./tmp/cover/" + id + ".png",
		videoURL:     uploadPath + suffix,
		coverURL:     uploadPath + ".png",
	}
	err = os.MkdirAll("./tmp/video", 0777)
	if err != nil && !os.IsExist(err) {
		return "", "", err
	}
	err = os.MkdirAll("./tmp/cover", 0777)
	if err != nil && !os.IsExist(err) {
		return "", "", err
	}

	videoFile, err := os.Create(info.videoTmpPath)
	if err != nil {
		return "", "", err
	}
	defer videoFile.Close()

	mpFile, err := fh.Open()
	if err != nil {
		return "", "", err
	}
	defer mpFile.Close()

	_, err = videoFile.ReadFrom(mpFile)
	if err != nil {
		return "", "", err
	}

	if err = getSnapShot(info.videoTmpPath, info.coverTmpPath); err != nil {
		hlog.Errorf("get snap shot error: %s", err.Error())
		return "", "", err
	}

	if _, err = s.minioClient.FPutObject(context.Background(), s.minioConfig.Bucket, info.coverURL, info.coverTmpPath, minio.PutObjectOptions{
		ContentType: "image/png",
	}); err != nil {
		hlog.Errorf("upload cover image error: %s", err.Error())
		return "", "", err
	}
	_ = os.Remove(info.coverTmpPath)
	if _, err = s.minioClient.FPutObject(context.Background(), s.minioConfig.Bucket, info.videoURL, info.videoTmpPath, minio.PutObjectOptions{
		ContentType: fmt.Sprintf("video/%s", suffix[1:]),
	}); err != nil {
		hlog.Errorf("upload video error: %s", err.Error())
		return "", "", err
	}
	_ = os.Remove(info.videoTmpPath)
	urlPrefix := fmt.Sprintf("http://%s:%d/%s/", constant.MinioPublicHost, s.minioConfig.Port, s.minioConfig.Bucket)
	return urlPrefix + info.videoURL, urlPrefix + info.coverURL, nil
}

func getSnapShot(videoPath, coverPath string) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	if err = imaging.Save(img, coverPath); err != nil {
		return err
	}
	return nil
}

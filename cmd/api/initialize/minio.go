package initialize

import (
	"context"
	"douyin/cmd/api/global"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func initMinio() {
	config := global.MinioConfig
	endpoint := fmt.Sprintf("%s:%d", config.Host, config.Port)
	c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UserSSL,
	})
	if err != nil {
		hlog.Fatalf("create minio client err: %s", err.Error())
	}
	exists, err := c.BucketExists(context.Background(), config.Bucket)
	if err != nil {
		hlog.Fatal(err)
	}
	if !exists {
		err = c.MakeBucket(context.Background(), config.Bucket, minio.MakeBucketOptions{Region: "cn-north-1"})
		if err != nil {
			hlog.Fatalf("make bucket err: %s", err.Error())
		}
	}
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + config.Bucket + `/*"],"Sid": ""}]}`
	err = c.SetBucketPolicy(context.Background(), config.Bucket, policy)
	if err != nil {
		hlog.Fatal("set bucket policy err:%s", err)
	}
	global.MinioClient = c
}

package cos

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/scSZn/blog/conf"
)

var ext2ContentImage = map[string]string{
	".gif":  "image/gif",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".bmp":  "image/bmp",
}

func UploadImage(ctx context.Context, filename string, reader io.Reader) (string, error) {
	baseURL := conf.GetSetting().COSSetting.BaseURL
	u, _ := url.Parse(baseURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.GetSetting().COSSetting.SecretID,
			SecretKey: conf.GetSetting().COSSetting.SecretKey,
		},
	})

	ext := strings.ToLower(path.Ext(filename))
	baseFilename := filename[:len(filename)-len(ext)]

	suffix := time.Now().Format(time.RFC3339Nano)
	name := fmt.Sprintf("%s-%s%s", baseFilename, suffix, ext)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: ext2ContentImage[ext],
		},
	}
	_, err := c.Object.Put(ctx, name, reader, opt)
	if err != nil {
		return "", errors.Wrapf(err, "upload file fail, filename: %+v", filename)
	}

	return fmt.Sprintf("%s/%s", baseURL, name), nil
}

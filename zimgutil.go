package utils

import (
	"encoding/base64"
	"github.com/levigross/grequests"
	"github.com/maczh/logs"
	"strconv"
	"strings"
)

const (
	ZIMG_HOST    = "http://47.111.191.6:4869/"
	ZIMG_GO_URL  = "http://47.111.191.6:8188/pic?id={0}&maxSize={1}"
	QINIU_HOST   = "https://img1.ququ.im/"
	QINIU_RESIZE = "?imageMogr2/size-limit/{1}k"
)

func GetZimgQiniuResizeImage(imgUrl string, size int) string {
	if strings.Contains(imgUrl, ZIMG_HOST) {
		id := strings.ReplaceAll(imgUrl, ZIMG_HOST, "")
		imgUrl = strings.ReplaceAll(strings.ReplaceAll(ZIMG_GO_URL, "{0}", id), "{1}", strconv.Itoa(size))
	}
	if strings.Contains(imgUrl, QINIU_HOST) {
		if strings.Contains(imgUrl, "?") {
			idx := strings.Index(imgUrl, "?")
			imgUrl = imgUrl[0:idx]
		}
		imgUrl = imgUrl + strings.ReplaceAll(QINIU_RESIZE, "{1}", strconv.Itoa(size))
	}
	return imgUrl
}

func DownloadImgToBase64(fileUrl string, maxSize int) (string, error) {
	url := GetZimgQiniuResizeImage(fileUrl, maxSize)
	rsp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: map[string]string{"Content-Type": "application/octet-stream"},
	})
	if err != nil {
		logs.Error("文件下载错误:{}", err.Error())
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rsp.Bytes()), nil
}

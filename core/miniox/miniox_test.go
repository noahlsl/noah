package miniox

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMinion(t *testing.T) {

	ctx := context.Background()
	cfg := Cfg{
		Address:  "127.0.0.1:9000",
		Username: "test",
		Password: "12345678",
		Bucket:   "public",
	}
	client := cfg.NewClient()

	// 打开要上传的本地文件
	file, err := os.Open("1.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	objectName := "oa/123.png"
	// 获取文件的大小和MIME类型
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Upload(ctx, file, objectName, fileInfo.Size())
	if err != nil {
		fmt.Println(err)
		return
	}

	url := client.GetUrl(ctx, objectName, Options{})
	fmt.Println(url)
}

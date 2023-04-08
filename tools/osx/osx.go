package osx

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
)

// Exists 判断所给路径文件/文件夹是否存在1
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Write(fileName string, context string) error {
	// 获取文件所在的目录
	dir := filepath.Dir(fileName)

	// 创建目录（如果不存在）
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	if Exists(fileName) {
		log.Printf("the %s is exist\n", fileName)
		return nil
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = f.Write([]byte(context))
	if err != nil {
		return errors.WithStack(err)
	}

	return f.Close()
}

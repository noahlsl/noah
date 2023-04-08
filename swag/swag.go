package swag

import (
	"fmt"
	"os"
	"strings"
)

var (
	oldStr = `"parameters": [`
	newStr = `"parameters": [
          {
            "name": "token",
            "in": "header",
            "required": true,
            "type": "string",
            "default": "{{token}}"
          },`
)

// SwaggerReplace 添加Token
func SwaggerReplace(path string) {
	fmt.Println(path)
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileStr := string(file)
	fileStr = strings.ReplaceAll(fileStr, oldStr, newStr)
	openFile, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	}

	_, err = openFile.WriteString(fileStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}

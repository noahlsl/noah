package imgx

import (
	"mime/multipart"
	"path"

	"github.com/noahlsl/noah/consts"
)

func ImgCheck(file multipart.File, header *multipart.FileHeader) error {

	//图片大小限制3M
	SizeLimit := consts.DefaultImgLimit << 20
	if header.Size > int64(SizeLimit) {
		return consts.ErrImageSizeLimit
	}

	filename := header.Filename
	fileFix := path.Ext(filename)

	// 后缀校验
	if fileFix != ".jpg" && fileFix != ".jpeg" && fileFix != ".png" {
		return consts.ErrImageSuffix
	}

	return nil
}

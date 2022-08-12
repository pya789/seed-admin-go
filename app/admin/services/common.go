package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"seed-admin/app/admin/entity"
	"seed-admin/common"
	"seed-admin/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CommonService struct{}

// 上传文件
func (*CommonService) Uploads(ctx *gin.Context) (map[string]any, error) {
	fileType := ctx.Request.FormValue("type")
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return nil, err
	}
	// 允许上传的后缀列表
	allowList := []string{"image/png", "image/jpeg", "image/jpg", "image/gif"}
	// 判断大小
	if header.Size > (1024*1024)*common.CONFIG.Int64("upload.fileSize") {
		return nil, errors.New("上传的文件大小超过限制")
	}
	// 判断格式
	if !utils.SliceIncludes(allowList, header.Header.Get("Content-Type")) {
		return nil, errors.New("禁止上传此格式")
	}
	url, id, err := local(file, header, fileType, ctx.Request.Host, ctx.Request.Header.Get("X-Forwarded-Proto"))
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"url": url,
		"id":  id,
	}, nil
}

// 本地上传
func local(file multipart.File, header *multipart.FileHeader, fileType string, host string, scheme string) (string, int, error) {
	// 读取byte
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", 0, err
	}
	// 如果uploads目录不存在则创建
	uploadsPath := fmt.Sprintf(".%v%v", common.CONFIG.String("app.staticPath"), common.CONFIG.String("upload.path"))
	if ok, _ := utils.PathExists(uploadsPath); !ok {
		if err := os.Mkdir(uploadsPath, 0777); err != nil {
			return "", 0, err
		}
	}
	// 当日的目录不存在则创建
	toDayPath := fmt.Sprintf("%v/%v", uploadsPath, time.Now().Format("2006-01-02"))
	if ok, _ := utils.PathExists(toDayPath); !ok {
		if err := os.Mkdir(toDayPath, 0777); err != nil {
			return "", 0, err
		}
	}
	// 生成UUID文件名
	fileName, err := uuid.NewV4()
	if err != nil {
		return "", 0, err
	}
	// 获取格式
	suffix := strings.Split(header.Filename, ".")[1]
	// 组装路径输出到文件
	filePath := toDayPath + "/" + fileName.String() + "." + suffix
	if err := ioutil.WriteFile(filePath, b, 0777); err != nil {
		return "", 0, err
	}
	// 判断是否有证书
	schemeStr := ""
	if scheme == "" || scheme == "http" {
		schemeStr = "http://"
	} else {
		schemeStr = "https://"
	}
	// 组装url
	url := fmt.Sprintf("%v%v%v", schemeStr, host, strings.TrimPrefix(filePath, "."))
	typeId := 1
	if ok, err := common.DB.Table("admin_uploads_type").Where("label = ?", fileType).Cols("id").Get(&typeId); !ok {
		if err != nil {
			return "", 0, err
		}
		return "", 0, errors.New("图片分类不存在")
	}
	// 入库
	upload := &entity.AdminUploads{
		Name: header.Filename,
		Url:  url,
		Type: typeId,
	}
	if _, err := common.DB.Insert(upload); err != nil {
		return "", 0, err
	}
	return upload.Url, upload.Id, nil
}

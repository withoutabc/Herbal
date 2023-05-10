package api

import (
	"archive/zip"
	"bytes"
	"github.com/gin-gonic/gin"
	"herbalBody/service"
	"herbalBody/util"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type ExcelServiceImpl struct {
	ExcelService
}

func NewExcelApi() *ExcelServiceImpl {
	return &ExcelServiceImpl{
		ExcelService: service.NewSubmissionDaoImpl(),
	}
}

type ExcelService interface {
	GenExcel(userId int) (filename string, code int, err error)
}

func (e *ExcelServiceImpl) GenExcel(c *gin.Context) {
	//先生成xlsx文件
	userId := c.Param("user_id")
	ID, err := strconv.Atoi(userId)
	if err != nil {
		util.RespInternalErr(c)
		log.Printf("strconv atoi err:%v\n", err)
		return
	}
	_, _, err = e.ExcelService.GenExcel(ID)
	if err != nil {
		if err.Error() == "用户答案未提交完全" {
			log.Printf("gen excel err:%v", err.Error())
			util.NormErr(c, 447, "用户答案未提交完全")
			return
		}
		util.RespInternalErr(c)
		log.Printf("gen excel err:%v\n", err)
		return
	}
	util.RespOK(c, "gen excel success")
	//打开文件
	//f, err := excelize.OpenFile("./excel/" + filename[1:])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 提供下载
	//c.Writer.Header().Set("Content-Type", "application/octet-stream")
	//disposition := "attachment; filename*=UTF-8''" + url.PathEscape(filename[1:])
	//c.Writer.Header().Set("Content-Disposition", disposition)
	//_ = f.Write(c.Writer)
}

func GetExcelZips(c *gin.Context) {
	// 创建最外层的zip文件
	outerZipFile, err := os.Create("excels.zip")
	if err != nil {
		log.Printf("os create err:%v", err)
		util.RespInternalErr(c)
		return
	}
	defer outerZipFile.Close()
	//创建外层
	outerZipWriter := zip.NewWriter(outerZipFile)
	// 创建一个缓冲区用于写内层zip文件
	buf := new(bytes.Buffer)
	innerZipWriter := zip.NewWriter(buf)
	// 定义内层压缩包计数器和文件计数器
	innerZipCount := 1
	innerFileCount := 0
	// 要压缩的目录路径
	dirPath := "./excel"
	// 遍历目录下的所有文件并添加到zip文件中
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果不是目录，则添加进zip文件中
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			// 在内层zip文件中创建一个新文件
			zipFile, err := innerZipWriter.Create(info.Name())
			if err != nil {
				return err
			}
			// 将Excel文件复制到内层zip文件
			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}
			innerFileCount++
			if innerFileCount == 100 {
				// 关闭内层zip写入器
				err = innerZipWriter.Close()
				if err != nil {
					return err
				}
				// 将内层zip文件写入外层zip文件中
				innerZipData := buf.Bytes()
				innerZipReader := bytes.NewReader(innerZipData)
				innerZipFileInfo, err := os.Stat("./inner.zip")
				if err != nil {
					return err
				}
				innerZipFileHeader, err := zip.FileInfoHeader(innerZipFileInfo)
				if err != nil {
					return err
				}
				innerZipFileHeader.Name = strconv.Itoa(innerZipCount) + ".zip"
				innerZipFileHeader.Method = zip.Deflate
				innerZipFile, err := outerZipWriter.Create(strconv.Itoa(innerZipCount) + ".zip")
				if err != nil {
					return err
				}
				_, err = io.Copy(innerZipFile, innerZipReader)
				if err != nil {
					return err
				}
				// 重置内层zip写入器和缓冲区
				buf = new(bytes.Buffer)
				innerZipWriter = zip.NewWriter(buf)
				innerZipCount++
				innerFileCount = 0
			}
		}
		return nil
	})
	//判断内层zip文件是否需要写入外层zip文件中
	if innerFileCount > 0 {
		// 关闭内层zip写入器
		err = innerZipWriter.Close()
		if err != nil {
			util.RespInternalErr(c)
			return
		}
		//内写外
		innerZipData := buf.Bytes()
		innerZipReader := bytes.NewReader(innerZipData)
		innerZipFileInfo, err := os.Stat("./inner.zip")
		if err != nil {
			return
		}
		innerZipFileHeader, err := zip.FileInfoHeader(innerZipFileInfo)
		if err != nil {
			return
		}
		innerZipFileHeader.Name = strconv.Itoa(innerZipCount) + ".zip"
		innerZipFileHeader.Method = zip.Deflate
		innerZipFile, err := outerZipWriter.CreateHeader(innerZipFileHeader)
		if err != nil {
			return
		}
		_, err = io.Copy(innerZipFile, innerZipReader)
		if err != nil {
			return
		}
	}

	// 关闭外层zip写入器
	err = outerZipWriter.Close()
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	// 将外层zip文件作为响应体返回给客户端
	outerZipData, err := os.ReadFile("excels.zip")
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	log.Println(len(outerZipData))
	c.Header("Content-Type", "application/binary")
	c.Header("Content-Disposition", "attachment;filename=excels.zip")
	c.Data(http.StatusOK, "application/binary", outerZipData)

}

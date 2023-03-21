/*
 * @Author: nijineko
 * @Date: 2023-03-16 15:59:52
 * @LastEditTime: 2023-03-21 12:48:17
 * @LastEditors: nijineko
 * @Description: 更新模块
 * @FilePath: \DataDownload\internal\Update\Update.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Flag"
	"BlueArchiveDataDownload/tools/CRC"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pierrec/xxHash/xxHash64"
)

/**
 * @description: 检查文件CRC，获取不一致的文件
 * @param {string} SavePath 文件保存路径
 * @param {[]Catalog.Data} CatalogData Catalog数据
 * @param {bool} xxHash OriginalFileSave模式下是否使用xxHash64计算文件名
 * @return {[]Catalog.Data} 差异数据
 */
func CheckFileCRC(SavePath string, CatalogData []Catalog.Data, xxHash bool) []Catalog.Data {
	var DifferenceDatas []Catalog.Data
	for _, Value := range CatalogData {
		// 计算文件CRC
		var FileCRC uint32
		if Flag.Data.OriginalFileSave && xxHash {
			FileName := fmt.Sprintf("%d", xxHash64.Checksum([]byte(Value.Name), 0))
			FileCRC = CRC.Checksum(path.Join(SavePath, path.Join(path.Dir(Value.Path), FileName)))
		} else {
			FileCRC = CRC.Checksum(path.Join(SavePath, Value.Path))
		}
		// 比较CRC
		if FileCRC != Value.Crc {
			DifferenceDatas = append(DifferenceDatas, Value)
		}
	}

	return DifferenceDatas
}

/**
 * @description: 复制文件到UpdateData文件夹下
 * @param {string} SavePath 文件保存路径
 * @param {[]Catalog.Data} CatalogData Catalog数据
 * @param {bool} xxHash OriginalFileSave模式下是否使用xxHash64计算文件名
 * @return {error} 错误
 */
func CopyFile(SavePath string, CatalogData []Catalog.Data, xxHash bool) error {
	for _, Value := range CatalogData {
		// 计算原始文件路径
		var FilePath string
		var FileName string
		if Flag.Data.OriginalFileSave && xxHash {
			FileName = fmt.Sprintf("%d", xxHash64.Checksum([]byte(Value.Name), 0))
			FilePath = path.Join(SavePath, path.Join(path.Dir(Value.Path), FileName))
		} else {
			FilePath = path.Join(SavePath, Value.Path)
		}

		// 计算复制文件路径
		var CopyFilePath string
		if FileName != "" {
			CopyFilePath = path.Join("./", "UpdateData", SavePath, path.Join(path.Dir(Value.Path), FileName))
		} else {
			CopyFilePath = path.Join("./", "UpdateData", SavePath, Value.Path)
		}

		// 创建文件夹
		err := os.MkdirAll(path.Dir(CopyFilePath), os.ModePerm)
		if err != nil {
			return err
		}

		// 打开源文件
		SourceFile, err := os.Open(FilePath)
		if err != nil {
			return err
		}
		defer SourceFile.Close()
		// 创建目标文件
		DestinationFile, err := os.OpenFile(CopyFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer DestinationFile.Close()

		// 复制文件
		_, err = io.Copy(DestinationFile, SourceFile)
		if err != nil {
			return err
		}
	}

	return nil
}

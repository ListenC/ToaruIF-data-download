/*
 * @Author: nijineko
 * @Date: 2023-03-16 15:59:52
 * @LastEditTime: 2023-03-26 03:57:55
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
	"path/filepath"

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
			FileCRC = CRC.Checksum(path.Join(SavePath, FileName))
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
			FilePath = path.Join(SavePath, FileName)
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

/**
 * @description: 删除多余文件
 * @param {string} SavePath	文件保存路径
 * @param {string} CatalogPath 本地Catalog文件路径
 * @param {[]Catalog.Data} CatalogData Catalog数据
 * @param {bool} xxHash OriginalFileSave模式下是否使用xxHash64计算文件名
 * @return {[]string} 删除文件路径
 */
func CleanFile(SavePath string, CatalogPath string, CatalogData []Catalog.Data, xxHash bool) ([]string, error) {
	// 遍历CatalogData，转换为map
	CatalogDataMap := make(map[string]struct{})
	for _, Value := range CatalogData {
		// 计算文件路径
		var FilePath string
		if Flag.Data.OriginalFileSave && xxHash {
			FileName := fmt.Sprintf("%d", xxHash64.Checksum([]byte(Value.Name), 0))
			FilePath = path.Join(SavePath, FileName)
		} else {
			FilePath = path.Join(SavePath, Value.Path)
		}

		CatalogDataMap[FilePath] = struct{}{}
	}

	// 遍历本地文件
	FilePaths, err := getFilePaths(SavePath)
	if err != nil {
		return nil, err
	}

	var DeleteFilePaths []string
	for _, Value := range FilePaths {
		// 转换路径为正斜杠
		FilePath := filepath.ToSlash(Value)

		// 跳过Catalog文件
		if FilePath == CatalogPath {
			continue
		}

		// 判断文件是否在CatalogDataMap中
		if _, Find := CatalogDataMap[FilePath]; !Find {
			DeleteFilePaths = append(DeleteFilePaths, FilePath)

			// 删除文件
			err := os.Remove(FilePath)
			if err != nil {
				return nil, err
			}
		}
	}

	return DeleteFilePaths, nil
}

/**
 * @description: 遍历出所有文件夹内文件路径
 * @param {string} DirPth
 * @return {[]string} 文件路径
 * @return {error} 错误
 */
func getFilePaths(DirPth string) ([]string, error) {
	DirPth = filepath.Clean(DirPth)
	var dirs []string
	dir, err := os.ReadDir(DirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Files []string
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
			getFilePaths(DirPth + PthSep + fi.Name())
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := getFilePaths(table)
		for _, temp1 := range temp {
			Files = append(Files, filepath.Clean(temp1))
		}
	}

	return Files, nil
}

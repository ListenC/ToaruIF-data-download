/*
 * @Author: nijineko
 * @Date: 2023-03-16 15:59:52
 * @LastEditTime: 2023-03-19 03:01:24
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

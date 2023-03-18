/*
 * @Author: nijineko
 * @Date: 2023-03-03 22:58:53
 * @LastEditTime: 2023-03-16 16:36:29
 * @LastEditors: nijineko
 * @Description: Catalog
 * @FilePath: \DataDownload\internal\Catalog\Catalog.go
 */
package Catalog

import (
	"BlueArchiveDataDownload/internal/Flag"
	"os"
	"path"
)

var TableBundlesCataLogPath = "/TableBundles/TableCatalog.json"        // TableBundles的CatLog文件路径
var AndroidAssetBundlsCataLogPath = "/Android/bundleDownloadInfo.json" // AssetBundles的CatLog文件路径(Android)
var MediaResourcesCataLogPath = "/MediaResources/MediaCatalog.json"    // MediaResources的CatLog文件路径

// 标准结构体
type Data struct {
	Name string `json:"Name"`
	Path string `json:"Path"`
	Crc  uint32    `json:"Crc"`
}

/**
 * @description: 保存JSON文件
 * @param {[]byte} JsonBody Json数据
 * @param {string} SavePath 保存路径
 * @return {error} 错误信息
 */
func SaveJson(JsonBody []byte, SavePath string) error {
	// 如果不保存CatLog文件则直接跳过
	if !Flag.Data.SaveCatalog {
		return nil
	}

	err := CreateFolder(path.Dir(SavePath))
	if err != nil {
		return err
	}

	return os.WriteFile(SavePath, JsonBody, 0666)
}

/**
 * @description: 创建文件夹
 * @param {string} Path 文件夹路径
 * @return {error} 错误信息
 */
func CreateFolder(Path string) error {
	_, err := os.Stat(Path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

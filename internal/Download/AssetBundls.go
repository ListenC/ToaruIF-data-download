/*
 * @Author: nijineko
 * @Date: 2023-03-04 00:52:56
 * @LastEditTime: 2023-03-04 01:21:31
 * @LastEditors: nijineko
 * @Description: 下载AssetBundls文件
 * @FilePath: \DataDownload\internal\Download\AssetBundls.go
 */
package Download

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"fmt"
)

/**
 * @description: 下载AssetBundls文件
 * @param {string} AddressablesCatalogURLRoot
 * @param {[]Catalog.Data} CatalogData
 * @param {string} SavePath
 * @return {error} 错误信息
 */
func AssetBundls(AddressablesCatalogURLRoot string, CatalogData []Catalog.Data, SavePath string) error {
	var AndroidAssetBundlsURLPath = "/Android/"

	fmt.Println("开始下载AssetBundls文件")
	return Resource(CatalogData, AddressablesCatalogURLRoot+AndroidAssetBundlsURLPath, SavePath, false)
}

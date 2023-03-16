/*
 * @Author: nijineko
 * @Date: 2023-03-04 00:56:50
 * @LastEditTime: 2023-03-04 01:20:49
 * @LastEditors: nijineko
 * @Description: 下载MediaResources文件
 * @FilePath: \DataDownload\internal\Download\MediaResources.go
 */
package Download

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"fmt"
)

/**
 * @description: 下载MediaResources文件
 * @param {string} AddressablesCatalogURLRoot
 * @param {[]Catalog.Data} CatalogData
 * @param {string} SavePath
 * @return {error} 错误信息
 */
func MediaResources(AddressablesCatalogURLRoot string, CatalogData []Catalog.Data, SavePath string) error {
	fmt.Println("开始下载MediaResources文件")
	return Resource(CatalogData, AddressablesCatalogURLRoot+MediaResourcesURLPath, SavePath, true)
}

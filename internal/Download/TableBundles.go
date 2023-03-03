/*
 * @Author: nijineko
 * @Date: 2023-03-04 00:36:47
 * @LastEditTime: 2023-03-04 00:53:14
 * @LastEditors: nijineko
 * @Description: 下载TableBundles文件
 * @FilePath: \DataDownload\internal\Download\TableBundles.go
 */
package Download

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"fmt"
)

/**
 * @description: 下载TableBundles文件
 * @param {string} AddressablesCatalogURLRoot 资源地址
 * @param {[]Catalog.Data} CatalogData  Catalog数据
 * @param {string} SavePath 保存路径
 * @return {error} 错误信息
 */
func TableBundles(AddressablesCatalogURLRoot string, CatalogData []Catalog.Data, SavePath string) error {
	var TableBundlesURLPath = "/TableBundles/"

	fmt.Println("开始下载TableBundles文件")
	return Resource(CatalogData, AddressablesCatalogURLRoot+TableBundlesURLPath, SavePath, true)
}

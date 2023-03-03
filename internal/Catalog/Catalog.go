/*
 * @Author: nijineko
 * @Date: 2023-03-03 22:58:53
 * @LastEditTime: 2023-03-03 23:19:58
 * @LastEditors: nijineko
 * @Description: Catalog
 * @FilePath: \DataDownload\internal\Catalog\Catalog.go
 */
package Catalog

var TableBundlesCataLogPath = "/TableBundles/TableCatalog.json"        // TableBundles的CatLog文件路径
var AndroidAssetBundlsCataLogPath = "/Android/bundleDownloadInfo.json" // AssetBundles的CatLog文件路径(Android)
var MediaResourcesCataLogPath = "/MediaResources/MediaCatalog.json"    // MediaResources的CatLog文件路径

// 标准结构体
type Data struct {
	Name string `json:"Name"`
	Path string `json:"Path"`
	Crc  int    `json:"Crc"`
}

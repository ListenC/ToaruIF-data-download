/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:04:06
 * @LastEditTime: 2023-03-04 16:44:26
 * @LastEditors: nijineko
 * @Description: 参数解析
 * @FilePath: \DataDownload\internal\Flag\Flag.go
 */
package Flag

import "flag"

type Flag struct {
	Version          string // 数据包版本
	OriginalFileSave bool   // 是否以原始文件的名字和路径保存
	MaxPool          int    // 最大并发数
	Filter           string // 字符串过滤器，只下载包含该字符串的文件
	SaveCatalog      bool   // 是否保存Catalog文件
	Update           bool   // 以更新模式启动程序
	AssetBundls      bool   // 下载/更新AssetBundls文件
	TableBundles     bool   // 下载/更新TableBundles文件
	MediaResources   bool   // 下载/更新MediaResources文件
}

var Data Flag

/**
 * @description: 初始化参数
 * @return {error} 错误
 */
func Init() error {
	// 参数解析
	Version := flag.String("data_version", "", "指定数据包版本")
	OriginalFileSave := flag.Bool("original_file_save", false, "是否以原始文件的名字和路径保存")
	MaxPool := flag.Int("max_pool", 10, "最大并发数")
	Filter := flag.String("filter", "", "字符串过滤器，只下载包含该字符串的文件")
	SaveCatalog := flag.Bool("save_catalog", true, "是否保存Catalog文件")
	Update := flag.Bool("update", false, "以更新模式启动程序")
	AssetBundls := flag.Bool("asset_bundls", false, "下载/更新AssetBundls文件")
	TableBundles := flag.Bool("table_bundles", false, "下载/更新TableBundles文件")
	MediaResources := flag.Bool("media_resources", false, "下载/更新MediaResources文件")
	flag.Parse()

	// 参数写入变量
	Data.Version = *Version
	Data.OriginalFileSave = *OriginalFileSave
	Data.MaxPool = *MaxPool
	Data.Filter = *Filter
	Data.SaveCatalog = *SaveCatalog
	Data.Update = *Update
	if *AssetBundls || *TableBundles || *MediaResources {
		Data.AssetBundls = *AssetBundls
		Data.TableBundles = *TableBundles
		Data.MediaResources = *MediaResources
	} else {
		Data.AssetBundls = true
		Data.TableBundles = true
		Data.MediaResources = true
	}

	return nil
}

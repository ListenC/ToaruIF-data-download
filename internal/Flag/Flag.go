/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:04:06
 * @LastEditTime: 2023-03-03 23:04:13
 * @LastEditors: nijineko
 * @Description: 参数解析
 * @FilePath: \DataDownload\internal\Flag\Flag.go
 */
package Flag

import "flag"

type Flag struct {
	Version          string // 数据包版本
	OriginalFileSave bool   // 是否以原始文件的名字和路径保存
	AssetBundls      bool   // 下载AssetBundls文件
	TableBundles     bool   // 下载TableBundles文件
	MediaResources   bool   // 下载MediaResources文件
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
	AssetBundls := flag.Bool("asset_bundls", false, "下载AssetBundls文件")
	TableBundles := flag.Bool("table_bundles", false, "下载TableBundles文件")
	MediaResources := flag.Bool("media_resources", false, "下载MediaResources文件")
	flag.Parse()

	// 参数写入变量
	Data.Version = *Version
	Data.OriginalFileSave = *OriginalFileSave
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

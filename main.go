/*
 * @Author: NyanCatda
 * @Date: 2023-03-03 22:42:20
 * @LastEditTime: 2023-03-16 17:13:21
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \DataDownload\main.go
 */
package main

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"BlueArchiveDataDownload/internal/Flag"
	"BlueArchiveDataDownload/internal/HTTP"
	"BlueArchiveDataDownload/internal/MateData"
	"BlueArchiveDataDownload/internal/Update"
	"fmt"
	"os"
	"path"
	"time"
)

var (
	AssetBundlsSavePath    string // AssetBundls资源保存路径
	TableBundlesSavePath   string // TableBundles资源保存路径
	MediaResourcesSavePath string // MediaResources资源保存路径
)

func main() {
	// 启动程序
	boot()

	// 读取元数据
	Mate, err := MateData.Get(Flag.Data.Version)
	if err != nil {
		fmt.Println(err)
		fmt.Println("元数据获取失败，可能是版本号不正确或者网络问题")
		os.Exit(1)
	}

	AddressablesCatalogUrlRoot := Mate.ConnectionGroups[0].OverrideConnectionGroups[len(Mate.ConnectionGroups[0].OverrideConnectionGroups)-1].AddressablesCatalogURLRoot

	// 判断资源服务器是否可用
	for {
		_, Response, err := HTTP.Get(AddressablesCatalogUrlRoot + Catalog.TableBundlesCataLogPath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("资源服务器信息获取失败")
			os.Exit(1)
		}
		if Response.StatusCode != 200 {
			fmt.Println("资源服务器不可用，将会每隔5秒尝试连接一次，直到成功为止")
			time.Sleep(5 * time.Second)
			continue
		} else {
			break
		}
	}

	// 判断是否以更新模式启动
	if Flag.Data.Update {
		bootUpdate(AddressablesCatalogUrlRoot)
	}

	// 默认以下载模式启动
	bootDownload(AddressablesCatalogUrlRoot)
}

func boot() {
	// 初始化参数
	err := Flag.Init()
	if err != nil {
		panic(err)
	}

	// 计算文件保存路径
	if Flag.Data.OriginalFileSave {
		AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "files", "AssetBundls")
	} else {
		AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "AssetBundls")
	}
	if Flag.Data.OriginalFileSave {
		TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "TableBundles")
	} else {
		TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "TableBundles")
	}
	if Flag.Data.OriginalFileSave {
		MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "MediaPatch")
	} else {
		MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "MediaResources")
	}
}

/**
 * @description: 启动下载模式
 * @param {string} AddressablesCatalogUrlRoot 资源服务器地址
 * @return {*}
 */
func bootDownload(AddressablesCatalogUrlRoot string) {
	// 下载AssetBundls资源
	if Flag.Data.AssetBundls {
		// 获取Catalog文件
		AssetBundlsCataLog, err := Catalog.GetAssetBundls(AddressablesCatalogUrlRoot, AssetBundlsSavePath)
		if err != nil {
			panic(err)
		}

		// 下载资源
		err = Download.AssetBundls(AddressablesCatalogUrlRoot, AssetBundlsCataLog, AssetBundlsSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 下载TableBundles资源
	if Flag.Data.TableBundles {
		// 获取Catalog文件
		TableBundlesCataLog, err := Catalog.GetTableBundles(AddressablesCatalogUrlRoot, TableBundlesSavePath)
		if err != nil {
			panic(err)
		}

		// 下载资源
		err = Download.TableBundles(AddressablesCatalogUrlRoot, TableBundlesCataLog, TableBundlesSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 下载MediaResources资源
	if Flag.Data.MediaResources {
		// 获取Catalog文件
		MediaResourcesCataLog, err := Catalog.GetMediaResources(AddressablesCatalogUrlRoot, MediaResourcesSavePath)
		if err != nil {
			panic(err)
		}

		// 下载资源
		err = Download.MediaResources(AddressablesCatalogUrlRoot, MediaResourcesCataLog, MediaResourcesSavePath)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("下载完成")
	os.Exit(0)
}

/**
 * @description: 启动更新模式
 * @param {string} AddressablesCatalogUrlRoot 资源服务器地址
 * @return {*}
 */
func bootUpdate(AddressablesCatalogUrlRoot string) {
	// 更新AssetBundls资源
	if Flag.Data.AssetBundls {
		err := Update.AssetBundls(AddressablesCatalogUrlRoot, AssetBundlsSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 更新TableBundles资源
	if Flag.Data.TableBundles {
		err := Update.TableBundles(AddressablesCatalogUrlRoot, TableBundlesSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 更新MediaResources资源
	if Flag.Data.MediaResources {
		err := Update.MediaResources(AddressablesCatalogUrlRoot, MediaResourcesSavePath)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("更新完成")
	os.Exit(0)
}

/*
 * @Author: NyanCatda
 * @Date: 2023-03-03 22:42:20
 * @LastEditTime: 2023-03-08 18:00:37
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
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	// 初始化参数
	err := Flag.Init()
	if err != nil {
		panic(err)
	}

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

	// 下载AssetBundls资源
	if Flag.Data.AssetBundls {
		var AssetBundlsSavePath string
		if Flag.Data.OriginalFileSave {
			AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "files", "AssetBundls")
		} else {
			AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "AssetBundls")
		}

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
		var TableBundlesSavePath string
		if Flag.Data.OriginalFileSave {
			TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "TableBundles")
		} else {
			TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "TableBundles")
		}

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
		var MediaResourcesSavePath string
		if Flag.Data.OriginalFileSave {
			MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "MediaPatch")
		} else {
			MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "MediaResources")
		}

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

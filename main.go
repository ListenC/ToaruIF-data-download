/*
 * @Author: NyanCatda
 * @Date: 2023-03-03 22:42:20
 * @LastEditTime: 2023-03-04 01:27:06
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \DataDownload\main.go
 */
package main

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"BlueArchiveDataDownload/internal/Flag"
	"BlueArchiveDataDownload/internal/MateData"
	"fmt"
	"os"
	"path"
)

func init() {
	// 初始化参数
	err := Flag.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	// 读取元数据
	Mate, err := MateData.Get(Flag.Data.Version)
	if err != nil {
		fmt.Println(err)
		fmt.Println("元数据获取失败，可能是版本号不正确或者网络问题")
		os.Exit(1)
	}

	AddressablesCatalogUrlRoot := Mate.ConnectionGroups[0].OverrideConnectionGroups[len(Mate.ConnectionGroups[0].OverrideConnectionGroups)-1].AddressablesCatalogURLRoot

	// 下载AssetBundls资源
	if Flag.Data.AssetBundls {
		AssetBundlsCataLog, err := Catalog.GetAssetBundls(AddressablesCatalogUrlRoot)
		if err != nil {
			panic(err)
		}
		var AssetBundlsSavePath string
		if Flag.Data.OriginalFileSave {
			AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "files", "AssetBundls")
		} else {
			AssetBundlsSavePath = path.Join("com.YostarJP.BlueArchive", "AssetBundls")
		}
		err = Download.AssetBundls(AddressablesCatalogUrlRoot, AssetBundlsCataLog, AssetBundlsSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 下载TableBundles资源
	if Flag.Data.TableBundles {
		TableBundlesCataLog, err := Catalog.GetTableBundles(AddressablesCatalogUrlRoot)
		if err != nil {
			panic(err)
		}
		var TableBundlesSavePath string
		if Flag.Data.OriginalFileSave {
			TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "TableBundles")
		} else {
			TableBundlesSavePath = path.Join("com.YostarJP.BlueArchive", "TableBundles")
		}
		err = Download.TableBundles(AddressablesCatalogUrlRoot, TableBundlesCataLog, TableBundlesSavePath)
		if err != nil {
			panic(err)
		}
	}

	// 下载MediaResources资源
	if Flag.Data.MediaResources {
		MediaResourcesCataLog, err := Catalog.GetMediaResources(AddressablesCatalogUrlRoot)
		if err != nil {
			panic(err)
		}
		var MediaResourcesSavePath string
		if Flag.Data.OriginalFileSave {
			MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "files", "MediaPatch")
		} else {
			MediaResourcesSavePath = path.Join("com.YostarJP.BlueArchive", "MediaResources")
		}
		err = Download.MediaResources(AddressablesCatalogUrlRoot, MediaResourcesCataLog, MediaResourcesSavePath)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("下载完成")
	os.Exit(0)
}

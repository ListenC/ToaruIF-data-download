/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:52:45
 * @LastEditTime: 2023-03-19 03:00:55
 * @LastEditors: nijineko
 * @Description: 更新AssetBundls
 * @FilePath: \DataDownload\internal\Update\AssetBundls.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"fmt"
)

/**
 * @description: 更新AssetBundls的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func AssetBundls(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取远程AssetBundls Catalog
	RemoteCatalogData, err := Catalog.GetAssetBundls(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 对比本地文件与远程Catalog的CRC
	fmt.Println("开始检查AssetBundls文件是否需要更新")
	NeedUpdateFiles := CheckFileCRC(SavePath, RemoteCatalogData, false)

	if len(NeedUpdateFiles) == 0 {
		fmt.Println("AssetBundls文件无需更新")
		return nil
	} else {
		fmt.Printf("共有%d个AssetBundls文件需要更新，开始下载\n", len(NeedUpdateFiles))
	}

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.AndroidAssetBundlsURLPath, SavePath, false)
	if err != nil {
		return err
	}

	return nil
}

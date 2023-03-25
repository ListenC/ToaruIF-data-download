/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:52:45
 * @LastEditTime: 2023-03-26 03:55:54
 * @LastEditors: nijineko
 * @Description: 更新AssetBundls
 * @FilePath: \DataDownload\internal\Update\AssetBundls.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"BlueArchiveDataDownload/internal/Flag"
	"fmt"
	"path"
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
	} else {
		fmt.Printf("共有%d个AssetBundls文件需要更新，开始下载\n", len(NeedUpdateFiles))
	}

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.AndroidAssetBundlsURLPath, SavePath, false)
	if err != nil {
		return err
	}

	// 复制更新的文件
	if Flag.Data.UpdateCopy {
		err = CopyFile(SavePath, NeedUpdateFiles, false)
		if err != nil {
			return err
		}
	}

	// 清理过期的文件
	if Flag.Data.UpdateClean {
		fmt.Println("开始清理过期的AssetBundls文件")
		DeleteFiles, err := CleanFile(SavePath, path.Join(SavePath, path.Base(Catalog.AndroidAssetBundlsCataLogPath)), RemoteCatalogData, false)
		if err != nil {
			return err
		}
		if len(DeleteFiles) != 0 {
			fmt.Printf("已删除%d个过期的AssetBundls文件\n", len(DeleteFiles))
		} else {
			fmt.Println("没有过期的AssetBundls文件")
		}
	}

	return nil
}

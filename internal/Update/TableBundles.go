/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:05:15
 * @LastEditTime: 2023-03-21 12:52:45
 * @LastEditors: nijineko
 * @Description: 更新TableBundles
 * @FilePath: \DataDownload\internal\Update\TableBundles.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"BlueArchiveDataDownload/internal/Flag"
	"fmt"
)

/**
 * @description: 更新TableBundles的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func TableBundles(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取远程TableBundles Catalog
	RemoteCatalogData, err := Catalog.GetTableBundles(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 对比本地文件与远程Catalog的CRC
	fmt.Println("开始检查TableBundles文件是否需要更新")
	NeedUpdateFiles := CheckFileCRC(SavePath, RemoteCatalogData, true)

	if len(NeedUpdateFiles) == 0 {
		fmt.Println("TableBundles文件无需更新")
		return nil
	} else {
		fmt.Printf("共有%d个TableBundles文件需要更新，开始下载\n", len(NeedUpdateFiles))
	}

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.TableBundlesURLPath, SavePath, true)
	if err != nil {
		return err
	}

	// 复制更新的文件
	if Flag.Data.UpdateCopy {
		err = CopyFile(SavePath, NeedUpdateFiles, true)
		if err != nil {
			return err
		}
	}

	return nil
}

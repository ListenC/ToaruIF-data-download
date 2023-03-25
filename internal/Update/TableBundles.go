/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:05:15
 * @LastEditTime: 2023-03-26 04:06:02
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
	"path"
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

	// 清理过期的文件
	if Flag.Data.UpdateClean {
		fmt.Println("开始清理过期的TableBundles文件")
		DeleteFiles, err := CleanFile(SavePath, path.Join(SavePath, path.Base(Catalog.TableBundlesCataLogPath)), RemoteCatalogData, true)
		if err != nil {
			return err
		}
		if len(DeleteFiles) != 0 {
			fmt.Printf("已删除%d个过期的TableBundles文件\n", len(DeleteFiles))
		} else {
			fmt.Println("没有过期的TableBundles文件")
		}
	}

	return nil
}

/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:58:21
 * @LastEditTime: 2023-03-26 04:07:05
 * @LastEditors: nijineko
 * @Description: 更新MediaResources
 * @FilePath: \DataDownload\internal\Update\MediaResources.go
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
 * @description: 更新MediaResources的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func MediaResources(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取远程MediaResources Catalog
	RemoteCatalogData, err := Catalog.GetMediaResources(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 对比本地文件与远程Catalog的CRC
	fmt.Println("开始检查MediaResources文件是否需要更新")
	NeedUpdateFiles := CheckFileCRC(SavePath, RemoteCatalogData, true)

	if len(NeedUpdateFiles) == 0 {
		fmt.Println("MediaResources文件无需更新")
	} else {
		fmt.Printf("共有%d个MediaResources文件需要更新，开始下载\n", len(NeedUpdateFiles))
	}

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.MediaResourcesURLPath, SavePath, true)
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
		fmt.Println("开始清理过期的MediaResources文件")
		DeleteFiles, err := CleanFile(SavePath, path.Join(SavePath, path.Base(Catalog.MediaResourcesCataLogPath)), RemoteCatalogData, true)
		if err != nil {
			return err
		}
		if len(DeleteFiles) != 0 {
			fmt.Printf("已删除%d个过期的MediaResources文件\n", len(DeleteFiles))
		} else {
			fmt.Println("没有过期的MediaResources文件")
		}
	}

	return nil
}

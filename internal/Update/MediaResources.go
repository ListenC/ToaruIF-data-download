/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:58:21
 * @LastEditTime: 2023-03-16 17:19:33
 * @LastEditors: nijineko
 * @Description: 更新MediaResources
 * @FilePath: \DataDownload\internal\Update\MediaResources.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Download"
	"fmt"
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
		return nil
	} else {
		fmt.Printf("共有%d个MediaResources文件需要更新，开始下载\n", len(NeedUpdateFiles))
	}

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.MediaResourcesURLPath, SavePath, true)
	if err != nil {
		return err
	}

	return nil
}

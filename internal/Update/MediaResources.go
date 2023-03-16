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
	"encoding/json"
	"os"
	"path"
	"path/filepath"
)

/**
 * @description: 更新MediaResources的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func MediaResources(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取本地MediaResources Catalog
	LocalCatalogFile, err := os.ReadFile(path.Join(SavePath, filepath.Base(Catalog.MediaResourcesCataLogPath)))
	if err != nil {
		return err
	}
	var LocalCatalog Catalog.MediaResourcesOrigin
	err = json.Unmarshal(LocalCatalogFile, &LocalCatalog)
	if err != nil {
		return err
	}
	// 转换为标准结构体
	LocalCatalogData := LocalCatalog.ToData()

	// 获取远程MediaResources Catalog
	RemoteCatalogData, err := Catalog.GetMediaResources(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 比较两个Catalog，获取需要更新的文件
	NeedUpdateFiles := CompareDataCrc(LocalCatalogData, RemoteCatalogData)

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.MediaResourcesURLPath, SavePath, true)
	if err != nil {
		return err
	}

	return nil
}

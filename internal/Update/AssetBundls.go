/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:52:45
 * @LastEditTime: 2023-03-16 17:19:45
 * @LastEditors: nijineko
 * @Description: 更新AssetBundls
 * @FilePath: \DataDownload\internal\Update\AssetBundls.go
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
 * @description: 更新AssetBundls的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func AssetBundls(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取本地AssetBundls Catalog
	LocalCatalogFile, err := os.ReadFile(path.Join(SavePath, filepath.Base(Catalog.AndroidAssetBundlsCataLogPath)))
	if err != nil {
		return err
	}
	var LocalCatalog Catalog.AssetBundlesOrigin
	err = json.Unmarshal(LocalCatalogFile, &LocalCatalog)
	if err != nil {
		return err
	}
	// 转换为标准结构体
	LocalCatalogData := LocalCatalog.ToData()

	// 获取远程AssetBundls Catalog
	RemoteCatalogData, err := Catalog.GetAssetBundls(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 比较两个Catalog，获取需要更新的文件
	NeedUpdateFiles := CompareDataCrc(LocalCatalogData, RemoteCatalogData)

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.AndroidAssetBundlsURLPath, SavePath, true)
	if err != nil {
		return err
	}

	return nil
}

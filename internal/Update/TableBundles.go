/*
 * @Author: nijineko
 * @Date: 2023-03-16 16:05:15
 * @LastEditTime: 2023-03-16 16:45:59
 * @LastEditors: nijineko
 * @Description: 更新TableBundles
 * @FilePath: \DataDownload\internal\Update\TableBundles.go
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
 * @description: 更新TableBundles的文件
 * @param {string} AddressablesCatalogUrlRoot
 * @param {string} SavePath
 * @return {error} error
 */
func TableBundles(AddressablesCatalogUrlRoot string, SavePath string) error {
	// 获取本地TableBundles Catalog
	LocalCatalogFile, err := os.ReadFile(path.Join(SavePath, filepath.Base(Catalog.TableBundlesCataLogPath)))
	if err != nil {
		return err
	}
	var LocalCatalog Catalog.TableBundlesOrigin
	err = json.Unmarshal(LocalCatalogFile, &LocalCatalog)
	if err != nil {
		return err
	}
	// 转换为标准结构体
	LocalCatalogData := LocalCatalog.ToData()

	// 获取远程TableBundles Catalog
	RemoteCatalogData, err := Catalog.GetTableBundles(AddressablesCatalogUrlRoot, SavePath)
	if err != nil {
		return err
	}

	// 比较两个Catalog，获取需要更新的文件
	NeedUpdateFiles := CompareDataCrc(LocalCatalogData, RemoteCatalogData)

	// 下载需要更新的文件
	err = Download.Resource(NeedUpdateFiles, AddressablesCatalogUrlRoot+Download.TableBundlesURLPath, SavePath, true)
	if err != nil {
		return err
	}

	return nil
}

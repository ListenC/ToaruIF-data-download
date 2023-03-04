/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:28:57
 * @LastEditTime: 2023-03-04 13:21:40
 * @LastEditors: nijineko
 * @Description: 读取AssetBundls的CataLog文件到标准结构体
 * @FilePath: \DataDownload\internal\Catalog\AssetBundls.go
 */
package Catalog

import (
	"BlueArchiveDataDownload/internal/HTTP"
	"encoding/json"
	"path"
)

type AssetBundlesOrigin struct {
	BundleFiles []struct {
		Name      string `json:"Name"`      // 文件名字
		Size      int    `json:"Size"`      // 文件大小
		IsInbuild bool   `json:"isInbuild"` // 是否为内置文件
		Crc       int    `json:"Crc"`       // 文件CRC
	} `json:"BundleFiles"`
}

/**
 * @description: 获取AssetBundls的CataLog数据
 * @param {string} AddressablesCatalogUrlRoot 资源地址
 * @param {string} SavePath json文件保存路径(为空则不保存)
 * @return {[]Data} CatLog数据
 * @return {error} 错误信息
 */
func GetAssetBundls(AddressablesCatalogUrlRoot string, SavePath string) ([]Data, error) {
	Body, _, err := HTTP.Get(AddressablesCatalogUrlRoot + AndroidAssetBundlsCataLogPath)
	if err != nil {
		return nil, err
	}

	if SavePath != "" {
		err := SaveJson(Body, path.Join(SavePath, path.Base(AndroidAssetBundlsCataLogPath)))
		if err != nil {
			return nil, err
		}
	}

	// 使用原始结构体解析JSON
	var AssetBundls AssetBundlesOrigin
	err = json.Unmarshal(Body, &AssetBundls)
	if err != nil {
		return nil, err
	}

	// 转换为标准结构体
	var AssetBundlsData []Data
	for _, Value := range AssetBundls.BundleFiles {
		AssetBundlsData = append(AssetBundlsData, Data{
			Name: Value.Name,
			Path: Value.Name,
			Crc:  Value.Crc,
		})
	}

	return AssetBundlsData, nil
}

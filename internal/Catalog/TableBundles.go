/*
 * @Author: nijineko
 * @Date: 2023-03-03 22:56:45
 * @LastEditTime: 2023-03-03 23:25:09
 * @LastEditors: nijineko
 * @Description: 读取TableBundles的CataLog文件到标准结构体
 * @FilePath: \DataDownload\internal\Catalog\TableBundles.go
 */
package Catalog

import (
	"BlueArchiveDataDownload/internal/HTTP"
	"encoding/json"
)

type TableBundlesOrigin struct {
	Table map[string]struct {
		Name      string   `json:"Name"`      // 文件名字
		Size      int      `json:"Size"`      // 文件大小
		Crc       int      `json:"Crc"`       // 文件CRC
		IsInbuild bool     `json:"isInbuild"` // 是否为内置文件
		IsChanged bool     `json:"isChanged"` // 是否被修改
		Includes  []string `json:"Includes"`  // 包含的文件
	} `json:"Table"`
	BundleMap any `json:"BundleMap"`
}

/**
 * @description: 获取TableBundles的CataLog数据
 * @param {string} AddressablesCatalogUrlRoot 资源地址
 * @return {[]Data} CatLog数据
 * @return {error} 错误信息
 */
func GetTableBundles(AddressablesCatalogUrlRoot string) ([]Data, error) {
	Body, _, err := HTTP.Get(AddressablesCatalogUrlRoot + TableBundlesCataLogPath)
	if err != nil {
		return nil, err
	}

	// 使用原始结构体解析JSON
	var TableBundles TableBundlesOrigin
	err = json.Unmarshal(Body, &TableBundles)
	if err != nil {
		return nil, err
	}

	// 转换为标准结构体
	var TableBundlesData []Data
	for _, Value := range TableBundles.Table {
		TableBundlesData = append(TableBundlesData, Data{
			Name: Value.Name,
			Path: Value.Name,
			Crc:  Value.Crc,
		})
	}

	return TableBundlesData, nil
}

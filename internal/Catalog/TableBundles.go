/*
 * @Author: nijineko
 * @Date: 2023-03-03 22:56:45
 * @LastEditTime: 2023-03-04 13:13:48
 * @LastEditors: nijineko
 * @Description: 读取TableBundles的CataLog文件到标准结构体
 * @FilePath: \DataDownload\internal\Catalog\TableBundles.go
 */
package Catalog

import (
	"BlueArchiveDataDownload/internal/Flag"
	"BlueArchiveDataDownload/internal/HTTP"
	"encoding/json"
	"path"
)

type TableBundlesOrigin struct {
	Table map[string]struct {
		Name      string   `json:"Name"`      // 文件名字
		Size      int      `json:"Size"`      // 文件大小
		Crc       uint32   `json:"Crc"`       // 文件CRC
		IsInbuild bool     `json:"isInbuild"` // 是否为内置文件
		IsChanged bool     `json:"isChanged"` // 是否被修改
		Includes  []string `json:"Includes"`  // 包含的文件
	} `json:"Table"`
	BundleMap any `json:"BundleMap"`
}

/**
 * @description: 获取TableBundles的CataLog数据
 * @param {string} AddressablesCatalogUrlRoot 资源地址
 * @param {string} SavePath json文件保存路径(为空则不保存)
 * @return {[]Data} CatLog数据
 * @return {error} 错误信息
 */
func GetTableBundles(AddressablesCatalogUrlRoot string, SavePath string) ([]Data, error) {
	Body, _, err := HTTP.Get(AddressablesCatalogUrlRoot + TableBundlesCataLogPath)
	if err != nil {
		return nil, err
	}

	if SavePath != "" {
		err := SaveJson(Body, path.Join(SavePath, path.Base(TableBundlesCataLogPath)))
		if err != nil {
			return nil, err
		}
	}

	// 使用原始结构体解析JSON
	var TableBundles TableBundlesOrigin
	err = json.Unmarshal(Body, &TableBundles)
	if err != nil {
		return nil, err
	}

	// 转换为标准结构体
	TableBundlesData := TableBundles.ToData()

	return TableBundlesData, nil
}

/**
 * @description: 将TableBundlesOrigin转换为标准Catalog结构体
 * @return {[]Data} CatLog数据
 */
func (Origin TableBundlesOrigin) ToData() []Data {
	var CatalogDatas []Data
	for _, Value := range Origin.Table {
		// 如果忽略内置文件并且当前文件为内置文件则跳过
		if !Flag.Data.IgnoreInbuild && Value.IsInbuild {
			continue
		}

		CatalogDatas = append(CatalogDatas, Data{
			Name: Value.Name,
			Path: Value.Name,
			Crc:  Value.Crc,
		})
	}

	return CatalogDatas
}

/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:34:33
 * @LastEditTime: 2023-03-03 23:40:14
 * @LastEditors: nijineko
 * @Description: 读取MediaResources的CataLog文件到标准结构体
 * @FilePath: \DataDownload\internal\Catalog\MediaResources.go
 */
package Catalog

import (
	"BlueArchiveDataDownload/internal/HTTP"
	"encoding/json"
)

type MediaResourcesOrigin struct {
	Table map[string]struct {
		IsInbuild bool   `json:"isInbuild"` // 是否为内置文件
		IsChanged bool   `json:"isChanged"` // 是否被修改
		MediaType int    `json:"mediaType"` // 媒体类型 (1:音频, 2:视频, 3:图片)
		Path      string `json:"path"`      // 路径
		FileName  string `json:"fileName"`  // 文件名
		Bytes     int    `json:"bytes"`     // 文件大小
		Crc       int    `json:"Crc"`       // 文件CRC
	} `json:"Table"`
	MediaList any `json:"MediaList"`
}

/**
 * @description: 获取MediaResources的CataLog数据
 * @param {string} AddressablesCatalogUrlRoot 资源地址
 * @return {[]Data} CatLog数据
 * @return {error} 错误信息
 */
func GetMediaResources(AddressablesCatalogUrlRoot string) ([]Data, error) {
	Body, _, err := HTTP.Get(AddressablesCatalogUrlRoot + MediaResourcesCataLogPath)
	if err != nil {
		return nil, err
	}

	// 使用原始结构体解析JSON
	var MediaResources MediaResourcesOrigin
	err = json.Unmarshal(Body, &MediaResources)
	if err != nil {
		return nil, err
	}

	// 转换为标准结构体
	var MediaResourcesData []Data
	for _, Value := range MediaResources.Table {
		MediaResourcesData = append(MediaResourcesData, Data{
			Name: Value.FileName,
			Path: Value.Path,
			Crc:  Value.Crc,
		})
	}

	return MediaResourcesData, nil
}

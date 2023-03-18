/*
 * @Author: nijineko
 * @Date: 2023-03-16 15:59:52
 * @LastEditTime: 2023-03-19 02:00:19
 * @LastEditors: nijineko
 * @Description: 更新模块
 * @FilePath: \DataDownload\internal\Update\Update.go
 */
package Update

import (
	"BlueArchiveDataDownload/internal/Catalog"
)

/**
 * @description: 对比两个Catalog的Crc，获取需要更新的文件
 * @param {[]Catalog.Data} SrcData 源数据
 * @param {[]Catalog.Data} DestData 目标数据
 * @return {[]Catalog.Data} 差异数据
 */
func CompareDataCrc(SrcData []Catalog.Data, DestData []Catalog.Data) []Catalog.Data {
	// 将源数据转换为以文件Path为Key，文件Crc为Value的Map
	SrcDataMap := make(map[string]uint32)
	for _, Value := range SrcData {
		SrcDataMap[Value.Path] = Value.Crc
	}

	// 遍历对比目标数据
	var DifferenceDatas []Catalog.Data
	for _, Value := range DestData {
		// 如果目标数据的Path在源数据中不存在，则将其添加到差异数据中
		if _, Find := SrcDataMap[Value.Path]; !Find {
			DifferenceDatas = append(DifferenceDatas, Value)
			continue
		}

		// 判断目标数据的Crc是否与源数据的Crc相同
		if SrcDataMap[Value.Path] != Value.Crc {
			DifferenceDatas = append(DifferenceDatas, Value)
		}
	}

	return DifferenceDatas
}

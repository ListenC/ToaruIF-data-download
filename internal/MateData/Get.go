package MateData

import (
	"DataDownload/internal/HTTP"
	"encoding/json"
	"fmt"
)

var MateDataURL = "https://yostar-serverinfo.bluearchiveyostar.com/%s.json" // 元数据地址

// 元数据结果(注释基于推测)
type MateData struct {
	ConnectionGroups []struct {
		Name                       string `json:"Name"`
		ManagementDataURL          string `json:"ManagementDataUrl"` // 公告数据地址
		IsProductionAddressables   bool   `json:"IsProductionAddressables"`
		APIURL                     string `json:"ApiUrl"`                     // API地址
		GatewayURL                 string `json:"GatewayUrl"`                 // 网关地址
		KibanaLogURL               string `json:"KibanaLogUrl"`               // 日志地址
		ProhibitedWordBlackListURI string `json:"ProhibitedWordBlackListUri"` // 词语黑名单地址
		ProhibitedWordWhiteListURI string `json:"ProhibitedWordWhiteListUri"` // 词语白名单地址
		CustomerServiceURL         string `json:"CustomerServiceUrl"`
		OverrideConnectionGroups   []struct {
			Name                       string `json:"Name"`                       // 版本号
			AddressablesCatalogURLRoot string `json:"AddressablesCatalogUrlRoot"` // 资源地址
		} `json:"OverrideConnectionGroups"` // 数据地址
		BundleVersion string `json:"BundleVersion"` // 资源版本号
	} `json:"ConnectionGroups"`
}

/**
 * @description: 获取元数据
 * @return {MateData} 元数据
 * @return {error} 错误信息
 */
func Get(Version string) (MateData, error) {
	// 获取元数据
	Body, _, err := HTTP.Get(fmt.Sprintf(MateDataURL, Version))
	if err != nil {
		return MateData{}, err
	}

	// 解析元数据
	var Data MateData
	err = json.Unmarshal(Body, &Data)
	if err != nil {
		return MateData{}, err
	}

	return Data, nil
}

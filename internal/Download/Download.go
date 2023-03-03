/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:43:42
 * @LastEditTime: 2023-03-04 01:43:21
 * @LastEditors: nijineko
 * @Description: 下载文件
 * @FilePath: \DataDownload\internal\Download\Download.go
 */
package Download

import (
	"BlueArchiveDataDownload/internal/Catalog"
	"BlueArchiveDataDownload/internal/Flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/pierrec/xxHash/xxHash64"
	"github.com/schollz/progressbar/v3"
)

/**
 * @description: 下载文件并显示进度条
 * @param {string} URL 文件地址
 * @param {string} SavePath 保存路径
 * @return {int64} 文件大小
 * @return {error} 错误信息
 */
func File(URL string, SavePath string) (int64, error) {
	Request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return 0, err
	}

	// 发起请求
	Client, err := http.DefaultClient.Do(Request)
	if err != nil {
		return 0, err
	}
	defer Client.Body.Close()

	err = CreateFolder(path.Dir(SavePath))
	if err != nil {
		return 0, err
	}
	File, err := os.OpenFile(SavePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer File.Close()

	// 创建进度条
	Progressbar := progressbar.NewOptions(int(Client.ContentLength),
		progressbar.OptionEnableColorCodes(true),                  // 启用颜色
		progressbar.OptionShowBytes(true),                         // 显示速度
		progressbar.OptionFullWidth(),                             // 宽度设置为Full
		progressbar.OptionSetDescription("正在下载: "+path.Base(URL)), // 设置描述
		progressbar.OptionClearOnFinish(),                         // 完成后清除进度条
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[light_blue]=[reset]", // 设置进度条的样式(中间)
			SaucerHead:    "[light_blue]>[reset]", // 设置进度条的样式(头部)
			SaucerPadding: " ",                    // 设置进度条的样式(空白部分)
			BarStart:      "[",                    // 设置进度条的开头
			BarEnd:        "]",                    // 设置进度条的结尾
		}))

	// 写入文件
	Size, err := io.Copy(io.MultiWriter(File, Progressbar), Client.Body)
	if err != nil {
		return 0, err
	}

	return Size, nil
}

/**
 * @description: 创建文件夹
 * @param {string} Path 文件夹路径
 * @return {error} 错误信息
 */
func CreateFolder(Path string) error {
	_, err := os.Stat(Path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: 下载资源文件
 * @param {[]Catalog.Data} CatalogData Catalog文件数据
 * @param {string} PathURL 资源文件地址
 * @param {string} SavePath 保存路径
 * @param {bool} xxHash OriginalFileSave模式下是否使用xxHash64计算文件名
 * @return {error} 错误信息
 */
func Resource(CatalogData []Catalog.Data, PathURL string, SavePath string, xxHash bool) error {
	// 遍历CatalogData
	for _, Value := range CatalogData {
		var FilePath string
		// 判断是否以原始文件的名字和路径保存
		if Flag.Data.OriginalFileSave && xxHash {
			// 计算文件名
			FileName := fmt.Sprintf("%d", xxHash64.Checksum([]byte(Value.Name), 0))
			FilePath = path.Join(SavePath, FileName)
		} else {
			FilePath = path.Join(SavePath, Value.Path)
		}

		// 下载文件
		Size, err := File(PathURL+Value.Path, FilePath)
		if err != nil {
			return err
		}

		fmt.Printf("文件 %s 下载完成，大小为 %dbytes\n", Value.Path, Size)
	}

	return nil
}

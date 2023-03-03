/*
 * @Author: nijineko
 * @Date: 2023-03-03 23:43:42
 * @LastEditTime: 2023-03-04 00:43:12
 * @LastEditors: nijineko
 * @Description: 下载文件
 * @FilePath: \DataDownload\internal\Download\Download.go
 */
package Download

import (
	"io"
	"net/http"
	"os"
	"path"

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

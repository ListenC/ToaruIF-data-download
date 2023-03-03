/*
 * @Author: NyanCatda
 * @Date: 2023-03-03 22:42:20
 * @LastEditTime: 2023-03-03 22:46:20
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \DataDownload\main.go
 */
package main

import "DataDownload/internal/Flag"

func init() {
	// 初始化参数
	err := Flag.Init()
	if err != nil {
		panic(err)
	}
}

func main() {

}

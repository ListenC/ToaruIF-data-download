<!--
 * @Author: nijineko
 * @Date: 2023-03-04 02:24:45
 * @LastEditTime: 2023-03-21 12:57:47
 * @LastEditors: nijineko
 * @Description: 
 * @FilePath: \DataDownload\README.md
-->
# blue-archive-data-download
BlueArchive游戏文件下载工具，使用Golang实现
![example](example.gif)

# 使用
前往[Releases](https://github.com/nijinekoyo/blue-archive-data-download/releases)下载程序后，执行以下命令即可
```
.\BlueArchiveDataDownload.exe -data_version xxx_xxxxxxxxxxxxxxxxxxxx
```

# 参数
```
-data_version string
    指定数据包版本
-original_file_save
    是否以原始文件的名字和路径保存
-max_pool int
    最大并发数 (default 10)
-filter string
    字符串过滤器，只下载包含该字符串的文件
-ignore_inbuild
    忽略APK内置文件
-save_catalog
    是否保存Catalog文件 (default true)
-update
    以更新模式启动程序
-update_copy
    更新模式下复制新文件到UpdateData目录下储存
-update_clean
    更新模式下清理过期文件
-asset_bundls
    下载/更新AssetBundls文件
-media_resources
    下载/更新MediaResources文件
-table_bundles
    下载/更新TableBundles文件
```

# 更新模式
更新模式会通过对比本地文件与远程Calalog记录的CRC值来判断文件是否需要更新，如果需要更新则会自动开始下载新的文件  
如果你需要使用更新模式，则必须指定和下载模式相同的`-original_file_save`参数来更新文件

# 关于数据包版本
暂时没有找到正确获取数据包版本的方案，目前只能尝试通过抓包获取

# 声明
此项目仅供学习使用，项目与Yostar或Nexon等其他第三方没有任何关系
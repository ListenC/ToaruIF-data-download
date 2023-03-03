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
-asset_bundls
    下载AssetBundls文件
-media_resources
    下载MediaResources文件
-table_bundles
    下载TableBundles文件
```

# 关于数据包版本
暂时没有找到正确获取数据包版本的方案，可以参考其他大佬拆出来的数据包版本或者尝试抓包

# 声明
此项目仅供学习使用，项目与Yostar或Nexon等其他第三方没有任何关系
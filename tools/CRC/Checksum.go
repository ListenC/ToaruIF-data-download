package CRC

import (
	"hash/crc32"
	"os"
)

/**
 * @description: 计算文件的CRC32
 * @param {string} FilePath 文件路径
 * @return {uint32} CRC32
 */
func Checksum(FilePath string) uint32 {
	FileBody, err := os.ReadFile(FilePath)
	if err != nil {
		return 0
	}
	CRC := crc32.ChecksumIEEE(FileBody)

	return CRC
}

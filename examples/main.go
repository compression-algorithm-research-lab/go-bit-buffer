package main

import (
	"fmt"
	bit_buffer "github.com/compression-algorithm-research-lab/go-bit-buffer"
)

func main() {

	// 创建一块缓存空间，可以一个bit一个bit的往里面写
	buffer := bit_buffer.New().WriteBit(1).WriteBit(0)

	// 转为二进制字符串
	binaryString := buffer.ToBinaryString()
	fmt.Println(binaryString) // Output: 10000000

	// 转为字节数组
	bytes := buffer.Bytes()
	fmt.Println(bytes) // Output: [128]

}

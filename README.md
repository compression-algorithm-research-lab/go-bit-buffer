# go-bit-buffer

# 一、这是什么？

可以申请一块bit缓存空间，可以一个bit一个bit的写入，用于方便进行位操作。

# 二、安装

```bash
go get -u github.com/compression-algorithm-research-lab/go-bit-buffer
```

# 三、API示例

```go
package main

import (
	"fmt"
	bit_buffer "github.com/compression-algorithm-research-lab/go-bit-buffer"
)

func main() {

	// 创建一块缓存空间，可以一个bit一个bit的往里面写
	buffer := bit_buffer.New().Append(1).Append(0)

	// 转为二进制字符串
	binaryString := buffer.ToBinaryString()
	fmt.Println(binaryString) // Output: 10000000

	// 转为字节数组
	bytes := buffer.Bytes()
	fmt.Println(bytes) // Output: [128]

}
```




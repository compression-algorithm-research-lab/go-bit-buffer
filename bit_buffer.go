package bit_buffer

import (
	"strings"
)

const BitBeginOffset = 7

// BitBuffer 提供bit级别的缓存
type BitBuffer struct {

	// 当前写到哪个下标了
	nextWriteIndex int

	// 当前写到的offset
	nextWriteBitOffset int

	// 当前的写缓存
	buffer []byte
}

// New 创建一个可以用来写数据的位缓存空间
func New() *BitBuffer {
	return &BitBuffer{
		nextWriteIndex:     0,
		nextWriteBitOffset: BitBeginOffset,
		buffer:             make([]byte, 0),
	}
}

// Append 往缓存空间追加一个bit，会获取 bitValue 的最低位的bit值，1或者0
func (x *BitBuffer) Append(bitValue int) *BitBuffer {

	// 字节数组在需要的时候扩容
	if x.nextWriteIndex >= len(x.buffer) {
		x.buffer = append(x.buffer, byte(0x0))
	}

	// 写入下标
	x.buffer[x.nextWriteIndex] ^= byte((bitValue & 0x01) << x.nextWriteBitOffset)

	// 移动指针
	x.nextWriteBitOffset--
	if x.nextWriteBitOffset < 0 {
		x.nextWriteBitOffset = BitBeginOffset
		x.nextWriteIndex++
	}

	return x
}

// AppendByte 往缓存空间一次追加一个byte
func (x *BitBuffer) AppendByte(b byte) *BitBuffer {
	for offset := BitBeginOffset; offset >= 0; offset-- {
		x.Append((int(b) >> offset) & 0x1)
	}
	return x
}

// Bytes 返回当前缓存空间所对应的字节数组
func (x *BitBuffer) Bytes() []byte {
	return x.buffer
}

// ToBinaryString 把当前的缓存空间转换为二进制字符串
func (x *BitBuffer) ToBinaryString() string {
	buff := strings.Builder{}
	for _, b := range x.buffer {
		for offset := BitBeginOffset; offset >= 0; offset-- {
			if ((b >> offset) & 0x1) == 1 {
				buff.WriteRune('1')
			} else {
				buff.WriteRune('0')
			}
		}
	}
	return buff.String()
}

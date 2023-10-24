package bit_buffer

import (
	"strings"
)

const BitBeginOffset = 7

// BitBuffer 在bit级别提供一些便于操作的API
type BitBuffer struct {

	// 下一个要操作的byte的下标
	nextIndex int

	// 下一个要操作的byte的offset
	nextBitOffset int

	// buff当前对应的字节数组
	bytes []byte
}

// New 创建一个可以用来写数据的位缓存空间
func New() *BitBuffer {
	return &BitBuffer{
		nextIndex:     0,
		nextBitOffset: BitBeginOffset,
		bytes:         make([]byte, 1),
	}
}

// SetBytes 设置buff底层的字节数组
func (x *BitBuffer) SetBytes(bytes []byte) *BitBuffer {
	x.bytes = bytes
	return x
}

// SeekFirst 移动操作指针
func (x *BitBuffer) SeekFirst() *BitBuffer {
	return x.Seek(0)
}

// SeekLast 移动指针到最后一个位置
func (x *BitBuffer) SeekLast() *BitBuffer {
	return x.Seek(len(x.bytes) * 8)
}

// Seek 移动指针
func (x *BitBuffer) Seek(offset int) *BitBuffer {

	// 保证空间是足够的
	x.nextIndex = offset / 8
	for len(x.bytes) <= x.nextIndex {
		x.bytes = append(x.bytes, byte(0))
	}

	// 把便宜设置正确
	x.nextBitOffset = BitBeginOffset - (offset % 8)

	return x
}

// GetSeek 获取当前写指针
func (x *BitBuffer) GetSeek() int {
	return x.nextIndex*8 + (BitBeginOffset - x.nextBitOffset)
}

// Capacity 获取当前的容量
func (x *BitBuffer) Capacity() int {
	return len(x.bytes) * 8
}

// WriteBit 往缓存空间追加一个bit，会获取 bitValue 的最低位的bit值，1或者0
func (x *BitBuffer) WriteBit(bitValue int) *BitBuffer {

	// 字节数组在需要的时候扩容
	if x.nextIndex >= len(x.bytes) {
		x.bytes = append(x.bytes, byte(0x0))
	}

	// 写入对应偏移的bit值，其他的bit位保持原值
	x.bytes[x.nextIndex] = byte((bitValue&0x01)<<x.nextBitOffset) | (x.bytes[x.nextIndex] & ((0x1 << x.nextBitOffset) ^ 0xFF))

	// 移动指针
	x.nextBitOffset--
	if x.nextBitOffset < 0 {
		x.nextBitOffset = BitBeginOffset
		x.nextIndex++
	}

	return x
}

// WriteByte 往缓存空间一次追加一个byte
func (x *BitBuffer) WriteByte(b byte) *BitBuffer {
	for offset := BitBeginOffset; offset >= 0; offset-- {
		x.WriteBit((int(b) >> offset) & 0x1)
	}
	return x
}

// ReadBit 从seek读取一个bit
func (x *BitBuffer) ReadBit() int {

	// 读取对应offset的bit值
	value := (x.bytes[x.nextIndex] & (0x1 << x.nextBitOffset)) >> x.nextBitOffset

	// offset往后移动
	x.nextBitOffset--
	if x.nextBitOffset < 0 {
		x.nextBitOffset = BitBeginOffset
		x.nextIndex++
	}

	return int(value)
}

// IsAtLast 是否已经存在最后一个位置了
func (x *BitBuffer) IsAtLast() bool {
	return x.GetSeek() == x.Capacity()
}

// Bytes 返回当前缓存空间所对应的字节数组
func (x *BitBuffer) Bytes() []byte {
	return x.bytes
}

// ToBinaryString 把当前的缓存空间转换为二进制字符串
func (x *BitBuffer) ToBinaryString() string {
	buff := strings.Builder{}
loop:
	for index, b := range x.bytes {
		for offset := BitBeginOffset; offset >= 0; offset-- {

			// 没有写内容的部分就不转换到字符串中了
			if index == x.nextIndex && offset <= x.nextBitOffset {
				break loop
			}

			if ((b >> offset) & 0x1) == 1 {
				buff.WriteRune('1')
			} else {
				buff.WriteRune('0')
			}
		}
	}
	return buff.String()
}

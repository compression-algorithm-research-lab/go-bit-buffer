package bit_buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitBuffer_WriteBit(t *testing.T) {

	buffer := New()
	assert.Equal(t, "1", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "10", buffer.WriteBit(0).ToBinaryString())
	assert.Equal(t, "101", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "1010", buffer.WriteBit(0).ToBinaryString())
	assert.Equal(t, "10101", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "101010", buffer.WriteBit(0).ToBinaryString())
	assert.Equal(t, "1010101", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "10101011", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "101010111", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "1010101111", buffer.WriteBit(1).ToBinaryString())

}

func TestBitBuffer_WriteByte(t *testing.T) {

	buffer := New()
	assert.Equal(t, "11111111", buffer.WriteByte(0xFF).ToBinaryString())
	assert.Equal(t, "111111110", buffer.WriteBit(0).ToBinaryString())
	assert.Equal(t, "1111111101", buffer.WriteBit(1).ToBinaryString())
	assert.Equal(t, "111111110111111111", buffer.WriteByte(0xFF).ToBinaryString())

}

func TestBitBuffer_SetBytes(t *testing.T) {
	buffer := New().SetBytes([]byte{0x1, 0x2})
	assert.Equal(t, []byte{0x1, 0x2}, buffer.Bytes())
}

func TestBitBuffer_Capacity(t *testing.T) {
	buffer := New()
	assert.Equal(t, 8, buffer.Capacity())
	assert.Equal(t, 8, buffer.WriteBit(0x1).Capacity())
	assert.Equal(t, 16, buffer.WriteByte(0xFF).Capacity())
}

func TestBitBuffer_GetSeek(t *testing.T) {
	buffer := New()
	assert.Equal(t, 0, buffer.GetSeek())
	assert.Equal(t, 1, buffer.WriteBit(0x1).GetSeek())
}

func TestBitBuffer_Seek(t *testing.T) {
	buffer := New().WriteByte(0xFF)
	assert.Equal(t, "11111111", buffer.ToBinaryString())
	assert.Equal(t, "", buffer.Seek(0).ToBinaryString())
	assert.Equal(t, "1", buffer.Seek(1).ToBinaryString())
	assert.Equal(t, "11", buffer.Seek(2).ToBinaryString())
	assert.Equal(t, "111", buffer.Seek(3).ToBinaryString())
	assert.Equal(t, "1111111", buffer.Seek(7).ToBinaryString())
	assert.Equal(t, "11111111", buffer.Seek(8).ToBinaryString())
	assert.Equal(t, "111111110", buffer.Seek(9).ToBinaryString())
	assert.Equal(t, "111111110000000", buffer.Seek(15).ToBinaryString())
	assert.Equal(t, "1111111100000000", buffer.Seek(16).ToBinaryString())
	assert.Equal(t, "101111110000000000000000", buffer.Seek(1).WriteBit(0).SeekLast().ToBinaryString())
}

func TestBitBuffer_SeekFirst(t *testing.T) {
	buffer := New().WriteByte(0xFF)
	assert.Equal(t, "11111111", buffer.ToBinaryString())
	assert.Equal(t, "", buffer.SeekFirst().ToBinaryString())
	assert.Equal(t, "1", buffer.WriteBit(0x1).ToBinaryString())
}

func TestBitBuffer_SeekLast(t *testing.T) {
	buffer := New().WriteByte(0xFF)
	assert.Equal(t, "11111111", buffer.ToBinaryString())
	assert.Equal(t, "11111111", buffer.SeekLast().ToBinaryString())
	assert.Equal(t, "111111111", buffer.WriteBit(0x1).ToBinaryString())
}

func TestBitBuffer_ReadBit(t *testing.T) {
	buffer := New().WriteByte(0xF0).SeekFirst()

	assert.Equal(t, 1, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 1, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 1, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 1, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 0, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 0, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 0, buffer.ReadBit())
	assert.False(t, buffer.IsAtLast())

	assert.Equal(t, 0, buffer.ReadBit())
	assert.True(t, buffer.IsAtLast())
}

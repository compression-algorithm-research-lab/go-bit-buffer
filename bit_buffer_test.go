package bit_buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitBuffer_Append(t *testing.T) {

	buffer := New()
	assert.Equal(t, "1", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "101", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "1010", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "10101", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "101010", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "1010101", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10101011", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "101010111", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "1010101111", buffer.Append(1).ToBinaryString())

}

func TestBitBuffer_AppendByte(t *testing.T) {

	buffer := New()
	assert.Equal(t, "11111111", buffer.AppendByte(0xFF).ToBinaryString())
	assert.Equal(t, "111111110", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "1111111101", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "111111110111111111", buffer.AppendByte(0xFF).ToBinaryString())

}

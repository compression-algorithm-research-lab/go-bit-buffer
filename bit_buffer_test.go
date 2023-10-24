package bit_buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitBuffer_Append(t *testing.T) {

	buffer := New()
	assert.Equal(t, "10000000", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10000000", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "10100000", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10100000", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "10101000", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10101000", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "10101010", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "10101011", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "1010101110000000", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "1010101111000000", buffer.Append(1).ToBinaryString())

}

func TestBitBuffer_AppendByte(t *testing.T) {

	buffer := New()
	assert.Equal(t, "11111111", buffer.AppendByte(0xFF).ToBinaryString())
	assert.Equal(t, "1111111100000000", buffer.Append(0).ToBinaryString())
	assert.Equal(t, "1111111101000000", buffer.Append(1).ToBinaryString())
	assert.Equal(t, "111111110111111111000000", buffer.AppendByte(0xFF).ToBinaryString())

}

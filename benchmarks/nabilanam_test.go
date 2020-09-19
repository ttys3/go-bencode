package benchmarks

import (
	"testing"

	"github.com/nabilanam/bencode/decoder"
	"github.com/nabilanam/bencode/encoder"
	"github.com/stretchr/testify/assert"
)

func BenchmarkNabilanamBencodeMarshal(b *testing.B) {
	namedBench(b, "nabilanam/bencode", func() {
		encoder.New(marshalTestData).Encode()
	})
}

func BenchmarkNabilanamBencodeUnmarshal(b *testing.B) {
	namedBench(b, "nabilanam/bencode", func() {
		decoder.New(unmarshalTestData).Decode()
	})
}

func TestNabilanamBencodeUnmarshal(t *testing.T) {
	assert := assert.New(t)
	nabilanam := decoder.New(unmarshalTestData).Decode()
	if !assert.Equal(nabilanam, marshalTestDataWithStrings) {
		return
	}
}

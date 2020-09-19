package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/jackpal/bencode-go"
)

func BenchmarkJackpalBencodeMarshal(b *testing.B) {
	namedBench(b, "jackpal/bencode-go", func() {
		bencode.Marshal(bytes.NewBuffer(nil), marshalTestData)
	})
}

func BenchmarkJackpalBencodeUnmarshal(b *testing.B) {
	namedBench(b, "jackpal/bencode-go", func() {
		bencode.Decode(bytes.NewReader(unmarshalTestData))
	})
}

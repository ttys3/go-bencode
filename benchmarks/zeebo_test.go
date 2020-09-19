package benchmarks

import (
	"testing"

	bencode "github.com/zeebo/bencode"
)

func BenchmarkZeeboBencodeMarshal(b *testing.B) {
	namedBench(b, "zeebo/bencode", func() {
		bencode.EncodeBytes(marshalTestData)
	})
}

func BenchmarkZeeboBencodeUnmarshal(b *testing.B) {
	namedBench(b, "zeebo/bencode", func() {
		var torrent interface{}
		bencode.DecodeBytes(unmarshalTestData, &torrent)
	})
}

package benchmarks

import (
	"bytes"
	"testing"

	bencode "github.com/marksamman/bencode"
)

func BenchmarkMarksammanBencodeMarshal(b *testing.B) {
	namedBench(b, "marksamman/bencode", func() {
		bencode.Encode(marshalTestData)
	})
}

func BenchmarkMarksammanBencodeUnmarshal(b *testing.B) {
	namedBench(b, "marksamman/bencode", func() {
		bencode.Decode(bytes.NewReader(unmarshalTestData))
	})
}

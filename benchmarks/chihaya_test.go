package benchmarks

import (
	"testing"

	bencode "github.com/chihaya/chihaya/frontend/http/bencode"
)


func BenchmarkChihayaBencodeMarshal(b *testing.B) {
	namedBench(b, "chihaya/bencode", func() {
		bencode.Marshal(marshalTestData)
	})
}

func BenchmarkChihayaBencodeUnmarshal(b *testing.B) {
	namedBench(b, "chihaya/bencode", func() {
		bencode.Unmarshal(unmarshalTestData)
	})
}

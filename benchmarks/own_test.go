package benchmarks

import (
	"testing"

	bencode "github.com/IncSW/go-bencode"
)

func BenchmarkMarshal(b *testing.B) {
	namedBench(b, "IncSW/go-bencode", func() {
		bencode.Marshal(marshalTestData)
	})
}

func BenchmarkUnmarshal(b *testing.B) {
	namedBench(b, "IncSW/go-bencode", func() {
		bencode.Unmarshal(unmarshalTestData)
	})
}

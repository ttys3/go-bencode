# go-bencode
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/IncSW/go-bencode.svg?style=flat-square)](https://travis-ci.org/IncSW/go-bencode)
[![Coverage Status](https://img.shields.io/coveralls/IncSW/go-bencode/master.svg?style=flat-square)](https://coveralls.io/github/IncSW/go-bencode)
[![Go Report Card](https://goreportcard.com/badge/github.com/IncSW/go-bencode?style=flat-square)](https://goreportcard.com/report/github.com/IncSW/go-bencode)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/IncSW/go-bencode)

## Installation

`go get github.com/IncSW/go-bencode`

```go
import bencode "github.com/IncSW/go-bencode"
```

## Quick Start

```go
data, err := bencode.Marshal(value)
```

```go
data, err := bencode.Unmarshal(value)
```

## Performance

### Go 1.15.2, Fedora 32, i7-8700K

```bash
goos: linux
goarch: amd64
pkg: benchmarks
BenchmarkChihayaBencodeMarshal/library=chihaya/bencode-12         	  695912	      1748 ns/op	    1010 B/op	      53 allocs/op
BenchmarkChihayaBencodeUnmarshal/library=chihaya/bencode-12       	  433002	      2772 ns/op	    5904 B/op	      61 allocs/op
BenchmarkJackpalBencodeMarshal/library=jackpal/bencode-go-12      	  234519	      4455 ns/op	    2032 B/op	      45 allocs/op
BenchmarkJackpalBencodeUnmarshal/library=jackpal/bencode-go-12    	  526748	      2554 ns/op	    1712 B/op	      59 allocs/op
BenchmarkMarksammanBencodeMarshal/library=marksamman/bencode-12   	 1568473	       793 ns/op	     400 B/op	       8 allocs/op
BenchmarkMarksammanBencodeUnmarshal/library=marksamman/bencode-12 	  418111	      2605 ns/op	    5792 B/op	      54 allocs/op
BenchmarkNabilanamBencodeMarshal/library=nabilanam/bencode-12     	  593919	      2065 ns/op	    1216 B/op	      44 allocs/op
BenchmarkNabilanamBencodeUnmarshal/library=nabilanam/bencode-12   	  978747	      1342 ns/op	    1264 B/op	      39 allocs/op
BenchmarkMarshal/library=IncSW/go-bencode-12                      	 2682212	       453 ns/op	      16 B/op	       2 allocs/op
BenchmarkUnmarshal/library=IncSW/go-bencode-12                    	  831754	      1364 ns/op	    1248 B/op	      25 allocs/op
BenchmarkZeeboBencodeMarshal/library=nabilanam/bencode-12         	  270471	      4858 ns/op	    1392 B/op	      33 allocs/op
BenchmarkZeeboBencodeUnmarshal/library=nabilanam/bencode-12       	  194533	      5967 ns/op	    6576 B/op	      99 allocs/op
PASS
ok  	benchmarks	16.971s
```

### Go 1.13.1, Debian 9.1, i7-7700

### Marshal

| Library            |    Time    | Bytes Allocated | Objects Allocated |
| :----------------- | :--------: | :-------------: | :---------------: |
| IncSW/go-bencode   | 485 ns/op  |     16 B/op     |    2 allocs/op    |
| marksamman/bencode | 897 ns/op  |    400 B/op     |    8 allocs/op    |
| chihaya/bencode    | 1846 ns/op |    1010 B/op    |   53 allocs/op    |
| nabilanam/bencode  | 2027 ns/op |    1216 B/op    |   44 allocs/op    |
| jackpal/bencode-go | 4968 ns/op |    2128 B/op    |   57 allocs/op    |
| zeebo/bencode      | 5147 ns/op |    1488 B/op    |   45 allocs/op    |

### Unmarshal

| Library            |    Time    | Bytes Allocated | Objects Allocated |
| :----------------- | :--------: | :-------------: | :---------------: |
| nabilanam/bencode  | 1341 ns/op |    1264 B/op    |   39 allocs/op    |
| IncSW/go-bencode   | 1435 ns/op |    1248 B/op    |   25 allocs/op    |
| jackpal/bencode-go | 2652 ns/op |    1712 B/op    |   59 allocs/op    |
| marksamman/bencode | 2877 ns/op |    5920 B/op    |   66 allocs/op    |
| chihaya/bencode    | 2896 ns/op |    5904 B/op    |   61 allocs/op    |
| zeebo/bencode      | 6595 ns/op |    6576 B/op    |   99 allocs/op    |

## License

[MIT License](LICENSE).

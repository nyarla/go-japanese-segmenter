go-japanese-segmenter
=====================

A TinySegmenter implementation of golang but zero allocation inside this library.

Usage
-----

```golang
package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/nyarla/go-japanese-segmenter/dicts/tinyseg"
	"github.com/nyarla/go-japanese-segmenter/segmenter"
)

func main() {
	src := strings.NewReader("今日は良い天気ですね")
	dst := new(strings.Builder)
	dict := segmenter.BiasCalculatorFunc(tinyseg.CalculateBias)
	seg := segmenter.New(dst, src)

	for {
		err := seg.Segment(dict)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Println(dst.String())
		dst.Reset()
	}

	fmt.Println(dst.String())
	dst.Reset()

	// Output:
	// 今日
	// は
	// 良い
	// 天気
	// ですね
}

```

Benchmark
---------

  * OS: NixOS on Hyper-V with Windows 10 Pro (build 19042.928)
  * Kernel: Linux 5.10.19
  * Machine: Desktop PC
  * CPU: AMD Ryzen 9 3950X (assign 16 threads)
  * Memory: assign 32GB 

```sh
$ go test -bench BenchmarkSegmenter -test.count=10
goos: linux
goarch: amd64
pkg: github.com/nyarla/go-japanese-segmenter/segmenter
cpu: AMD Ryzen 9 3950X 16-Core Processor
BenchmarkSegmenter-16             213902              5556 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             214588              5538 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             215530              5593 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             214365              5560 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             214882              5551 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             215779              5546 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             215514              5577 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             213500              5649 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             213333              5652 ns/op               0 B/op          0 allocs/op
BenchmarkSegmenter-16             199754              5645 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/nyarla/go-japanese-segmenter/segmenter       12.498s
```

Documentation
-------------

  * <https://pkg.go.dev/github.com/nyarla/go-japanese-segmenter@v0.1.0>

How to rebuild dictionary code (bias calculator)
------------------------------------------------

```sh
$ cat dictionary.json
{
  "BC1": {
    "HH": 6
    ...
  },
  ...
}

$ go get github.com/go-japanese-segmenter/cmd/tinydictgen
$ tinydictgen -pkg mydict -bias "-332" -json ./dictionary.json
$ go fmt
$ cat mydict_generated.go
// this code is auto-generated. DO NOT EDIT.
package defaults

const initialBias = -332

func CalculateBias(p1, p2, p3, r1, r2, r3, r4, r5, r6, t1, t2, t3, t4, t5, t6 rune) int64 {
	n := int64(initialBias)
...
```

Running Benchmark
-----------------

```sh
$ cd /path/to/go-japanese-segmenter/segmenter
$ curl -LO http://www.genpaku.org/timemachine/timemachineu8j.txt
$ go test -bench BenchmarkSegmenter -test.count=10
```

Copyrights Notice
-----------------

This library contians dictionary data (`dicts/tinyseg.json`) splitted from original [TinySegmenter.js](http://chasen.org/~taku/software/TinySegmenter/),
and dictionary code (`dicts/tinyseg_generated.go`) is generated from splitted dictionary data. 

Please looking at _Licenses section_ in this `README.md` about license of original TinySegmenter.js 

Licenses
--------

### Original Implemenation written by JavaScript

  * [TinySegmenter.js](http://chasen.org/~taku/software/TinySegmenter/)

```
Copyright (c) 2008, Taku Kudo

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright notice,
this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.
    * Neither the name of the <ORGANIZATION> nor the names of its
contributors may be used to endorse or promote products derived from this
software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

### go-japanese-segmenter

```
Copyright (c) 2019, Naoki OKAMURA <nyarla@thotep.net>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright notice,
this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.
    * Neither the name of the <ORGANIZATION> nor the names of its
contributors may be used to endorse or promote products derived from this
software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

Contacts
--------

  * Naoki OKAMURA <nyarla@kalaclista.com>


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

	"github.com/nyarla/go-japanese-segmenter/defaults"
	"github.com/nyarla/go-japanese-segmenter/segmenter"
)

func main() {
	src := strings.NewReader("今日は良い天気ですね")
	dst := new(strings.Builder)
	dict := segmenter.BiasCalculatorFunc(defaults.CalculateBias)
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

Documentation
-------------

  * <https://godoc.org/github.com/nyarla/go-japanese-segmenter/segmenter>

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

$ go get github.com/go-japanese-segmenter/cmds/tinydictgen
$ tinydictgen -pkg mydict -bias "-332" -json ./dictionary.json
$ got fmt
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
$ cd /path/to/go-japanese-segmenter
$ sh ./download.sh # or cd segmenter && curl -LO http://www.genpaku.org/timemachine/timemachineu8j.txt
$ go test -bench Benchmark ./...
```

Copyrights Notice
-----------------

This library contians dictionary data (`defaults/defaults.json`) splitted from original [TinySegmenter.js](http://chasen.org/~taku/software/TinySegmenter/),
and dictionary code (`defaults/defaults_generated.go`) is generated from splitted dictionary data. 

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

  * Naoki OKAMURA <nyarla@thotep.net>


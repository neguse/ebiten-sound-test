# ebiten sound performance test

## Build

export GO111MODULE=off
gopherjs build -o main.js main.go
GOOS=js GOARCH=wasm go build -o main.wasm main.go
python serve.py
open http://localhost:8000/main.html # for wasm version
open http://localhost:8000/js.html   # for js version

## License

wasm_exec.jsは、Goのものを利用しています。

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style

damage.wav は、 [go-inovation](https://github.com/hajimehoshi/go-inovation/)のものを利用しています。
statik.goにも damage.wav ファイルが含まれます。

Copyright 2015 Omega (The original auther)
Copyright 2015 Haneda3 (JavaScript port)
Copyright 2015 Hajime Hoshi (Go port)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

それ以外のファイルはご自由に.

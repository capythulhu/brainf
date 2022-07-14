# brainf

**brainf** is a [brainfuck](https://en.wikipedia.org/wiki/Brainfuck) compiler/interpreter written in Go.

## 📃 Prerequisites

You must have [Go 1.16+](https://go.dev/doc/install) in order to compile **brainf**.

## ⚙️ Compiling

First, clone this repository with the following command:

```
git clone https://github.com/thzoid/brainf.git
```

Then, compile it:

```
cd brainf
go build
```

## 🚀 Usage

**brainf** can be used to run `.bf` scripts:

```
./brainf <relative file path>
```

You can try to run one of the example codes like this:

```
./brainf examples/sum.bf
3
2
5
```

## 📝 License

Copyright © 2022 [Thiago Antunes](https://github.com/thzoid).<br />
This project is [MIT](https://github.com/kefranabg/readme-md-generator/blob/master/LICENSE) licensed.

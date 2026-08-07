# go-conv

常见进制与编码转换的小工具：十进制/十六进制/二进制互转，以及 URL、Base64、十六进制编解码。

## 安装

```bash
go build -o go-conv.exe
```

## 用法

```bash
go-conv dec2hex 255      # ff
go-conv hex2dec ff       # 255
go-conv dec2bin 5        # 101
go-conv bin2dec 101      # 5
go-conv urlenc "a b"     # a+b
go-conv urldec "a%2Fb"   # a/b
go-conv b64enc ab        # YWI=
go-conv b64dec YWI=      # ab
go-conv hexenc hi        # 6869
go-conv hexdec 6869      # hi
```

## 说明

零依赖纯 Go。所有转换逻辑收在 `conv()` 函数里，命令行和测试共用同一份实现。

# go-conv

编解码这种小事，犯不着每次都跑在线工具网站溜一圈。

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

// go-conv 做常见的进制/编码转换：十进制<->十六进制、十进制<->二进制、
// 以及 URL 编码/解码、Base64 编解码、十六进制编解码。子模式用第一个参数指定。
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

// conv 按模式对值做转换，返回结果或错误。所有转换逻辑都收在这，
// 这样既能给命令行用，也能直接被测试调用。
func conv(mode, val string) (string, error) {
	switch mode {
	case "dec2hex":
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%x", n), nil
	case "hex2dec":
		n, err := strconv.ParseInt(val, 16, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", n), nil
	case "dec2bin":
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%b", n), nil
	case "bin2dec":
		n, err := strconv.ParseInt(val, 2, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", n), nil
	case "urlenc":
		return url.QueryEscape(val), nil
	case "urldec":
		return url.QueryUnescape(val)
	case "b64enc":
		return base64.StdEncoding.EncodeToString([]byte(val)), nil
	case "b64dec":
		b, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case "hexenc":
		return hex.EncodeToString([]byte(val)), nil
	case "hexdec":
		b, err := hex.DecodeString(val)
		if err != nil {
			return "", err
		}
		return string(b), nil
	default:
		return "", fmt.Errorf("未知模式: %s", mode)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "用法: go-conv <模式> <值>  模式: dec2hex hex2dec dec2bin bin2dec urlenc urldec b64enc b64dec hexenc hexdec")
		os.Exit(1)
	}
	out, err := conv(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "错误:", err)
		os.Exit(1)
	}
	fmt.Println(out)
}

package logs

import (
	"fmt"
	"strconv"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色

// 3 位前景色, 4 位背景色

// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

// Color defines a single SGR Code
type Color int

func (c Color) String() string {
	return strconv.Itoa(int(c))
}

// Foreground text colors
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func CLogf(f func() string, c ...string) {
	arg := make([]string, 0)
	arg = append(arg, f())
	arg = append(arg, c...)
	CLog(arg...)
}

func CLog(c ...string) {
	var f string
	switch len(c) {
	case 2:
		f = fmt.Sprintf("\033[1;%sm%s\033[0m", c[1], c[0])
	case 3:
		f = fmt.Sprintf("\033[1;%s;%sm%s\033[0m", c[1], c[2], c[0])
	default:
		f = c[0]
	}
	fmt.Println(f)
}

func CDebug(c string) {
	CLog(c, FgGreen.String())
}

func CInfo(c string) {
	CLog(c, FgBlue.String())
}

func CWaring(c string) {
	CLog(c, FgYellow.String())
}

func CError(c string) {
	CLog(c, FgRed.String())
}

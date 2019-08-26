/*
* @Author: scottxiong
* @Date:   2019-08-26 22:48:33
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-26 23:28:29
 */
package color

import (
	"github.com/fatih/color"
)

var (
	red         = color.New(color.FgRed)
	green       = color.New(color.FgGreen)
	yellow      = color.New(color.FgYellow)
	cyan        = color.New(color.FgCyan)
	blue        = color.New(color.FgBlue)
	magenta     = color.New(color.FgMagenta)
	white       = color.New(color.FgWhite)
	BoldRed     = red.Add(color.Bold)
	BoldGreen   = green.Add(color.Bold)
	BoldYellow  = yellow.Add(color.Bold)
	BoldMagenta = magenta.Add(color.Bold)
	BoldBlue    = blue.Add(color.Bold)
	BoldCyan    = cyan.Add(color.Bold)
	BoldWhite   = white.Add(color.Bold)
)

package tags

import (
	"strings"
)

type TextNode struct {
	BaseNode
	line    string
	newline bool
	notrim  bool
}

func CreateTextNode(line string, newline, notrim bool) *TextNode {
	line = strings.Replace(line, "\\&.", ".", -1)
	line = strings.Replace(line, "\\-", "-", -1)
	line = strings.Replace(line, ">", "&gt;", -1)
	line = strings.Replace(line, "<", "&lt;", -1)
	line = strings.Replace(line, "\\\\", "\\", -1)
	if notrim {
		if strings.HasPrefix(line, "    ") {
			line = line[4:]
		}
	} else {
		line = strings.TrimSpace(line)
	}
	line = convertFormatters(line)
	return &TextNode{
		line:    line,
		newline: newline,
	}
}

func (l *TextNode) Execute(out OutputWriter) {
	prefix := ""
	if l.newline {
		prefix = "\n"
	}

	out.Write(prefix+l.line+" ", false)
}

func convertFormatters(line string) string {
	italic := strings.Index(line, "\\fI")
	bold := strings.Index(line, "\\fB")

	if italic == -1 && bold == -1 {
		return line
	} else if italic == -1 {
		line = handleBoldConversion(line)
		return convertFormatters(line)
	} else if bold == -1 {
		line = handleItalicConversion(line)
		return convertFormatters(line)
	} else if bold < italic {
		line = handleBoldConversion(line)
		return convertFormatters(line)
	} else {
		line = handleItalicConversion(line)
		return convertFormatters(line)
	}
}

func handleBoldConversion(line string) string {
	line = strings.Replace(line, "\\fB", "<xref href=\"", 1)
	idx := strings.Index(line, "\\fR")
	end := "\\fR"
	if idx > 0 && (idx+5) < len(line) {
		if string(line[idx+3]) == "(" && string(line[idx+5]) == ")" {
			end = line[idx : idx+6]
		}
	}
	line = strings.Replace(line, end, ".dita\"/>", 1)
	return line
}

func handleItalicConversion(line string) string {
	line = strings.Replace(line, "\\fI", "<i>", 1)
	line = strings.Replace(line, "\\fR", "</i>", 1)
	return line
}

package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*

type LineToken struct {
	BaseNode
	line string
}*/

type Token string

const (
	MAN_MACRO       = "'\\\"" // Denotes that this is a man page
	COMMENT         = ".\\\"" // Comment line
	PARAGRAPH       = ".PP"   // A paragraph tag
	SECTION         = ".SH"   // Section header
	TITLE           = ".TH"   // Skipped since we inject our own title
	IFELSE          = ".ie"   // Macro for man portability
	ELSE            = ".el"   // Macro for man portability
	HYPH            = ".nh"   // Macro to disable man hyphenation
	ADJUST          = ".ad"   // Macro to disable man justification
	LINE_BREAK      = ".sp"   // Line breaks will be ignored
	BULLET_LIST     = ".IP"   // Index paragraph, used for lists
	INDENT          = ".RS"   // Index margin
	UNINDENT        = ".RE"   // Unindent marging
	START_CODEBLOCK = ".DS"   // Start of codeblock
	END_CODEBLOCK   = ".DE"   // End of codeblock
	SUBSECTION      = ".SS"   // Subsection header
	TEXT            = "text"  // A text line
	END             = "eof"   // End of file
)

type Tokenizer struct {
	infile   *os.File
	scanner  *bufio.Scanner
	curTok   Token
	curArg   string
	finished bool
}

func CreateTokenizer(infile string) (*Tokenizer, error) {
	in, err := os.OpenFile(infile, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	ret := &Tokenizer{
		infile:   in,
		scanner:  scanner,
		finished: false,
	}

	// Read in the first token
	ret.Next()
	return ret, nil
}

func (t *Tokenizer) GetName() string {
	_, name := filepath.Split(t.infile.Name())
	return name
}

func (t *Tokenizer) Peek() (Token, string) {
	return t.curTok, t.curArg
}

func (t *Tokenizer) Next() bool {
	if t.finished {
		return false
	}

	if t.scanner.Scan() {
		line := t.scanner.Text()

		if !strings.HasPrefix(line, ".") && !strings.HasPrefix(line, MAN_MACRO) {
			t.curTok = TEXT
			t.curArg = line
			return true
		}

		if len(line) < 3 {
			return t.Next()
		}

		tag := line[0:3]
		if tag == ADJUST || tag == COMMENT || tag == ELSE || tag == HYPH ||
			tag == IFELSE || tag == TITLE || tag == MAN_MACRO ||
			tag == LINE_BREAK {
			return t.Next()
		} else if tag == BULLET_LIST || tag == INDENT || tag == UNINDENT ||
			tag == SECTION || tag == SUBSECTION || tag == PARAGRAPH ||
			tag == START_CODEBLOCK || tag == END_CODEBLOCK {
			t.curTok = Token(tag)
			t.curArg = line[3:]
		} else {
			fmt.Printf("Unexpected line: %s\n", line)
			os.Exit(1)
		}
		return true
	} else {
		t.curTok = END
		t.curArg = ""
		t.finished = true
		return false
	}
}

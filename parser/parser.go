package parser

import (
	"fmt"

	"github.com/couchbaselabs/man2dita/tags"
	"github.com/couchbaselabs/man2dita/tokenizer"
)

type Parser struct {
	halt   bool
	tokens *tokenizer.Tokenizer
}

func CreateParser(tokens *tokenizer.Tokenizer) *Parser {
	return &Parser{
		halt:   false,
		tokens: tokens,
	}
}

func (p *Parser) Parse(name string) tags.AbstractNode {
	// Add XML Root
	root := tags.CreateXMLNode()

	// Add the topic tag
	topic := tags.CreateTopicNode(p.tokens.GetName())
	root.Add(topic)

	// Add the page title tag
	title := tags.CreateTitleNode(name)
	topic.Add(title)

	// Start the document body
	body := tags.CreateBodyNode()
	p.parseBody(body)
	topic.Add(body)

	return root
}

func (p *Parser) parseBody(body tags.AbstractNode) {
	token, arg := p.tokens.Peek()
	switch token {
	case tokenizer.SECTION:
		section := tags.CreateSectionNode(arg)
		p.tokens.Next()
		p.parseSection(section)
		body.Add(section)
		p.parseBody(body)
	case tokenizer.END:
		break
	default:
		fmt.Printf("Unexpect line in body: %s\n", token)
		fmt.Printf("\tLine: %s%s\n", token, arg)
	}
}

func (p *Parser) parseSection(section tags.AbstractNode) {
	token, arg := p.tokens.Peek()
	switch token {
	case tokenizer.PARAGRAPH:
		paragraph := tags.CreateParagraphNode()
		p.tokens.Next()
		p.parseParagraph(paragraph)
		section.Add(paragraph)
		p.parseSection(section)
	case tokenizer.BULLET_LIST:
		list := tags.CreateListNode()
		p.tokens.Next()
		p.parseList(list)
		section.Add(list)
		p.parseSection(section)
	case tokenizer.SUBSECTION:
		subsection := tags.CreateSubsectionNode(arg)
		p.tokens.Next()
		p.parseSubsection(subsection)
		section.Add(subsection)
		p.parseSection(section)
	case tokenizer.START_CODEBLOCK:
		codeblock := tags.CreateCodeblockNode()
		p.tokens.Next()
		p.parseCodeblock(codeblock, false)
		section.Add(codeblock)
		p.parseSection(section)
	case tokenizer.INDENT:
		indent := tags.CreateIndentNode()
		p.tokens.Next()
		p.parseIndent(indent, false)
		section.Add(indent)
		p.parseSubsection(section)
	default:
		break
	}
}

func (p *Parser) parseSubsection(subsection tags.AbstractNode) {
	token, _ := p.tokens.Peek()
	switch token {
	case tokenizer.PARAGRAPH:
		paragraph := tags.CreateParagraphNode()
		p.tokens.Next()
		p.parseParagraph(paragraph)
		subsection.Add(paragraph)
		p.parseSubsection(subsection)
	case tokenizer.BULLET_LIST:
		list := tags.CreateListNode()
		p.tokens.Next()
		p.parseList(list)
		subsection.Add(list)
		p.parseSubsection(subsection)
	case tokenizer.START_CODEBLOCK:
		codeblock := tags.CreateCodeblockNode()
		p.tokens.Next()
		p.parseCodeblock(codeblock, false)
		subsection.Add(codeblock)
		p.parseSubsection(subsection)
	case tokenizer.INDENT:
		indent := tags.CreateIndentNode()
		p.tokens.Next()
		p.parseIndent(indent, false)
		subsection.Add(indent)
		p.parseSubsection(subsection)
	default:
		break
	}
}

func (p *Parser) parseIndent(indent tags.AbstractNode, firstCall bool) {
	token, args := p.tokens.Peek()
	switch token {
	case tokenizer.TEXT:
		indent.Add(tags.CreateTextNode(args, firstCall, false))
		p.tokens.Next()
		p.parseIndent(indent, true)
	case tokenizer.UNINDENT:
		p.tokens.Next()
		break
	default:
		break
	}
}

func (p *Parser) parseCodeblock(codeblock tags.AbstractNode, firstCall bool) {
	token, args := p.tokens.Peek()
	switch token {
	case tokenizer.TEXT:
		codeblock.Add(tags.CreateTextNode(args, firstCall, true))
		p.tokens.Next()
		p.parseCodeblock(codeblock, true)
	case tokenizer.END_CODEBLOCK:
		p.tokens.Next()
		break
	default:
		break
	}
}

func (p *Parser) parseParagraph(paragraph tags.AbstractNode) {
	token, args := p.tokens.Peek()
	switch token {
	case tokenizer.TEXT:
		paragraph.Add(tags.CreateTextNode(args, false, false))
		p.tokens.Next()
		p.parseParagraph(paragraph)
	default:
		break
	}
}

func (p *Parser) parseList(list tags.AbstractNode) {
	token, args := p.tokens.Peek()
	switch token {
	case tokenizer.TEXT:
		list.Add(tags.CreateTextNode(args, false, false))
		p.tokens.Next()
		p.parseList(list)
	default:
		break
	}
}

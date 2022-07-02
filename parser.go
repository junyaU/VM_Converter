package VM_Converter

import (
	"bufio"
	"io"
	"strings"
)

type Parser struct {
	commands    []string
	currentLine int
}

func NewParser(f io.Reader) *Parser {
	scanner := bufio.NewScanner(f)
	var rowTexts []string

	for scanner.Scan() {
		text := scanner.Text()

		commentOutIndex := strings.Index(text, "//")
		if commentOutIndex != -1 {
			text = text[:commentOutIndex]
		}

		if text == "" {
			continue
		}

		rowTexts = append(rowTexts, strings.TrimSpace(text))
	}

	return &Parser{
		commands:    rowTexts,
		currentLine: 0,
	}
}

func (p Parser) HasMoreCommands() bool {
	return len(p.commands) > p.currentLine
}

func (p *Parser) Advance() {
	p.currentLine++
}

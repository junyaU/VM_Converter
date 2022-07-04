package VM_Converter

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Parser struct {
	texts       []string
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
		texts:       rowTexts,
		currentLine: 0,
	}
}

func (p Parser) Commands() []string {
	return p.texts
}

func (p Parser) HasMoreCommands() bool {
	return len(p.texts) > p.currentLine
}

func (p *Parser) Advance() {
	p.currentLine++
}

func (p Parser) CommandType() (VMCommand, error) {
	command := strings.Split(p.texts[p.currentLine], " ")[0]

	switch command {
	case "push":
		return C_PUSH, nil
	case "pop":
		return C_POP, nil
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ARITHMETIC, nil
	default:
		return 0, errors.New("このコマンドは存在しない")
	}
}

func (p Parser) Arg1() (string, error) {
	commandType, err := p.CommandType()
	if err != nil {
		return "", err
	}

	lineTexts := strings.Split(p.texts[p.currentLine], " ")

	switch commandType {
	case C_ARITHMETIC:
		return lineTexts[0], nil
	default:
		return lineTexts[1], nil
	}
}

func (p Parser) Arg2() (string, error) {
	commandType, err := p.CommandType()
	if err != nil {
		return "", err
	}

	switch commandType {
	case C_POP, C_PUSH, C_FUNCTION, C_CALL:
		return strings.Split(p.texts[p.currentLine], " ")[2], nil
	default:
		return "", errors.New("このコマンドタイプにはたいおうしておらず")
	}
}

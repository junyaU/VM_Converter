package VM_Converter

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Parser struct {
	texts          []string
	currentLine    int
	isExistSysFile bool
}

func NewParser(fs []*os.File) *Parser {
	var rowTexts []string

	var isSys bool
	if len(fs) > 1 {
		isSys = true
	}

	for _, f := range fs {
		scanner := bufio.NewScanner(f)

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

	}
	return &Parser{
		texts:          rowTexts,
		currentLine:    0,
		isExistSysFile: isSys,
	}
}

func (p Parser) Commands() []string {
	return p.texts
}

func (p Parser) IsExistSys() bool {
	return p.isExistSysFile
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
	case "label":
		return C_LABEL, nil
	case "if-goto":
		return C_IF, nil
	case "goto":
		return C_GOTO, nil
	case "function":
		return C_FUNCTION, nil
	case "return":
		return C_RETURN, nil
	case "call":
		return C_CALL, nil
	default:
		return 0, errors.New("this command does not exist")
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
		return "", errors.New("this command does not exist")
	}
}

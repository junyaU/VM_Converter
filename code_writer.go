package VM_Converter

import (
	"errors"
	"os"
	"strings"
)

type CodeWriter struct {
	file *os.File
}

func NewCodeWriter(f *os.File) *CodeWriter {

	return &CodeWriter{
		file: f,
	}
}

func (w *CodeWriter) WriteArithmetic(command string) error {
	switch command {
	case "add":
		_, err := w.file.WriteString("@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nM=M+D\n@SP\nM=M+1")
		return err
	default:
		return errors.New("this command does not exist")
	}
}

func (w *CodeWriter) WritePush(segment string, index string) error {
	var textToWrite strings.Builder

	switch segment {
	case "constant":
		textToWrite.WriteString("@")
		textToWrite.WriteString(index)
		textToWrite.WriteString("\n")
		textToWrite.WriteString("D=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n")

		_, err := w.file.WriteString(textToWrite.String())
		return err
	default:
		return errors.New("this segment does not exist")
	}
}

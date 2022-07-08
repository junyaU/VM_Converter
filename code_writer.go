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
	var asmText string
	switch command {
	case "add":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nM=M+D\n@SP\nM=M+\n"
	case "sub":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nM=M-D\n@SP\nM=M+\n"
	case "neg":
		asmText = "@SP\nAM=M-1\nM=-M\n@SP\nM=M+1\n"
	case "eq":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nD=M-D\n@COMP.0.TRUE\nD;JEQ\n@COMP.0.FALSE\n0;JMP\n(COMP.0.TRUE)\n@SP\nA=M\nM=-1\n@SP\nM=M+1\n@COMP.0.END\n0;JMP\n(COMP.0.FALSE)\n@SP\nA=M\nM=0\n@SP\nM=M+1\n(COMP.0.END)\n"
	case "gt":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nD=M-D\n@COMP.6.TRUE\nD;JGT\n@COMP.6.FALSE\n0;JMP\n(COMP.6.TRUE)\n@SP\nA=M\nM=-1\n@SP\nM=M+1\n@COMP.6.END\n0;JMP\n(COMP.6.FALSE)\n@SP\nA=M\nM=0\n@SP\nM=M+1\n(COMP.6.END)\n"
	case "lt":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nD=M-D\n@COMP.3.TRUE\nD;JLT\n@COMP.3.FALSE\n0;JMP\n(COMP.3.TRUE)\n@SP\nA=M\nM=-1\n@SP\nM=M+1\n@COMP.3.END\n0;JMP\n(COMP.3.FALSE)\n@SP\nA=M\nM=0\n@SP\nM=M+1\n(COMP.3.END)\n"
	case "and":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nM=D&M\n@SP\nM=M+\n"
	case "or":
		asmText = "@SP\nAM=M-1\nD=M\n@SP\nAM=M-1\nM=D|M\n@SP\nM=M+\n"
	case "not":
		asmText = "@SP\nAM=M-1\nM=!M\n@SP\nM=M+1\n"
	default:
		return errors.New("this command does not exist")
	}

	_, err := w.file.WriteString(asmText)
	return err
}

func (w *CodeWriter) WritePush(segment, index string) error {
	textToWrite := func(label, index string) string {
		var t strings.Builder

		t.WriteString("@" + index + "\n")
		t.WriteString("D=A\n@" + label + "\nA=M\nD=D+A\nA=D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n")
		return t.String()
	}

	var label string
	switch segment {
	case "local":
		label = "LCL"
	case "constant":
		var t strings.Builder
		t.WriteString("@")
		t.WriteString(index)
		t.WriteString("\n")
		t.WriteString("D=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n")

		_, err := w.file.WriteString(t.String())
		return err
	default:
		return errors.New("this segment does not exist")
	}

	_, err := w.file.WriteString(textToWrite(label, index))
	return err
}

func (w *CodeWriter) WritePop(segment, index string) error {
	textToWrite := func(label, index string) string {
		var text strings.Builder
		indexARegister := "@" + index + "\n"

		text.WriteString(indexARegister)
		text.WriteString("D=A\n@" + label + "\nA=M\nD=D+A\n@" + label + "\nM=D\n@SP\nM=M-1\nA=M\nD=M\n@" + label + "\nA=M\nM=D\n")
		text.WriteString(indexARegister)
		text.WriteString("D=A\n" + label + "A=M\nD=A-D\n" + label + "M=D\n")

		return text.String()
	}

	var label string
	switch segment {
	case "local":
		label = "LCL"
	case "argument":
		label = "ARG"
	case "this":
		label = "THIS"
	case "that":
		label = "THAT"
	default:
		return errors.New("this segment does not exist")
	}

	_, err := w.file.WriteString(textToWrite(label, index))
	return err
}

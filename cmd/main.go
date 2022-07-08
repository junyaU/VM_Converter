package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	VM_Converter "vm_converter"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("specify asm file as an argument")
		os.Exit(1)
	}

	fileName := flag.Args()[0]
	dataDir := "testdata/"

	f, err := os.Open(dataDir + fileName)
	if err != nil {
		fmt.Println("the specified file does not exist")
		os.Exit(1)
	}

	defer f.Close()

	p := VM_Converter.NewParser(f)

	isVmFile := strings.Contains(fileName, ".vm")
	if !isVmFile {
		fmt.Println("cannot specify anything other than a vm file")
		os.Exit(1)
	}

	vmExtensionIndex := strings.Index(fileName, ".vm")
	asmFile, err := os.Create(dataDir + fileName[:vmExtensionIndex] + ".asm")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer asmFile.Close()

	writer := VM_Converter.NewCodeWriter(asmFile)

	if err := WriteToAsmFile(writer, p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("success")
}

func WriteToAsmFile(w *VM_Converter.CodeWriter, p *VM_Converter.Parser) error {
	for i := 0; i < len(p.Commands()); i++ {
		if !p.HasMoreCommands() {
			break
		}

		commandType, err := p.CommandType()
		if err != nil {
			return err
		}

		switch commandType {
		case VM_Converter.C_PUSH:
			seg, _ := p.Arg1()
			index, _ := p.Arg2()
			if err := w.WritePush(seg, index); err != nil {
				return err
			}

		case VM_Converter.C_POP:
			seg, _ := p.Arg1()
			index, _ := p.Arg2()

			if err := w.WritePop(seg, index); err != nil {
				return err
			}

		case VM_Converter.C_ARITHMETIC:
			cmd, _ := p.Arg1()
			if err := w.WriteArithmetic(cmd); err != nil {
				return err
			}
		}

		p.Advance()
	}

	return nil
}

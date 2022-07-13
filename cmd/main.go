package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	VM_Converter "vm_converter"
)

const DATA_DIR = "testdata/"

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("specify asm file as an argument")
		os.Exit(1)
	}

	filePath := flag.Args()[0]

	files, err := GetFiles(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := VM_Converter.NewParser(files)

	asmFile, err := CreateAsmFile(filePath)
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

	for _, f := range files {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("success")
}

func GetFiles(filePath string) ([]*os.File, error) {
	fileInfo, err := os.Stat(DATA_DIR + filePath)
	if err != nil {
		return nil, err
	}

	var fs []*os.File
	if fileInfo.IsDir() {
		files, _ := ioutil.ReadDir(DATA_DIR + filePath)

		for _, file := range files {
			f, err := os.Open(DATA_DIR + filePath + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			fs = append(fs, f)
		}
	} else {
		f, err := os.Open(DATA_DIR + filePath)
		if err != nil {
			return nil, err
		}

		fs = append(fs, f)
	}

	return fs, nil
}

func CreateAsmFile(filePath string) (*os.File, error) {
	var asmFileName string
	isVmFile := strings.Contains(filePath, ".vm")
	if isVmFile {
		asmFileName = DATA_DIR + filePath[:strings.Index(filePath, ".vm")] + ".asm"
	} else {
		asmFileName = DATA_DIR + filePath + "/" + filePath + ".asm"
	}

	asmFile, err := os.Create(asmFileName)
	if err != nil {
		return nil, err
	}

	return asmFile, nil
}

func WriteToAsmFile(w *VM_Converter.CodeWriter, p *VM_Converter.Parser) error {
	if p.IsExistSys() {
		if err := w.WriteInit(); err != nil {
			return err
		}
	}

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

		case VM_Converter.C_LABEL:
			cmd, _ := p.Arg1()
			if err := w.WriteLabel(cmd); err != nil {
				return err
			}

		case VM_Converter.C_IF:
			cmd, _ := p.Arg1()
			if err := w.WriteIf(cmd); err != nil {
				return err
			}

		case VM_Converter.C_GOTO:
			cmd, _ := p.Arg1()
			if err := w.WriteGoto(cmd); err != nil {
				return err
			}

		case VM_Converter.C_FUNCTION:
			cmd, _ := p.Arg1()
			numLocals, _ := p.Arg2()
			n, _ := strconv.Atoi(numLocals)
			if err := w.WriteFunction(cmd, n); err != nil {
				return err
			}

		case VM_Converter.C_RETURN:
			if err := w.WriteReturn(); err != nil {
				return err
			}

		case VM_Converter.C_CALL:
			cmd, _ := p.Arg1()
			numLocals, _ := p.Arg2()
			n, _ := strconv.Atoi(numLocals)
			if err := w.WriteCall(cmd, n); err != nil {
				return err
			}
		}

		p.Advance()
	}

	return nil
}

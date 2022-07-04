package VM_Converter

type VMCommand int

const (
	C_ARITHMETIC VMCommand = iota + 1
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

func (t VMCommand) String() string {
	v := [...]string{"C_ARITHMETIC", "C_PUSH", "C_POP", "C_LABEL", "C_GOTO", "C_IF", "C_FUNCTION", "C_RETURN", "C_CALL"}
	return v[t-1]
}

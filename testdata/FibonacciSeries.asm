@1
D=A
@ARG
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@THAT
M=D
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@THAT
A=M
D=D+A
@THAT
M=D
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@0
D=A
THAT
A=M
D=A-D
THAT
M=D
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@THAT
A=M
D=D+A
@THAT
M=D
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@1
D=A
THAT
A=M
D=A-D
THAT
M=D
@0
D=A
@ARG
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
@SP
AM=M-1
M=M-D
@SP
M=M+1
@0
D=A
@ARG
A=M
D=D+A
@ARG
M=D
@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D
@0
D=A
ARG
A=M
D=A-D
ARG
M=D
(MAIN_LOOP_START)
@0
D=A
@ARG
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@COMPUTE_ELEMENT
D;JNE
@END_PROGRAM
0;JMP
(COMPUTE_ELEMENT)
@0
D=A
@THAT
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@THAT
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
@SP
AM=M-1
M=M+D
@SP
M=M+1
@2
D=A
@THAT
A=M
D=D+A
@THAT
M=D
@SP
M=M-1
A=M
D=M
@THAT
A=M
M=D
@2
D=A
THAT
A=M
D=A-D
THAT
M=D
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
@SP
AM=M-1
M=M+D
@SP
M=M+1
@SP
M=M-1
A=M
D=M
@THAT
M=D
@0
D=A
@ARG
A=M
D=D+A
A=D
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
@SP
AM=M-1
M=M-D
@SP
M=M+1
@0
D=A
@ARG
A=M
D=D+A
@ARG
M=D
@SP
M=M-1
A=M
D=M
@ARG
A=M
M=D
@0
D=A
ARG
A=M
D=A-D
ARG
M=D
@MAIN_LOOP_START
0;JMP
(END_PROGRAM)
# AntlrPractice

> This repo is a demo for studying parser and programming language.

## CSPec

grammar file source: [antlr/grammars-v4](https://github.com/antlr/grammars-v4/blob/master/c/C.g4)

Initial aim is to study array of pointer and pointer of array by the operator precedence and associativity.

[doc for c++ operator precedence and associativity](https://docs.microsoft.com/zh-cn/cpp/cpp/cpp-built-in-operators-precedence-and-associativity?view=msvc-160)

Read the parse tree of code to study.

```c
int * a[10];
int (* a)[10];
```
# Monkey Interpreter (Go)

An implementation of the **Monkey** programming language in Go, based
on\
*Writing an Interpreter in Go* by Thorsten Ball.

The goal of this project is to deeply understand how an interpreter
works by building one from scratch: lexing, parsing, AST construction,
and evaluation.

------------------------------------------------------------------------

## 📌 Project Status

**Current Progress: Chapter 1 Completed (Lexer)**

✔ Fully implemented Lexer\
✔ Token and keyword definitions\
✔ Support for operators:

-   `=`
-   `+ - * /`
-   `< >`
-   `== !=`
-   `!`
-   `, ; ( ) { }`

✔ Support for: - Identifiers - Integer literals - Keywords: `let`, `fn`,
`if`, `else`, `return`, `true`, `false`

🧪 Lexer tests implemented and passing.

------------------------------------------------------------------------

## 📂 Project Structure

    monkey/
    ├── lexer/        # Lexical analysis
    ├── token/        # Token definitions and keywords
    ├── go.mod
    └── README.md

------------------------------------------------------------------------

## 🚀 Running Tests

Run lexer tests:

``` bash
go test ./lexer
```

Run all project tests:

``` bash
go test ./...
```

------------------------------------------------------------------------

## 🧪 Example Monkey Code Supported by the Lexer

``` monkey
let add = fn(x, y) {
  x + y;
};

if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
```

------------------------------------------------------------------------

## 🎯 Next Step

Chapter 2 --- Implementing the Parser and building the AST.

------------------------------------------------------------------------

## 📚 Reference

Thorsten Ball --- *Writing an Interpreter in Go*\
https://interpreterbook.com/

------------------------------------------------------------------------

## 🛠 Requirements

-   Go 1.20+

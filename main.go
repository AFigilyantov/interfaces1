// Message Formatter
// Реализуйте интерфейс Formatter с методом Format, который возвращает
// отформатированную строку.
// Определите структуры, удовлетворяющие интерфейсу Formatter: обычный
// текст(как есть), жирным шрифтом(** **), код(` `), курсив(_ _)
// Опционально: иметь возможность задавать цепочку модификаторов
// chainFormatter.AddFormatter(plainText)
// chainFormatter.AddFormatter(bold)
// chainFormatter.AddFormatter(code)

package main

import (
	"fmt"
	"log"
)

func TextFormating(f Formatter, text string) {
	formatedText, err := f.Format(text)
	if err != nil {
		log.Printf("Formatting is failed with error %s", err)
		return
	}
	fmt.Println(formatedText)
}

func main() {
	var format Formatter

	format = PlainText{}

	TextFormating(format, "plain text")

	format = BoldText{}

	TextFormating(format, "bold text")

	format = ItalicText{}

	TextFormating(format, "italic text")

	format = CodeText{}

	TextFormating(format, "code text")

}

type Formatter interface {
	Format(string) (string, error)
}

type PlainText struct {
}

func (p PlainText) Format(text string) (string, error) {
	return fmt.Sprintf(" %s ", text), nil
}

type BoldText struct {
}

func (b BoldText) Format(text string) (string, error) {
	return fmt.Sprintf("** %s **", text), nil
}

type CodeText struct {
}

func (c CodeText) Format(text string) (string, error) {
	return fmt.Sprintf("` %s `", text), nil
}

type ItalicText struct {
}

func (i ItalicText) Format(text string) (string, error) {
	return fmt.Sprintf("_ %s _", text), nil
}

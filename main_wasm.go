package main

import (
    "fmt"
    "syscall/js"
    "naknak/naknak"
)

func encode(this js.Value, inputs []js.Value) any {
	if len(inputs) != 1 {
        return fmt.Sprintf("Expected one argument, got `%+v` arguments", inputs)
    }
    if inputs[0].Type() != js.TypeString     {
        return "Expected a string argument"
    }
    text := inputs[0].String()
    return naknak.Encode(text)
}

func decode(this js.Value, inputs []js.Value) any {
	if len(inputs) != 1 {
        return fmt.Sprintf("Expected one argument, got `%+v` arguments", inputs)
    }
    if inputs[0].Type() != js.TypeString     {
        return "Expected a string argument"
    }
    text := inputs[0].String()
    return naknak.Decode(text)
}

func main() {
    c := make(chan struct{}, 0)
    js.Global().Set("encode", js.FuncOf(encode))
    js.Global().Set("decode", js.FuncOf(decode))
    <-c
}

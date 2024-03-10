package naknak

import (
    "fmt"
    "slices"
    "strings"
)

var nakNak = []string{
    "Nak",
    "Nanak",
    "Nananak",
    "Nanananak",
    "Nak?",
    "nak?",
    "Naknak",
    "Naknaknak",
    "Nak.",
    "Naknak.",
    "Naknaknaknak",
    "nanak",
    "naknak",
    "nak!",
    "nak.",
    "naknaknak",
}

func Encode(text string) string {
    sb := strings.Builder{}
    for _, r := range text {
        y := int(r) % 16
        x := int(r) / 16
        if x < 0 || x >= len(nakNak) {
            continue
        }
        if sb.Len() > 0 {
            sb.WriteString(" ")
        }
        sb.WriteString(nakNak[x])
        sb.WriteString(" ")
        sb.WriteString(nakNak[y])
    }
    return sb.String()
}

func Decode(encodedText string) string {
    sb := strings.Builder{}
    words := strings.Split(encodedText, " ")
    for i := 0; i < len(words)-1; {
        x := slices.Index(nakNak, words[i])
        y := slices.Index(nakNak, words[i+1])
        if x != -1 && y != -1 {
            ord := x*16 + y
            sb.WriteString(fmt.Sprintf("%c", ord))
            i += 2
        } else {
            i += 1
        }
    }
    return sb.String()
}

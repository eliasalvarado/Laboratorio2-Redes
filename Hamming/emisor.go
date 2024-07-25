package main

import (
	"fmt"
	"strings"
)


func convertirStringABits(mensaje string) string {
    var bits strings.Builder
    for _, c := range mensaje {
        bits.WriteString(fmt.Sprintf("%08b", c))
    }
    return bits.String()
}

func agregarBitsDeParidad(bits string) string {
    if len(bits) != 4 {
        panic("La entrada debe ser de 4 bits")
    }

    d := make([]int, 4)
    for i := 0; i < 4; i++ {
        d[i] = int(bits[i] - '0')
    }

    p1 := (d[0] + d[1] + d[3]) % 2
    p2 := (d[0] + d[2] + d[3]) % 2
    p3 := (d[1] + d[2] + d[3]) % 2

    hamming := fmt.Sprintf("%d%d%d%d%d%d%d", p1, p2, d[0], p3, d[1], d[2], d[3])
    return hamming
}

func convertirBitsAHamming(bits string) string {
    var hamming strings.Builder
    for i := 0; i < len(bits); i += 4 {
        if i+4 <= len(bits) {
            bloque := bits[i : i+4]
            hamming.WriteString(agregarBitsDeParidad(bloque))
        }
    }
    return hamming.String()
}

func main() {
    mensaje := "Hola"
    bits := convertirStringABits(mensaje)
    hamming := convertirBitsAHamming(bits)
    fmt.Println("Mensaje original:", mensaje)
    fmt.Println("Bits:", bits)
    fmt.Println("Código Hamming:", hamming)
}
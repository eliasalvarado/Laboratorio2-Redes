package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// agregarBitsDeParidad toma un bloque de 4 bits y agrega 3 bits de paridad para formar un bloque de 7 bits (código Hamming)
func agregarBitsDeParidad(bits string) string {
	if len(bits) != 4 {
		panic("La entrada debe ser de 4 bits")
	}

	d := make([]int, 4)
	for i := 0; i < 4; i++ {
		d[i] = int(bits[i] - '0')
	}

	// Calcular los bits de paridad
	p1 := (d[0] + d[1] + d[3]) % 2
	p2 := (d[0] + d[2] + d[3]) % 2
	p3 := (d[1] + d[2] + d[3]) % 2

	// El bloque de 7 bits es de la forma p1, p2, d0, p3, d1, d2, d3
	hamming := fmt.Sprintf("%d%d%d%d%d%d%d", p1, p2, d[0], p3, d[1], d[2], d[3])
	return hamming
}

// convertirBitsAHamming toma una cadena de bits y la convierte en una cadena de códigos Hamming
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el mensaje en cadena de bits:")
	inputMessage, _ := reader.ReadString('\n')
	inputMessage = strings.TrimSpace(inputMessage)

	hamming := convertirBitsAHamming(inputMessage)
	fmt.Println("Código Hamming:", hamming)
}

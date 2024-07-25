package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func agregarBitsDeParidad(bits string, m, n int) string {
	numParidades := n - m

	hamming := make([]rune, n)
	j := 0
	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			hamming[i-1] = 'P'
		} else {
			hamming[i-1] = rune(bits[j])
			j++
		}
	}

	for i := 0; i < numParidades; i++ {
		pos := 1 << i
		val := 0
		for j := pos; j <= n; j += pos * 2 {
			for k := 0; k < pos && j+k <= n; k++ {
				if hamming[j+k-1] == '1' {
					val ^= 1
				}
			}
		}
		if val == 1 {
			hamming[pos-1] = '1'
		} else {
			hamming[pos-1] = '0'
		}
	}

	return string(hamming)
}

func convertirBitsAHamming(bits string, m, n int) string {
	var hamming strings.Builder
	for i := 0; i < len(bits); i += m {
		if i+m <= len(bits) {
			bloque := bits[i : i+m]
			hamming.WriteString(agregarBitsDeParidad(bloque, m, n))
		}
	}
	return hamming.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Leer n y m
	fmt.Println("Ingrese la combinación (n, m) para el código de Hamming:")
	nmInput, _ := reader.ReadString('\n')
	nmInput = strings.TrimSpace(nmInput)
	nm := strings.Split(nmInput, ",")
	n, errN := strconv.Atoi(nm[0])
	m, errM := strconv.Atoi(nm[1])

	if errN != nil || errM != nil || n <= m {
		fmt.Println("Combinación (n, m) inválida.")
		return
	}

	fmt.Println("Ingrese el mensaje en cadena de bits:")
	inputMessage, _ := reader.ReadString('\n')
	inputMessage = strings.TrimSpace(inputMessage)

	hamming := convertirBitsAHamming(inputMessage, m, n)
	fmt.Println("Código Hamming:", hamming)
}

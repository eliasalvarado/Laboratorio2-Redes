package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
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

func stringToBinary(s string) string {
	var binary strings.Builder
	for _, c := range s {
		binary.WriteString(fmt.Sprintf("%08b", c))
	}
	return binary.String()
}

func aplicarRuido(bits string, prob float64) string {
	rand.Seed(time.Now().UnixNano())
	var noisyBits strings.Builder
	for _, bit := range bits {
		if rand.Float64() < prob {
			if bit == '0' {
				noisyBits.WriteRune('1')
			} else {
				noisyBits.WriteRune('0')
			}
		} else {
			noisyBits.WriteRune(bit)
		}
	}
	return noisyBits.String()
}

func enviarMensajeHamming(host string, port int, mensaje string) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(mensaje))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Capa 1 - Aplicacion

	fmt.Println("Ingrese el mensaje de texto:")
	inputMessage, _ := reader.ReadString('\n')
	inputMessage = strings.TrimSpace(inputMessage)

	binaryMessage := stringToBinary(inputMessage)
	hamming := convertirBitsAHamming(binaryMessage, 4, 7)

	fmt.Println("Ingrese la probabilidad p de ruido (1/p):")
	probInput, _ := reader.ReadString('\n')
	probInput = strings.TrimSpace(probInput)
	prob, errP := strconv.ParseFloat(probInput, 64)

	if errP != nil || prob <= 0 {
		fmt.Println("Probabilidad inválida.")
		return
	}

	

	// Capa 2 - Presentacion

	fmt.Printf("Mensaje codificado a Hamming: %s\n",hamming)
	
	// Ruido
	noisyHamming := aplicarRuido(hamming, 1/prob)
	fmt.Printf("Mensaje codificado con Ruido: %s\n",noisyHamming)


	// Capa 3 - Enlace

	fmt.Println("Ingrese la dirección IP del servidor:")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Println("Ingrese el puerto del servidor:")
	portInput, _ := reader.ReadString('\n')
	portInput = strings.TrimSpace(portInput)
	port, errPort := strconv.Atoi(portInput)

	if errPort != nil {
		fmt.Println("Puerto inválido.")
		return
	}
	
	// Capa 4 - Transmision
	
	err := enviarMensajeHamming(host, port, noisyHamming+"\n")
	if err != nil {
		fmt.Printf("Error al enviar el mensaje: %v\n", err)
	} else {
		fmt.Println("Mensaje Hamming con ruido enviado exitosamente.")
	}
}

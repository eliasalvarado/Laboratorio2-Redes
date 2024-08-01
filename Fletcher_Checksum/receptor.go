package main

import (
	"fmt"
	"strconv"
	"bufio"
	"strings"
	"net"
	"os"
)

func fletcher16(data []byte) (uint16, uint16) {
	var sum1, sum2 uint16
	sum1 = 0
	sum2 = 0

	for _, b := range data {
		sum1 = (sum1 + uint16(b)) % 255
		sum2 = (sum2 + sum1) % 255
	}

	return sum1, sum2
}

func verifyFletcher16Checksum(data []byte, checksum1, checksum2 uint16) bool {
	sum1, sum2 := fletcher16(data)
	return sum1 == checksum1 && sum2 == checksum2
}

func bitsStringToBytes(bitsStr string) ([]byte, error) {
	var data []byte

	for i := 0; i < len(bitsStr); i += 8 {
		bitsSegment := bitsStr[i:min(i+8, len(bitsStr))]

		byteValue, err := strconv.ParseUint(bitsSegment, 2, 8)
		if err != nil {
			return nil, err
		}
		data = append(data, byte(byteValue))
	}

	return data, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bytesToASCIIString(data []byte) string {
	return string(data)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := bufio.NewReader(conn)
	inputMessage, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el mensaje:", err, inputMessage)
		return
	}
	inputMessage = inputMessage[:len(inputMessage)-1]

	if len(inputMessage) < 16 {
		fmt.Println("Mensaje demasiado corto para contener checksum")
		return
	}
	bitsMessage := inputMessage[:len(inputMessage)-16]
	bitsChecksum := inputMessage[len(inputMessage)-16:]

	message, err := bitsStringToBytes(bitsMessage)
	if err != nil {
		fmt.Println("Error al convertir el mensaje:", err)
		return
	}

	checksum1, err := strconv.ParseUint(bitsChecksum[8:], 2, 8)
	if err != nil {
		fmt.Println("Error al convertir el checksum1:", err)
		return
	}
	checksum2, err := strconv.ParseUint(bitsChecksum[:8], 2, 8)
	if err != nil {
		fmt.Println("Error al convertir el checksum2:", err)
		return
	}

	isValid := verifyFletcher16Checksum(message, uint16(checksum1), uint16(checksum2))

	if isValid {
		asciiMessage := bytesToASCIIString(message)
		fmt.Printf("El mensaje fue aceptado: %s\n", asciiMessage)
	} else {
		fmt.Println("El mensaje contiene errores: Checksum inválido")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto para escuchar:")
	portInput, _ := reader.ReadString('\n')
	portInput = strings.TrimSpace(portInput)
	port, errPort := strconv.Atoi(portInput)

	if errPort != nil {
		fmt.Println("Puerto inválido.")
		return
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error al abrir el puerto:", err)
		return
	}
	defer ln.Close()

	fmt.Printf("Escuchando en el puerto %d...\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

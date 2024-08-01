import socket
import random

def fletcher16(message):
    sum1 = 0
    sum2 = 0
    for byte in message:
        sum1 = (sum1 + byte) % 255
        sum2 = (sum2 + sum1) % 255
    return (sum2 << 8) | sum1

def calcularIntegridad(data):
    message = [int(data[i:i+8], 2) for i in range(0, len(data), 8)]

    checksum = fletcher16(message)

    return "".join(data) + format(checksum, '016b')

def stringToBinary(data):
    return ''.join(format(ord(i), '08b') for i in data)

def ruido(message, probability):
    message = list(message)
    for i in range(len(message)):
        if random.random() < probability:
            message[i] = str(int(message[i]) ^ 1)
    return ''.join(message)

def enviarInformacion(ip, port, message):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((ip, port))
        s.sendall(message.encode())
    print("Mensaje exitosamente enviado.")

if __name__ == '__main__':
    # Capa de aplicación
    data = input("Introduce el mensaje: ")
    
    # Capa de presentación
    data = stringToBinary(data + "\n")
    print("Mensaje en binario: ", data)

    # Capa de enlace
    message = calcularIntegridad(data)
    print("Mensaje con Fletcher-16: ", message)
    ip = "127.0.0.1"
    port = 65432

    # Ruido
    probability = float(input("Introduce la probabilidad p de ruido en su mensaje (1/p): "))
    if probability <= 0:
        print("Probabilidad inválida")
        exit()

    message = ruido(message,1/probability)
    print("Mensaje con ruido: ", message)

    # Capa transmisión
    enviarInformacion(ip, port, message+"\n")


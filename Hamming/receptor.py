
def hammingReciever(message):
    message = list(message)
    n = len(message)

    parityNum = [2**i for i in range(n) if 2**i <= n]
    
    error = 0
    for par in parityNum:
        xor = 0
        for i in range(1, n + 1):
            if i & par == par:
                xor ^= int(message[-i])
        if xor != 0:
            error += par

    if error >= 0:
        print("Error en el bit: ", error + 1)
        message[error - 1] = str(int(message[error - 1])^1)
        print("Mensaje corregido: ", "".join(message))
    else:
        print("No se detectaron errores.")

    originalMessage = []
    for i in range(n):
        if i not in parityNum:
            originalMessage.append(message[i])
    
    return "".join(originalMessage)


if __name__ == "__main__":
    message = input("Ingrese el mensaje recibido: ")
    print("Mensaje decodificado: ", hammingReciever(message))

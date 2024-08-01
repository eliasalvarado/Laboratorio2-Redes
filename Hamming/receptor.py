def hammingReciever(message):
    def correct_block(block):
        n = len(block)
        parityNum = [2**i for i in range(n) if 2**i <= n]
        
        error = 0
        for par in parityNum:
            xor = 0
            for i in range(1, n + 1):
                if i & par != 0:
                    xor ^= int(block[i-1])
            if xor != 0:
                error += par

        if error != 0:
            block[error - 1] = str(int(block[error - 1]) ^ 1)

        originalMessage = []
        for i in range(1, n + 1):
            if i not in parityNum:
                originalMessage.append(block[i - 1])
        
        return "".join(originalMessage)

    blocks = [message[i:i+7] for i in range(0, len(message), 7)]
    corrected_message = []

    for block in blocks:
        corrected_message.append(correct_block(list(block)))

    return "".join(corrected_message)

if __name__ == "__main__":
    message = input("Ingrese el mensaje recibido: ")
    print("Mensaje decodificado: ", hammingReciever(message))

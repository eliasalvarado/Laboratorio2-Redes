import socket


def verificarIntegridad(message):
    def correjirMensaje(block):
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
            print("Error en el bit ", error, " del bloque ", block)
            block[error - 1] = str(int(block[error - 1]) ^ 1)

        originalMessage = []
        for i in range(1, n + 1):
            if i not in parityNum:
                originalMessage.append(block[i - 1])
        
        return "".join(originalMessage)

    blocks = [message[i:i+7] for i in range(0, len(message), 7)]
    corrected_message = []

    for block in blocks:
        corrected_message.append(correjirMensaje(list(block)))

    return "".join(corrected_message)

def recibirInformacion(ip, port):
    print("Esperando conexión...")
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind((ip, port))
    s.listen()
    conn, addr = s.accept()
    with conn:
        print('Conexión entrante desde ', addr)
        while True:
            data = conn.recv(1024)
            if not data:
                break
            message = data.decode()
            return message[:-1]

def decodificarMensaje(messageBinario):
    message = ""
    for i in range(0, len(messageBinario), 8):
        message += chr(int(messageBinario[i:i+8], 2))
    return message
    

if __name__ == "__main__":
    ip = "127.0.0.1"
    port = 65432

    # Capa transmisión
    messageBinario = recibirInformacion(ip, port)
    print("Mensaje recibido en binario: ", messageBinario)

    # Capa de enlace
    messageHamming = verificarIntegridad(messageBinario)
    print("Mensaje recibido con Hamming: ", messageHamming)

    # Capa de presentación
    message = decodificarMensaje(messageHamming)

    # Capa de aplicación
    print("Mensaje recibido: '", message, "'")

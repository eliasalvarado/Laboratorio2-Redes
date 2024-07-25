
def fletcher16(message):
    sum1 = 0
    sum2 = 0
    for byte in message:
        sum1 = (sum1 + byte) % 255
        sum2 = (sum2 + sum1) % 255
    return (sum2 << 8) | sum1

def emisor(data):
    message = [int(data[i:i+8], 2) for i in range(0, len(data), 8)]

    checksum = fletcher16(message)

    return "".join(data) + format(checksum, '016b')

if __name__ == '__main__':
    data = input("Introduce el mensaje: ")
    print(emisor(data))



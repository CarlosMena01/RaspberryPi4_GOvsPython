import time
from PIL import Image
import numpy as np
from scipy.fftpack import fft2, fftshift

# Abrir imagen
img = Image.open('../image.jpg')
pixels = np.array(img)

# Crear archivo de tiempo
time_file = open("time.txt", "w")

for n in range(1,100):
    # Realizar n transformadas FFT
    start = time.time()
    for i in range(n):
        # Realizar FFT
        fft = fft2(pixels)
        fft_shift = fftshift(fft)
        pixels = fft_shift

    elapsed = time.time() - start

    # Escribir iteración/tiempo en archivo
    time_file.write(str(n) + "/" + str(elapsed) + "\n")
time_file.close()

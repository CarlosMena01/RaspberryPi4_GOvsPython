import time
from PIL import Image
import numpy as np
from scipy.fftpack import fft2, fftshift

from progress.bar import Bar

# Parámetros del programa
m = 30 # Máximo número de FFTs consecutivas

# Creamos la barra de progreso
bar = Bar('Procesando:', max = m)

# Abrir imagen
img = Image.open('../image.jpg')
pixels = np.array(img)

# Crear archivo de tiempo
time_file = open("time.txt", "w")

for n in range(1,m + 1):
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

    # Actualizamos la barra de progreso
    bar.next()

bar.finish()
time_file.close()

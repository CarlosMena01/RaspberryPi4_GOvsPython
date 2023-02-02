import time
from PIL import Image
import numpy as np
from scipy.fftpack import fft2, fftshift

# Abrir imagen
img = Image.open('../image.jpg')
pixels = np.array(img)

# Crear archivo de tiempo
time_file = open("time.txt", "w")

# Realizar n transformadas FFT
n = 10
start = time.time()
for i in range(n):
    # Realizar FFT
    fft = fft2(pixels)
    fft_shift = fftshift(fft)

elapsed = time.time() - start

# Escribir tiempo en archivo
time_file.write(str(elapsed))
time_file.close()

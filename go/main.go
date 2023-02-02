package main

import (
	"fmt"
    "os"
    "time"
    "image"
    _ "image/jpeg"
    "github.com/mjibson/go-dsp/fft"
    "sync"
)

func main() {
	// Parámetros del programa
	m := 40 //Máximo número de FFTs consecutivas

    // Abrir imagen
    file, err := os.Open("../image.jpg")
    if err != nil {
    panic(err)
    }
    defer file.Close()
    img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }
	// Obtener matriz de pixeles
	bounds := img.Bounds()

	pixels := make([][]complex128, bounds.Dy())

	for y := 0; y < bounds.Dy(); y++ {
		pixels[y] = make([]complex128, bounds.Dx())
		for x := 0; x < bounds.Dx(); x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			pixels[y][x] = complex(float64(r),0)
		}
	}

    // Crear archivo de tiempo
    timeFile, err := os.Create("time.txt")
    if err != nil {
        panic(err)
    }
    defer timeFile.Close()

	for n := 0; n < (m+1); n++ {
		// Realizar n transformadas FFT de forma concurrente
		var wg sync.WaitGroup
		start := time.Now()
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// Realizar FFT
				pixels = fft.FFT2(pixels)
			}()
		}
		wg.Wait()
		elapsed := time.Since(start)
	
		// Escribir tiempo en archivo
		message := fmt.Sprintf("%d/%.6f \n",n,elapsed.Seconds())
		timeFile.WriteString(message)	
	}
}

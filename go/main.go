package main

import (
    "os"
    "time"
    "image"
    _ "image/jpeg"
    "github.com/mjibson/go-dsp/fft"
    "sync"
)

func main() {
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

    // Realizar n transformadas FFT de forma concurrente
    n := 10
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
    timeFile.WriteString(elapsed.String())
}

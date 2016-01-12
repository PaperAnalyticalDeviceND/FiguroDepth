package main

import (
	"bufio"
	"image"
	_ "image/jpeg"
	"log"
	"math"
	"os"
)

func main() {
	fCenter, err := os.Open("center.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	iCenter, _, err := image.Decode(bufio.NewReader(fCenter))
	if err != nil {
		log.Fatalln(err)
	}
	fCenter.Close()

	fRight, err := os.Open("right.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	iRight, _, err := image.Decode(bufio.NewReader(fRight))
	if err != nil {
		log.Fatalln(err)
	}
	fCenter.Close()

	b := iCenter.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		cR, cG, cB, _ := iCenter.At(x, 124).RGBA()

		bestX := 0
		bestDiff := math.MaxFloat64

		for rX := b.Min.X; rX < b.Max.X; rX++ {
			rR, rG, rB, _ := iRight.At(rX, 124).RGBA()

			dR := float64(cR - rR)
			dG := float64(cG - rG)
			dB := float64(cB - rB)

			diff := math.Sqrt((dR * dR) + (dG * dG) + (dB * dB))
			if diff < bestDiff {
				bestX = rX
				bestDiff = diff
			}
		}

		log.Println("[", x, "] Best:", bestX, " Diff:", bestDiff)
	}
}

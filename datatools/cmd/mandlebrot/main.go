// Reference: https://gist.github.com/esimov/970a3816dda12e4059e3
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/shiena/ansicolor"
)

const (
	WIDTH  int     = 70
	HEIGHT int     = 120
	MIN_X  float64 = -2.0
	MAX_X  float64 = 1.0
	MIN_Y  float64 = -1.0
	MAX_Y  float64 = 1.0
	MAX_IT int     = 1000
)

var isColor bool = true
var wg sync.WaitGroup

func main() {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(`Usage go run mandelbrot_cli.go [--]
			-c --color		generate ASCII mandelbrot in color
			-m --mono		generate ASCII mandelbrot in monochrome`)
			os.Exit(1)
		}
		switch os.Args[1] {
		case "--color", "-c":
			isColor = true
		case "--mono", "-m":
			isColor = false
		}
	}
	charTable := map[int]string{1: "~", 2: "#", 3: "+", 4: "$", 5: "%", 6: "^", 7: "*", 8: "'", 9: "`"}
	ansiColors := map[int]string{1: "\x1b[41m", 2: "\x1b[42m", 3: "\x1b[43m", 4: "\x1b[44m", 5: "\x1b[45m", 6: "\x1b[47m", 7: "\x1b[100m", 8: "\x1b[46m", 9: "\x1b[101m"}

	for row := 0; row < WIDTH; row++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()

			for col := 0; col < HEIGHT; col++ {
				var x float64 = MIN_X + (MAX_X-MIN_X)*float64(row)/float64(WIDTH)
				var y float64 = MIN_Y + (MAX_Y-MIN_Y)*float64(col)/float64(HEIGHT)
				var i = mandelIter(x, y, MAX_IT)
				if i < MAX_IT {
					if i > 5 {
						if isColor {
							fmt.Fprintf(w, "@%s%s%s%s%s", "\x1b[37m", "\x1b[1m", "\x1b[30m", "\x1b[41;32m", "\x1b[0m")
						} else {
							fmt.Printf("@")
						}
					} else {
						if _, ok := charTable[i]; ok {
							if isColor {
								fmt.Fprintf(w, "%s%s", charTable[i], ansiColors[i])
							} else {
								fmt.Printf("%s", charTable[i])
							}
						}
					}
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}(row)
	}

	wg.Wait()
}

func mandelIter(cx, cy float64, maxIter int) int {
	var x, y float64 = 0.0, 0.0
	var iteration int = 0

	for x*x+y*y <= 4 && iteration < MAX_IT {
		var xx float64 = x*x - y*y + cx
		var xy float64 = 2*x*y + cy
		x = xx
		y = xy

		iteration++
	}
	return iteration
}

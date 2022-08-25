package main

import (
	"fmt"
	//"math"
	"time" // used to calculate seed for terrain randomization
	noise "github.com/ojrac/opensimplex-go" // used for calculating terrain height and variation
)


const (
	chunks_amount = 185
	columns = 56
)

var (
	// for testing
	// seed = int64(123)
	seed = time.Now().UnixNano()
	octaves = 16 // 16 octaves is the max for a 185x56 world size (after that nothing changes)
	frequency = 0.03
	amplitude = 1.0
)

type chunk struct {
	blocks [columns]string
}

// IN PLAN
// checking command-line arguments & settings variables
/*func init() {
	args := os.Args
	
	// width
	if len(args) >= 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		return i
	}
	return 100
}*/

func main() {
	chunks := generate_chunks(seed)
	
	for i := 0; i < columns; i++ {
		var column string
		for j := 0; j < len(chunks); j++ {
			column += chunks[j].blocks[i]
		}
		fmt.Println(column)
	}
}

func generate_chunks(seed int64) [chunks_amount]chunk {
	var chunks [chunks_amount]chunk
	n := noise.New(seed)

	for i := 0; i < chunks_amount; i++ {
		chunks[i] = chunk{blocks: randomize_blocks(n, i)}
	}

	return chunks
}

func randomize_blocks(n noise.Noise, c int) [columns]string {
	var blocks [columns]string
	
	for i := 0; i < columns; i++ {
		var height float64
		
		// initialize the frequency & amplitude for later change
		freq := frequency
		ampl := amplitude

		// loop through each octave
		for j := 0; j < octaves; j++ {
			// each higher octave has half the amplitude & double the frequency
			x := float64(c)*freq
			y := float64(i)*freq
			height += float64(n.Eval2(x, y)*ampl)
			freq *= 2
			ampl *= 0.5
		}

		if height >= 0.5 {
			blocks[i] = "#"
		} else {
			blocks[i] = " "
		}
	}
	return blocks
}

/*func generate_island_height_map() ([][]float64) {
	//var height_map [chunks_amount][columns]float64
	height_map := make([][]float64, chunks_amount, columns)
	size_x := chunks_amount
	size_y := columns

	// define center
	center_x := size_x * 0.5
	center_y := size_y * 0.5

	// generate height map
	for x := 0; x < size_x; x++ {
		for y := 0; y < size_y; y++ {
			// distance = sqrt( (x2 - x1)^2 + (y2 - y1)^2 )
			distance_x := (center_x - x) * (center_x - x) // x^2
			distance_y := (center_y - y) * (center_y - y) // y^2

			distance := float64(math.Sqrt(distance_x + distance_y))

			// normalize values
			distance /= 128.0
			
			// define distance
			height_map[x][y] = distance
		}
	}

	return height_map
}*/

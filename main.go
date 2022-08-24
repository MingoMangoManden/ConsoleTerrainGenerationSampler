package main

import (
	"fmt"
	"time" // used to calculate seed for terrain randomization
	noise "github.com/ojrac/opensimplex-go" // used for calculating terrain height and variation
)


const (
	chunks_amount = 185
	columns = 56
	
	frequency = 0.03
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
	seed := time.Now().UnixNano()
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
		height := n.Eval2(float64(c)*frequency, float64(i)*frequency)

		if height >= 0.5 {
			blocks[i] = "#"
		} else {
			blocks[i] = " "
		}
	}
	return blocks
}

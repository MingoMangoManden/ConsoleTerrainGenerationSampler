package main

import (
	"fmt"
	//"math/rand"
	"time"
	noise "github.com/ojrac/opensimplex-go"
)

const (
	chunks_amount = 256
	coloumns = 32

	frequency = 0.03
)

type chunk struct {
	blocks [coloumns]string
}

func main() {
	chunks := generate_chunks()
	
	for i := 0; i < coloumns; i++ {
		var coloumn string
		for j := 0; j < len(chunks); j++ {
			coloumn += chunks[j].blocks[i]
		}
		fmt.Println(coloumn)
	}
}

// seed int32
func generate_chunks() [chunks_amount]chunk {
	var chunks [chunks_amount]chunk
	seed := time.Now().UnixNano()
	n := noise.New(seed)

	for i := 0; i < chunks_amount; i++ {
		chunks[i] = chunk{blocks: randomize_blocks(n, i)}
	}

	return chunks
}

func randomize_blocks(n noise.Noise, c int) [coloumns]string {
	var blocks [coloumns] string
	
	//seed := time.Now().UnixNano()
	//rand.Seed(seed)

	for i := 0; i < coloumns; i++ {
		height := n.Eval2(float64(c)*frequency, float64(i)*frequency)
		//fmt.Println(height)

		if height >= 0.5 {
			blocks[i] = "#"
		} else {
			blocks[i] = " "
		}
	}
	return blocks
}

package main

import (
	"math"
	"math/rand"
)

type DNA struct {
	Gene    string
	Fitness float32
}

func CreateRandomDNA(length int) DNA {
	var randomGene []byte = make([]byte, length)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

	for i := 0; i < length; i++ {
		randomGene[i] = letters[rand.Intn(len(letters))]
	}
	return DNA{Gene: string(randomGene), Fitness: 0}
}

func (dna *DNA) CalculateFitness(target string) {
	var correctChars int = 0

	for i := 0; i < len(target); i++ {
		if target[i] == dna.Gene[i] {
			correctChars++
		}
	}
	dna.Fitness = float32(correctChars) / float32(len(target))
}

func (dna *DNA) Mutate(mutationRate float32) {
	if rand.Float32() < mutationRate {
		const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
		randomVal := rand.Intn(int(math.Max(float64(len(letters)), float64(len(dna.Gene)))))
		randomLetter := letters[randomVal%len(letters)]
		dna.Gene = dna.Gene[0:randomVal%len(dna.Gene)] + string(randomLetter) + dna.Gene[randomVal%len(dna.Gene)+1:]
	}
}

func (dna *DNA) Crossover(pair DNA) DNA {
	var midPoint int = rand.Intn(len(dna.Gene))
	return DNA{Gene: dna.Gene[:midPoint] + pair.Gene[midPoint:]}
}

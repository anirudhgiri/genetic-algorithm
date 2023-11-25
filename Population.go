package main

import (
	"math/rand"
)

type GenePool struct {
	Population   []DNA
	MutationRate float32
	Target       string
	Generations  int
}

func (pool *GenePool) CalculateAllFitnesses() {
	for i := 0; i < len(pool.Population); i++ {
		pool.Population[i].CalculateFitness(pool.Target)
	}
}

func (pool *GenePool) NaturalSelection() {
	var matingPool []DNA
	var maxFitness = pool.GetBestDNA().Fitness
	for _, dna := range pool.Population {
		n := int((dna.Fitness / maxFitness) * 500)
		for i := 0; i < n; i++ {
			matingPool = append(matingPool, dna)
		}
	}
	nextGenerationPopulation := []DNA{}

	for i := 0; i < len(pool.Population); i++ {
		parentA := matingPool[rand.Intn(len(matingPool))]
		parentB := matingPool[rand.Intn(len(matingPool))]
		child := parentA.Crossover(parentB)
		child.Mutate(pool.MutationRate)
		nextGenerationPopulation = append(nextGenerationPopulation, child)
	}
	pool.Population = nextGenerationPopulation
	pool.CalculateAllFitnesses()
	pool.Generations++
}

func (pool *GenePool) GetBestDNA() DNA {
	var maxFitness float32 = 0
	var bestDNA DNA
	for _, dna := range pool.Population {
		if dna.Fitness > maxFitness {
			maxFitness = dna.Fitness
			bestDNA = dna
		}
	}
	return bestDNA
}

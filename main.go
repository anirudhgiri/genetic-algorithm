package main

import "fmt"

func main() {
	genePool := GenePool{MutationRate: 0.01, Target: "To Be Or Not To Be", Generations: 1}
	for i := 0; i < 1000; i++ {
		genePool.Population = append(genePool.Population, CreateRandomDNA(len(genePool.Target)))
	}
	genePool.CalculateAllFitnesses()

	var maxFitness float32 = 0

	for genePool.GetBestDNA().Fitness != 1 {
		genePool.NaturalSelection()

		if genePool.GetBestDNA().Fitness > maxFitness {
			maxFitness = genePool.GetBestDNA().Fitness
			fmt.Println(maxFitness, genePool.GetBestDNA().Gene, genePool.Generations)
		}

	}
}

package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Report map[int]int

type GA struct {
	uniformRate float64
	mutationRate float64
	tournamentSize int
	elitism bool //keep best individual from old population
	individualConfig IndividualConfig
	populationSize int
	loop int
}

func NewGA(uniform float64, mutationRate float64, tournamentSize int, elitism bool, config IndividualConfig, populationSize int, loop int) *GA {
	rand.Seed(time.Now().UTC().UnixNano())
	return &GA{
		uniform, mutationRate, tournamentSize, elitism, config,populationSize,loop,
	}
}

func (ga *GA) tournamentSelection(population *Population) *Individual {
	tournament := NewPopulation(ga.tournamentSize, false, ga.individualConfig)

	for i := 0; i < ga.tournamentSize; i++ {
		randomIndex := rand.Intn(population.GetSize())
		tournament.SaveIndividual(i, population.GetIndividual(randomIndex))
	}
	return tournament.GetFittest()
}

func (ga *GA) crossOver(indiv1 *Individual, indiv2 *Individual) *Individual {
	newIndividual := NewIndividual(ga.individualConfig)

	for i := 0; i < indiv1.GetSize(); i++ {
		if rand.Float64() <= ga.uniformRate {
			newIndividual.SetGene(i, indiv1.GetGene(i))
		} else {
			newIndividual.SetGene(i, indiv2.GetGene(i))
		}
	}
	return newIndividual
}

func (ga *GA) mutate(individual *Individual) {

	for i := 0; i < individual.GetSize(); i++ {
		if rand.Float64() <= ga.mutationRate {
			gene := rand.Intn(2) == 0
			individual.SetGene(i, gene)
		}
	}
}

func (ga *GA) EvolutionPopulation(population *Population) *Population {
	newPopulation := NewPopulation(population.GetSize(), false, ga.individualConfig)

	elitismOffset := 0
	if ga.elitism {
		newPopulation.SaveIndividual(0, population.GetFittest())
		elitismOffset = 1
	}

	for i := elitismOffset; i < population.GetSize(); i++ {
		individual1 := ga.tournamentSelection(population)
		individual2 := ga.tournamentSelection(population)
		newIndividual := ga.crossOver(individual1, individual2)

		ga.mutate(newIndividual)

		newPopulation.SaveIndividual(i, newIndividual)
	}

	return newPopulation
}

func (ga *GA) Run() {
	report := make(map[int]int)
	population := NewPopulation(ga.populationSize, true, ga.individualConfig)

	for i:=0; i<ga.loop - 1; i++ {
		report[i+1] = population.GetFittest().GetFitness()

		fmt.Printf("Generation: %d > Fittest: %d > Gene: %s\n", i+1, population.GetFittest().GetFitness(), population.GetFittest().ToString())

		population = ga.EvolutionPopulation(population)
	}

	fmt.Printf("Generation: %d > Fittest: %d > Gene: %s\n", ga.loop, population.GetFittest().GetFitness(), population.GetFittest().ToString())
}

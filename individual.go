package main

import (
	"math/rand"
	"time"
)
type IndividualConfig struct {
	GeneLength int
	FitnessCalc FitnessCalcBase
}

type Individual struct {
	genes []bool
	fitness int
	fitnessCalc FitnessCalcBase
}

func NewIndividual(c IndividualConfig) *Individual{
	individual := Individual{
		genes: make([]bool, c.GeneLength),
		fitness: 0,
		fitnessCalc: c.FitnessCalc,
	}

	return &individual
}

func (i *Individual) GetSize() int {
	return len(i.genes)
}

func (i *Individual) GenerateIndividual() {
	rand.Seed(time.Now().UTC().UnixNano())
	for index, _ := range i.genes {
		i.genes[index] = rand.Intn(2) == 0
	}
}

func (i *Individual) GetGene(index int) bool{
	if index < len(i.genes) {
		return i.genes[index]
	} else {
		return false
	}
}

func (i *Individual) SetGene(index int, value bool) {
	if index < len(i.genes) {
		i.genes[index] = value
	}
}

func (i *Individual) GetFitness() int {
	if i.fitness == 0 {
		i.fitness = i.fitnessCalc.GetFitness(i)
	}
	return i.fitness
}

func (i *Individual) ToString() string {
	str := ""
	for _, gene := range i.genes {
		if gene {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}



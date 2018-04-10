package main

type FitnessCalcBase interface {
	GetFitness(individual *Individual) int
}

type FitnessKnapsackCalc struct {
	items []KnapSackWeightPoint
}

type KnapSackWeightPoint struct {
	Weight int
	Price int
}

func NewFitnessKnapsack() FitnessKnapsackCalc {
	weight := []int{35,25,8,51,48,20,3,3,50,18,57,30,44,65,52,67,9,71,13,88,66,77,77,15,69,99,88,5,74,78,20,92,50,15,75,93,99,40,37,92,96,62,14,86,75,41,37,10,58,35}
	price := []int{47,32,66,21,41,24,88,62,29,45,35,173,22,92,85,12,66,82,2,55,96,28,42,10,38,29,4,64,76,88,77,46,49,61,51,40,37,160,84,34,70,93,69,40,60,65,43,98,65,84}
	items := make([]KnapSackWeightPoint, len(weight))

	for i, _ := range weight {
		items[i].Weight = weight[i]
		items[i].Price = price[i]
	}
	return FitnessKnapsackCalc{items}
}

func (f FitnessKnapsackCalc) GetFitness(individual *Individual) int {
	fitness := 0
	totalWeight := 0

	for i := 0; i < individual.GetSize(); i++ {
		if individual.GetGene(i) {
			totalWeight += f.items[i].Weight
			fitness += f.items[i].Price * 1
		}
	}

	if totalWeight > 200 {
		fitness = totalWeight * -1
	}

	return fitness
}
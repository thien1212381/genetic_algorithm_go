package main


type Population struct {
	individuals []*Individual
	individualConfig IndividualConfig
}

func NewPopulation(populationSize int, isInitialise bool, conf IndividualConfig) *Population{
	population := &Population{individualConfig: conf}
	population.individuals = make([]*Individual, populationSize)

	if isInitialise {
		for i := 0; i < population.GetSize(); i++ {
			individual := NewIndividual(population.individualConfig)
			individual.GenerateIndividual()
			population.SaveIndividual(i, individual)
		}
	}
	return population
}

func (p *Population) GetSize() int{
	return len(p.individuals)
}

func (p *Population) SaveIndividual(index int, individual *Individual) {
	if index < p.GetSize() {
		p.individuals[index] = individual
	}
}

func (p *Population) GetIndividual(index int) *Individual {
	if index < p.GetSize() {
		return p.individuals[index]
	} else {
		return nil
	}
}

func (p *Population) GetFittest() *Individual {
	result := p.GetIndividual(0)
	for i:=1; i<p.GetSize(); i++ {
		if p.GetIndividual(i).GetFitness() > result.GetFitness() {
			result = p.GetIndividual(i)
		}
	}
	return result
}

func (p *Population) GetIndividualConfig() IndividualConfig {
	return p.individualConfig
}

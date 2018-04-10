package main

func main()  {
	conf := IndividualConfig{
		GeneLength: 50,
		FitnessCalc: NewFitnessKnapsack(),
	}

	ga := NewGA(0.5,0.015, 10, true, conf, 500, 100)
	ga.Run()
}

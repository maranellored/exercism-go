package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	finalMap := FreqMap{}

	commChannel := make(chan FreqMap)
	for _, s := range strings {
		go func(s string) {
			commChannel <- Frequency(s)
		}(s)
	}

	for j := 0; j < len(strings); j++ {
		tempMap := <-commChannel
		for key, value := range tempMap {
			finalMap[key] += value
		}
	}

	return finalMap
}

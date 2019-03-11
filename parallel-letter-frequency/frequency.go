package letter

//import "time"
import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var wg sync.WaitGroup

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func SafeConcurrentFrequency(word string, ch chan FreqMap) {
	defer wg.Done()
	m := make(FreqMap)

	for _, r := range word {
		m[r]++
	}

	ch <- m
}



func ConcurrentFrequency(words []string) FreqMap {
	
	m := make(FreqMap)
	ch := make(chan FreqMap, len(words))
	for _, word := range words {
		wg.Add(1)
		go SafeConcurrentFrequency(word, ch)
	}

	wg.Wait()
	close(ch)

	for freqMap := range ch {
		
		for key, value := range freqMap {
			m[key] += value
		}
	}

	return m
}



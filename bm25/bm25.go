package main

import (
	"fmt"
	"math"
	"strings"
)

//varies from 1.2-2.0
var k1 float64 = 1.2

//should be around 0.75
var b float64 = 0.75
var docs []string

func main() {
	docs = []string{"Shipment of gold damaged in a fire", "Delivery of silver arrived in a silver truck", "Shipment of gold arrived in a truck"}
	terms := []string{"gold", "silver", "truck", "delivery"}
	// docs = []string{"this is a sample a", "this is another example", "final doc here"}
	// terms := []string{"a", "query", "example", "this"}

	//average documnet length
	L_ave := AverageDocLength()
	N := float64(len(docs))
	for i := 0; i < len(docs); i++ {
		println(docs[i])
		fmt.Printf("%f\n", BM25(terms, docs[i], L_ave, N))
	}
}

//BM25 first param is terms, second is current document
func BM25(terms []string, d string, L_ave float64, N float64) float64 {
	//total number of docs

	D := float64(len(strings.Fields(d)))

	sum := 0.0
	for j := 0; j < len(terms); j++ {
		//current term
		i := terms[j]

		//this is the number of times term i appears in any doc
		numberOfDocOccurences := NumberOfDocOccurrences(i)
		//left hand side of the multiplication
		idf := math.Log10((N - numberOfDocOccurences + 0.5) / (numberOfDocOccurences + 0.5))
		// idf := 0.0
		// if numberOfDocOccurences > 0 {
		// 	idf = math.Log10(N / NumberOfDocOccurrences(i))
		// }

		//right hand side of the multiplication
		tf_id := TermFreq(d, i)

		rightHandSide := ((k1 + 1) * tf_id) / (k1*((1-b)+b*(D/L_ave)) + tf_id)

		//add it to the sum
		sum += (idf * rightHandSide)
	}
	return sum
}

func docLength(d string) float64 {
	return float64(len(d)) - float64(len(strings.Fields(d))-1)
}

func AverageDocLength() float64 {
	sum := 0
	for i := 0; i < len(docs); i++ {
		sum += len(docs[i])
		sum -= (len(strings.Fields(docs[i])) - 1)
	}
	return float64(sum) / float64(len(docs))
}

func NumberOfDocOccurrences(i string) float64 {
	counter := 0
	for k := 0; k < len(docs); k++ {
		words := strings.Fields(docs[k])
		for j := 0; j < len(words); j++ {
			if strings.ToLower(i) == strings.ToLower(words[j]) {

				counter++
				break
			}
		}
	}
	return float64(counter)
}

func TermFreq(d string, term string) float64 {
	words := strings.Fields(d)
	counter := 0
	for i := 0; i < len(words); i++ {

		if strings.ToLower(term) == strings.ToLower(words[i]) {
			counter++
		}
	}
	return float64(counter) // / float64(len(words))
}

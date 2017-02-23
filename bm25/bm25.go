package main

import (
	"math"
	"strings"
)

//varies from 1.2-2.0
var k1 float64 = 1.2
var b float64 = 0.75
var docs []string

func main() {

	docs = []string{"Shipment of gold damaged in a fire", "Delivery of silver arrived in a silver truck", "Shipment of gold arrived in a truck"}
	terms := []string{"gold", "silver", "truck"}
	for i := 0; i < len(docs); i++ {
		println(docs[i])
		println(BM25(terms, docs[i], AverageDocLength()))
	}
}

//BM25 first param is terms, second is current document
func BM25(terms []string, d string, L_ave float64) float64 {

	//total number of docs
	N := float64(len(docs))

	//search terms

	//document length
	L_d := float64(len(d))

	sum := 0.0
	for j := 0; j < len(terms); j++ {
		//current term
		i := terms[j]
		tf_id := TermFreq(d, i)
		idf := math.Log10(N / NumberOfDocOccurrences(i))
		// docsWithTerm := NumberOfDocOccurrences(i)
		// idf := math.Log10(N/docsWithTerm + 0.5)
		rightHandSide := ((k1 + 1) * tf_id) / (k1*((1-b)+b*(L_d/L_ave)) + tf_id)

		sum += (idf * rightHandSide)
	}
	return sum
}

func AverageDocLength() float64 {
	sum := 0
	for i := 0; i < len(docs); i++ {
		sum += len(docs[i])
	}
	return float64(sum) / float64(len(docs))
}

func NumberOfDocOccurrences(i string) float64 {
	counter := 0
	for k := 0; k < len(docs); k++ {
		words := strings.Fields(docs[k])
		for j := 0; j < len(words); j++ {
			if i == words[j] {
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
		if term == words[i] {
			counter++
		}
	}
	return float64(counter) / float64(len(words))
}

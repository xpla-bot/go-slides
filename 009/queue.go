package main

import "log"

// zdefiniowanie własnego typu queue
type queue []string

// metoda enqueue dodaje do kolejki
func (q *queue) enqueue(s string) {
	// wywoałeni funckji wbudowanej append
	// append zwraca nową tablicę, więc trzeba zachować to w zmiennej q
	*q = append(*q, s)
}

func (q *queue) dequeue() string {
	// sprawdzenie czy kolejka jest pusta
	if len(*q) == 0 {
		// jeśli jest to zwracamy pusty element
		return ""
	}

	// zachowanie pierwszego elementu kolejki
	s := (*q)[0]

	// skrócenie kolejki
	*q = (*q)[1:len(*q)]

	// zwrócenie pierwszego elementu z kolejki
	return s
}

func main() {
	// stworzenie kolejki
	var q queue

	// dodanie jednej wartości
	q.enqueue("1")

	// dodanie drugiej wartości
	q.enqueue("2")

	// Wypisanie pierwszej wartości
	log.Print(q.dequeue())

	// Wypisanie pierwszej wartości
	log.Print(q.dequeue())

	// Kolejka jest pusta i zwraca pusty ciąg znków
	log.Print(q.dequeue())
}

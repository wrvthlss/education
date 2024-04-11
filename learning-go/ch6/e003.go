package main

type Prsn struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	const numEntries = 10_000_000

	// Pre-allocating slice with capacity for 10_000_000 entries.
	people := make([]Prsn, numEntries)

	for i := 0; i < numEntries; i++ {
		people[i] = Prsn{
			FirstName: "John",
			LastName:  "Doe",
			Age:       40,
		}
	}

}

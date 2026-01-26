package magaya

type FieldElement struct {
	prime   uint64
	element uint64
}

// NewFieldElement initializes a new FieldElement with the passed arguments
// If prime is not a prime number || element >= prime || element < 0, FieldElement will be nil and an error will be returned
func (fe *FieldElement) NewFieldElement(prime uint64, element uint64) (*FieldElement, error) {
	return &FieldElement{prime: prime, element: element}, nil
}

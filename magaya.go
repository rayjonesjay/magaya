package magaya

type FieldElement struct {
	prime   uint64
	element uint64
}

func NewFieldElement(prime uint64, element uint64) (*FieldElement, error) {
	return &FieldElement{prime: prime, element: element}, nil
}

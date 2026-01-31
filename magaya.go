package magaya

import (
	"fmt"
)

type FieldElement struct {
	prime   uint64
	element uint64
}

func NewFieldElement(prime uint64, element uint64) (*FieldElement, error) {
	return &FieldElement{prime: prime, element: element}, nil
}

func (self *FieldElement) Add(other FieldElement) (*FieldElement,error){
	if self.prime != other.prime {
		return nil , fmt.Errorf("cannot add elements from different fields")
	}
	var result uint64 = (self.element+other.element) % self.prime
	return &FieldElement{self.prime,result},nil
}

func (self *FieldElement) Sub(other FieldElement) (*FieldElement,error){
	if self.prime != other.prime {
		return nil, fmt.Errorf("cannot subtract elements from different fields")
	}
	var result uint64 = (self.element-other.element) % self.prime
	return &FieldElement{self.prime,result},nil
}

func (self *FieldElement) Mul(other FieldElement) (*FieldElement,error){
	if self.prime != other.prime {
		return nil, fmt.Errorf("cannot multiply elements from different fields")
	}
	var result uint64 = (self.element*other.element) % self.prime
	return &FieldElement{self.prime,result},nil 
}

func (self *FieldElement) Div(other FieldElement) (*FieldElement, error){
	if !self.haveSamePrime(other) {
		return nil, fmt.Errorf("cannot divide elements from different fields")
	}
	var inverse uint64 = 0;
	var i uint64 = 0;
	for i = 1; i < self.prime; i++ {
		if (other.element*i) % self.prime == 1 {
			inverse = i;
			break;
		}
	}
	return &FieldElement{self.prime,(self.element*inverse)%self.prime}, nil
}

func (self *FieldElement) haveSamePrime(other FieldElement) bool {
	return self.prime == other.prime
}
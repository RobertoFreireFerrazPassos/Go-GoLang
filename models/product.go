package models

type Product struct {
	Name     string
	Category Category
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type CategorySpecification struct {
	Category Category
}

func (spec CategorySpecification) IsSatisfied(p *Product) bool {
	return p.Category == spec.Category
}

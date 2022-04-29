package books

type Specification interface {
	IsSatisfied(p *Book) bool
}

type CategorySpecification struct {
	Category Category
}

func (spec CategorySpecification) IsSatisfied(p *Book) bool {
	return p.Category == spec.Category
}

type Filter struct{}

func (f *Filter) Filter(
	books []Book, spec Specification) []*Book {
	result := make([]*Book, 0)
	for i, b := range books {
		if spec.IsSatisfied(&b) {
			result = append(result, &books[i])
		}
	}
	return result
}

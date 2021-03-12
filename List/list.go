package List

type List interface {
	GetByIndex(index int) (interface{}, error)
	Add(ell interface{})
}

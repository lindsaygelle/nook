package nook

type Code struct {
	Value string
}

func (v Code) Ok() bool {
	return v.Value != ""
}

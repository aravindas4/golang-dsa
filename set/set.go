package set

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) GetElements() map[int]bool {
	return set.integerMap
}

func (set *Set) ContainsElement(value int) bool {
	var exists bool
	_, exists = set.integerMap[value]
	return exists
}

func (set *Set) AddElement(value int) {
	if !set.ContainsElement(value) {
		set.integerMap[value] = true
	}
}

func (set *Set) DeleteElement(value int) {
	delete(set.integerMap, value)
}

func (set *Set) CountElements() int {
	return len(set.integerMap)
}

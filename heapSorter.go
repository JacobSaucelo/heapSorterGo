package main

func (c *Cars) InsertHeap(car CarTypes) {
	c.cars = append(c.cars, car)
	c.HeapifyUp(len(c.cars) - 1)
}

func (c *Cars) HeapifyUp(index int) {
	for c.cars[Parent(index)].Price_in_thousands < c.cars[index].Price_in_thousands {
		c.Swap(Parent(index), index)
		index = Parent(index)
	}
}

func (c *Cars) Swap(index1, index2 int) {
	c.cars[index1], c.cars[index2] = c.cars[index2], c.cars[index1]
}

func Parent(i int) int {
	return (i - 1) / 2
}

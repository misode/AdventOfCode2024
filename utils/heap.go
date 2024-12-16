package utils

type Heap[T any] struct {
	data     []*T
	priority func(*T) int
}

func MakeHeap[T any](priority func(*T) int) Heap[T] {
	return Heap[T]{[]*T{}, priority}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Push(node *T) {
	h.data = append(h.data, node)
	h.up(len(h.data) - 1)
}

func (h *Heap[T]) Pop() *T {
	if len(h.data) == 0 {
		panic("Heap is empty")
	}
	top := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.down(0)
	}
	return top
}

func (h *Heap[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || h.priority(h.data[parent]) <= h.priority(h.data[i]) {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *Heap[T]) down(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.priority(h.data[left]) < h.priority(h.data[smallest]) {
			smallest = left
		}
		if right < n && h.priority(h.data[right]) < h.priority(h.data[smallest]) {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

package network

import (
	"container/list"
	"sync"
)

const (
	CAPACITY_INFINITE = -1
)

/* Session Index 관리 */
type Deque struct {
	sync.RWMutex
	container *list.List
	capacity  int
}

func (d *Deque) Append(item interface{}) (int, bool) {
	d.Lock()
	defer d.Unlock()

	count := d.container.Len()
	if d.capacity < 0 || count < d.capacity {
		d.container.PushBack(item)
		count += 1
		return count, true
	}

	return 0, false
}

func (d *Deque) Prepend(item interface{}) (int, bool) {
	d.Lock()
	defer d.Unlock()

	count := d.container.Len()
	if d.capacity < 0 || count < d.capacity {
		d.container.PushFront(item)
		count += 1
		return count, true
	}

	return 0, false
}

func (d *Deque) Pop() interface{} {
	d.Lock()
	defer d.Unlock()

	var item interface{} = nil
	var lastContainerItem *list.Element = nil

	lastContainerItem = d.container.Back()
	if lastContainerItem != nil {
		item = d.container.Remove(lastContainerItem)
	}

	return item
}

func (d *Deque) Shift() interface{} {
	d.Lock()
	defer d.Unlock()

	var item interface{} = nil
	var firstContainerItem *list.Element = nil

	firstContainerItem = d.container.Front()
	if firstContainerItem != nil {
		item = d.container.Remove(firstContainerItem)
	}

	return item
}

func (d *Deque) First() interface{} {
	d.RLock()
	defer d.RUnlock()

	item := d.container.Front()
	if item != nil {
		return item.Value
	}

	return nil
}

func (d *Deque) Last() interface{} {
	d.RLock()
	defer d.RUnlock()

	item := d.container.Back()
	if item != nil {
		return item.Value
	}

	return nil
}

func (d *Deque) Size() int {
	d.RLock()
	defer d.RUnlock()

	return d.container.Len()
}

func (d *Deque) Capacity() int {
	d.RLock()
	defer d.RUnlock()

	return d.capacity
}

func (d *Deque) Empty() bool {
	d.RLock()
	defer d.RUnlock()

	return d.container.Len() == 0
}

func (d *Deque) Full() bool {
	d.RLock()
	defer d.RUnlock()

	return d.capacity >= 0 && d.container.Len() >= d.capacity
}

func NewDeque() *Deque {
	return NewCappedDeque(CAPACITY_INFINITE)
}

func NewCappedDeque(capacity int) *Deque {
	return &Deque{
		container: list.New(),
		capacity:  capacity,
	}
}

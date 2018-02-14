package hash

import "myalgorithms/list"
import "fmt"

type bucket struct {
	key  interface{}
	data *lists.List
}

type hashalgo func(d interface{}) int

var index int

//Hashbucket data structure for creating hash table
type Hashbucket struct {
	size int
	b    []bucket
	fn   hashalgo
}

func (h *Hashbucket) getHashKey(d interface{}) int {

	if h.fn == nil {
		index++
		index = index % len(h.b)
		return index
	}
	return h.fn(d)

}

func Newhashtable(len int, algo hashalgo) *Hashbucket {
	return &Hashbucket{
		b:    make([]bucket, len),
		size: 0,
		fn:   algo,
	}
}

//Search hash table
func (h *Hashbucket) Search(d interface{}) bool {
	key := h.getHashKey(d)

	return h.b[key].data.Search(d)

}

func (h *Hashbucket) Insert(d interface{}) bool {
	key := h.getHashKey(d)

	if h.b[key].data == nil {
		h.b[key].key = key
		h.b[key].data, _ = lists.New()
	}
	h.size++
	return h.b[key].data.Insert(d)

}

func (h *Hashbucket) Delete(d interface{}) bool {
	key := h.getHashKey(d)

	return h.b[key].data.Delete(d)
}

func (h *Hashbucket) Print() {
	for i, b := range h.b {
		fmt.Println("bucket ", i)
		b.data.Print()
	}
}

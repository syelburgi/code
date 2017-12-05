package hash

import "myalgorithms/list"

type bucket struct {
	key  interface{}
	data *lists.List
}

type hashalgo interface {
	getHashKey(d interface{}) int
}

var index int

//Hashbucket data structure for creating hash table
type Hashbucket struct {
	size int
	b    []bucket
	hashalgo
}

func (h *Hashbucket) getHashKey(d interface{}) int {
	index = index % (len(h.b) - 1)
	index++
	return index
}

func Newhashtable(len int, algo hashalgo) *Hashbucket {
	return &Hashbucket{
		b:        make([]bucket, len),
		size:     0,
		hashalgo: algo,
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

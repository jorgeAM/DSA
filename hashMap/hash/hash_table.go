package hash

const arraySize = 7

type HashTable struct {
	array [arraySize]*Bucket
}

func Init() *HashTable {
	hashTable := new(HashTable)

	for i := range hashTable.array {
		hashTable.array[i] = new(Bucket)
	}

	return hashTable
}

func (h *HashTable) hash(key string) int {
	var counter int

	for _, w := range key {
		counter += int(w)
	}

	return counter % arraySize
}

func (h *HashTable) Insert(key string) {
	hash := h.hash(key)

	h.array[hash].insert(key)
}

func (h *HashTable) Search(key string) bool {
	hash := h.hash(key)

	return h.array[hash].search(key)
}

func (h *HashTable) Delete(key string) {
	hash := h.hash(key)

	h.array[hash].delete(key)
}

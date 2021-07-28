package hash

type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func (b *Bucket) insert(key string) {
	node := &BucketNode{key: key}
	node.next = b.head
	b.head = node
}

func (b *Bucket) search(key string) bool {
	if b.head == nil {
		return false
	}

	current := b.head

	for current != nil {
		if current.key == key {
			return true
		}

		current = current.next
	}

	return false
}

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

func (b *Bucket) delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next

		return
	}

	prevNode := b.head

	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next

			return
		}

		prevNode = prevNode.next
	}
}

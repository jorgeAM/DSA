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

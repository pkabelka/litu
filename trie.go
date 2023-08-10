package litu

type Trie[K comparable, V any] struct {
	Value    V
	next     map[K]*Trie[K, V]
	previous *Trie[K, V]
}

func (node *Trie[K, V]) Insert(value V, keyPath ...K) *Trie[K, V] {
	tmp := node
	for _, nodeKey := range keyPath {

		nextNode, nodeExists := tmp.next[nodeKey]
		if !nodeExists {
			if tmp.next == nil {
				tmp.next = map[K]*Trie[K, V]{}
			}

			nextNode = &Trie[K, V]{}
			nextNode.previous = tmp
			tmp.next[nodeKey] = nextNode
		}

		tmp = nextNode
	}

	tmp.Value = value
	return tmp
}

func (root *Trie[K, V]) Get(keyPath ...K) *Trie[K, V] {
	tmp := root
	for _, nodeKey := range keyPath {

		nextNode, nodeExists := tmp.next[nodeKey]
		if !nodeExists {
			return nil
		}

		tmp = nextNode
	}

	return tmp
}

func (root *Trie[K, V]) Remove(keyPath ...K) bool {
	tmp := root.Get(keyPath...)

	if tmp == nil || len(tmp.next) > 0 {
		return false
	}

	i := len(keyPath) - 1
	for tmp.previous != root.previous {
		delete(tmp.previous.next, keyPath[i])

		if len(tmp.previous.next) > 0 {
			break
		}

		tmp.previous.next = nil

		tmp = tmp.previous
		i--
	}

	return true
}

func (node *Trie[K, V]) Each(callback func(key K, node *Trie[K, V])) {
	nodeStack := make([]*Trie[K, V], 0, len(node.next))
	keyStack := make([]K, 0, len(node.next))
	visited := map[*Trie[K, V]]struct{}{}

	for k, next := range node.next {
		nodeStack = append(nodeStack, next)
		keyStack = append(keyStack, k)
	}

	for len(nodeStack) > 0 {
		lastNode := nodeStack[len(nodeStack)-1]
		lastKey := keyStack[len(keyStack)-1]

		_, nodeVisited := visited[lastNode]
		if len(lastNode.next) > 0 && !nodeVisited {
			for k, next := range lastNode.next {
				nodeStack = append(nodeStack, next)
				keyStack = append(keyStack, k)
				visited[lastNode] = struct{}{}
			}
			continue
		}

		callback(lastKey, lastNode)
		nodeStack = nodeStack[:len(nodeStack)-1]
		keyStack = keyStack[:len(keyStack)-1]
	}
}

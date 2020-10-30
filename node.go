package hedera

import (
	"math"
	"math/rand"
	"time"
)

type node struct {
	accountID AccountID
	address   string
	delay     int64
	lastUsed  *int64
}

type nodes struct {
	nodes []node
}

func newNode(accountID AccountID, address string) node {
	return node{
		accountID: accountID,
		address:   address,
		delay:     250,
		lastUsed:  nil,
	}
}

func (node *node) isHealthy() bool {
	if node.lastUsed != nil {
		lastUsed := *node.lastUsed
		return lastUsed+node.delay < time.Now().UTC().UnixNano()/1e6
	}

	return true
}

func (node *node) increaseDelay() {
	*node.lastUsed = time.Now().UTC().UnixNano() / 1e6
	node.delay = int64(math.Min(float64(node.delay)*2, 8000))
}

func (node *node) decreaseDelay() {
	node.delay = int64(math.Max(float64(node.delay)/2, 250))
}

func (node *node) wait() {
	delay := *node.lastUsed
	if node.lastUsed != nil {
		delay = delay + node.delay - time.Now().UTC().UnixNano()/1e6
	} else {
		delay = node.delay - time.Now().UTC().UnixNano()/1e6
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func (s nodes) Len() int {
	return len(s.nodes)
}
func (s nodes) Swap(i, j int) {
	s.nodes[i], s.nodes[j] = s.nodes[j], s.nodes[i]
}
func (s nodes) Less(i, j int) bool {
	if s.nodes[i].isHealthy() && s.nodes[j].isHealthy() {
		if rand.Int63n(1) == 0 {
			return true
		} else {
			return false
		}
	} else if s.nodes[i].isHealthy() && !s.nodes[j].isHealthy() {
		return true
	} else if !s.nodes[i].isHealthy() && s.nodes[j].isHealthy() {
		return false
	} else {
		var aLastUsed = s.nodes[i].lastUsed
		var bLastUsed = s.nodes[j].lastUsed
		if aLastUsed == nil {
			*aLastUsed = 0
		}

		if bLastUsed == nil {
			*bLastUsed = 0
		}

		if *aLastUsed+s.nodes[i].delay < *bLastUsed+s.nodes[j].delay {
			return true
		} else {
			return false
		}
	}
}
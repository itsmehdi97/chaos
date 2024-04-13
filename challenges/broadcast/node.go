package broadcast

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Broadcast struct {
	node      *maelstrom.Node
	seen      []int
	neighbors []string
}

func NewBroadcast() *Broadcast {
	return &Broadcast{
		node:      maelstrom.NewNode(),
		seen:      make([]int, 0),
		neighbors: make([]string, 0),
	}
}

func (b *Broadcast) Run() {
	b.setupHandlers()

	if err := b.node.Run(); err != nil {
		log.Fatal(err)
	}
}

func (b *Broadcast) setupHandlers() {
	b.node.Handle("topology", b.onTopology)
	b.node.Handle("broadcast", b.onBroadcast)
	b.node.Handle("read", b.onRead)
}

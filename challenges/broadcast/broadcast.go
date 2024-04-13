package broadcast

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastMsg struct {
	Typ       string `json:"type"`
	Message   int    `json:"message"`
	MessageID int    `json:"msg_id"`
}

func (b *Broadcast) onBroadcast(msg maelstrom.Message) error {
	var body BroadcastMsg
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	go func() {
		for _, nodeID := range b.neighbors {
			if nodeID != msg.Src && newMsg(&b.seen, body.Message) {
				b.node.RPC(nodeID, map[string]any{
					"type":    "broadcast",
					"message": body.Message,
				}, nil)
			}
		}
		b.seen = append(b.seen, body.Message)
	}()

	return b.node.Reply(msg, map[string]string{
		"type": "broadcast_ok",
	})
}

func newMsg(seen *[]int, i int) bool {
	for _, item := range *seen {
		if item == i {
			return false
		}
	}
	return true
}

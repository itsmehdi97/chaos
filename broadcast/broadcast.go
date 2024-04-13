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

	b.seen = append(b.seen, body.Message)
	go func() {
		for _, nodeID := range b.neighbors {
			if nodeID != msg.Src {
				b.node.RPC(nodeID, map[string]any{
					"type":    "broadcast",
					"message": body.Message,
				}, nil)
			}
		}
	}()

	return b.node.Reply(msg, map[string]string{
		"type": "broadcast_ok",
	})
}

package broadcast

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type TopoMsg struct {
	Typ       string              `json:"type"`
	Topology  map[string][]string `json:"topology"`
	MessageID int                 `json:"msg_id"`
}

func (b *Broadcast) onTopology(msg maelstrom.Message) error {
	var body TopoMsg
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	b.neighbors = body.Topology[b.node.ID()]

	return b.node.Reply(msg, map[string]any{
		"type": "topology_ok",
	})
}

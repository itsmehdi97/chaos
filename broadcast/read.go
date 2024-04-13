package broadcast

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ReadResponse struct {
	Typ       string `json:"type"`
	Messages  []int  `json:"messages"`
	MessageID int    `json:"msg_id"`
}

func (b *Broadcast) onRead(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	return b.node.Reply(msg, ReadResponse{
		Typ:      "read_ok",
		Messages: b.seen,
	})

}

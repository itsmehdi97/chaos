package uid

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Run() {
	n := maelstrom.NewNode()

	i := 0

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = fmt.Sprintf("%s:%d", n.ID(), i)
		i += 1

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

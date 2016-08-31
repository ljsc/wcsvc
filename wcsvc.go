package wordcount

import (
	"strings"

	context "golang.org/x/net/context"
)

type Server struct{}

func (s Server) GetStats(ctx context.Context, m *Message) (*MessageStats, error) {
	words := strings.Split(m.Text, " ")
	wc := int32(len(words))
	return &MessageStats{Words: wc}, nil
}

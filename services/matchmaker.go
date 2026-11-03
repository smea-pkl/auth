package services

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Username string
}

type Matchmaker struct {
	mu sync.Mutex
	Queue []*Client
}

func NewMatchmaker() *Matchmaker {
	return  &Matchmaker{
		Queue: make([]*Client, 0),
	}
}

func (m *Matchmaker) FindMatch(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Queue = append(m.Queue, client)
	fmt.Println(client.Username, "bergabung ke antrean matchmaking...")

	if len(m.Queue) >= 2 {
		p1 := m.Queue[0]
		p2 := m.Queue[1]

		m.Queue = m.Queue[2:]

		roomID := uuid.New().String()

		p1.Conn.WriteJSON(map[string]string{
			"type": "MATCH_FOUND",
			"roomID": roomID,
			"opponent": p2.Username,
		})

		p2.Conn.WriteJSON(map[string]string{
			"type": "MATCH_FOUND",
			"roomID": roomID,
			"opponent": p1.Username,
		})

		fmt.Printf("Match found! Room: %s (%s vs %s)\n", roomID, p1.Username, p2.Username)
	}
}

func (m *Matchmaker) RemoveClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.Queue {
		if c == client {
			m.Queue = append(m.Queue[:i], m.Queue[i+1:]...)
			fmt.Println(client.Username, "dihapus dari antrean karena terputus.")
			break
		}
	}
}
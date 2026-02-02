package wsserver

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"crypto-monitor/internal/alerts"

	"github.com/gorilla/websocket"
)

// Server accepts client connections for alerts.
type Server struct {
	addr     string
	upgrader websocket.Upgrader
	mu       sync.Mutex
	clients  map[*websocket.Conn]struct{}
}

// NewServer builds an alert WebSocket server.
func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients: make(map[*websocket.Conn]struct{}),
	}
}

// Run starts the HTTP server and blocks until shutdown.
func (s *Server) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleWS)

	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Broadcast sends alerts to all connected clients.
func (s *Server) Broadcast(ctx context.Context, alertStream <-chan alerts.Alert) {
	for {
		select {
		case <-ctx.Done():
			return
		case alert, ok := <-alertStream:
			if !ok {
				return
			}
			payload, err := json.Marshal(alert)
			if err != nil {
				continue
			}
			s.broadcast(payload)
		}
	}
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	s.addClient(conn)
	go s.readPump(conn)
}

func (s *Server) readPump(conn *websocket.Conn) {
	defer func() {
		s.removeClient(conn)
		_ = conn.Close()
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			return
		}
	}
}

func (s *Server) broadcast(payload []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.clients {
		_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			_ = conn.Close()
			delete(s.clients, conn)
		}
	}
}

func (s *Server) addClient(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[conn] = struct{}{}
}

func (s *Server) removeClient(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, conn)
}

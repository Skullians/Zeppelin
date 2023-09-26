package server

import (
	"net"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	player2 "github.com/aimjel/minecraft/player"
	"github.com/dynamitemc/dynamite/server/network/handlers"
	"github.com/dynamitemc/dynamite/server/player"
)

type Session struct {
	conn *minecraft.Conn

	state *player.Player
}

func New(c *minecraft.Conn, s *player.Player) *Session {
	return &Session{conn: c, state: s}
}

func (s *Session) HandlePackets() error {
	for {
		p, err := s.conn.ReadPacket()
		if err != nil {
			return err
		}

		switch pk := p.(type) {
		case *packet.ChatMessageServer:
			handlers.ChatMessagePacket(pk.Message)
		case *packet.ChatCommandServer:
			handlers.ChatCommandPacket(pk.Command)
		}
		switch p.ID() {
		case 0x14, 0x15, 0x16, 0x17:
			{
				handlers.PlayerMovement(s.state, p)
			}
		}
	}
}

func (s *Session) SendPacket(p packet.Packet) error {

	return s.conn.SendPacket(p)
}

func (s *Session) Info() *player2.Info {
	return s.conn.Info
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}
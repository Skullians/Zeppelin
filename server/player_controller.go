package server

import (
	"slices"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type PlayerController struct {
	player  *player.Player
	session *Session
	Server  *Server

	UUID string
}

func (p *PlayerController) JoinDimension(d *world.Dimension) error {
	if err := p.session.SendPacket(&packet.JoinGame{
		EntityID:           0, //TODO
		IsHardcore:         p.player.IsHardcore(),
		GameMode:           p.player.GameMode(),
		PreviousGameMode:   p.player.PreviousGameMode(),
		DimensionNames:     []string{d.Type()},
		DimensionName:      d.Type(),
		DimensionType:      d.Type(),
		HashedSeed:         d.Seed(),
		ViewDistance:       p.player.ViewDistance(),
		SimulationDistance: p.player.SimulationDistance(),
	}); err != nil {
		return err
	}

	return p.session.SendPacket(&packet.SetDefaultSpawnPosition{})
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.session.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) Position() (x float64, y float64, z float64) {
	return p.player.X, p.player.Y, p.player.Z
}

func (p *PlayerController) Rotation() (yaw float32, pitch float32) {
	return p.player.Yaw, p.player.Pitch
}

func (p *PlayerController) OnGround() bool {
	return p.player.OnGround
}

func (p *PlayerController) SendCommands(graph commands.Graph) {
	for i, command := range graph.Commands {
		if !p.HasPermissions(command.RequiredPermissions) {
			graph.Commands = slices.Delete(graph.Commands, i, i+1)
		}
	}
	p.session.SendPacket(graph.Data())
}

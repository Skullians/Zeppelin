package play

import (
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/slot"
)

// clientbound
const PacketIdSetCreativeModeSlot = 0x32

type SetCreativeModeSlot struct {
	Slot        int16
	ClickedItem slot.Slot
}

func (SetCreativeModeSlot) ID() int32 {
	return PacketIdSetCreativeModeSlot
}

func (s *SetCreativeModeSlot) Encode(w io.Writer) error {
	if err := w.Short(s.Slot); err != nil {
		return err
	}
	return s.ClickedItem.Encode(w)
}

func (s *SetCreativeModeSlot) Decode(r io.Reader) error {
	if err := r.Short(&s.Slot); err != nil {
		return err
	}
	return s.ClickedItem.Decode(r)
}

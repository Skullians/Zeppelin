package dimension

import (
	"fmt"
	"math"
	"os"
	"sync"

	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/region"
)

func NewDimension(regionPath string, typ string, generator region.Generator) *Dimension {
	return &Dimension{
		regions: make(map[uint64]*region.RegionFile),

		regionPath: regionPath,
		typ:        typ,
		generator:  generator,
	}
}

type Dimension struct {
	reg_mu  sync.Mutex
	regions map[uint64]*region.RegionFile

	generator region.Generator

	typ string

	regionPath string
}

func (s *Dimension) Type() string {
	return s.typ
}

func (s *Dimension) GetChunk(x, z int32) (*chunk.Chunk, error) {
	rx, rz := s.chunkPosToRegionPos(x, z)
	region, err := s.getRegion(rx, rz)
	if err != nil {
		if s.generator != nil {
			region = s.newRegion(rx, rz)
		} else {
			return nil, err
		}
	}

	return region.GetChunk(x, z, s.generator)
}

func (s *Dimension) newRegion(rx, rz int32) *region.RegionFile {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	hash := s.regionHash(rx, rz)
	s.regions[hash] = new(region.RegionFile)
	region.EmptyRegion(s.regions[hash])

	return s.regions[hash]
}

func (s *Dimension) getRegion(rx, rz int32) (*region.RegionFile, error) {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	if r, ok := s.regions[s.regionHash(rx, rz)]; ok {
		return r, nil
	}

	reg, err := s.openRegion(rx, rz)
	if err != nil {
		return nil, err
	}

	return reg, err
}

func (s *Dimension) regionHash(rx, rz int32) uint64 {
	return uint64(rx) | uint64(rz)<<32
}

func (s *Dimension) chunkPosToRegionPos(x, z int32) (rx, rz int32) {
	return int32(math.Floor(float64(x) / 32)), int32(math.Floor(float64(z) / 32))
}

func (s *Dimension) openRegion(rx, rz int32) (*region.RegionFile, error) {
	path := fmt.Sprintf("%s/r.%d.%d.mca", s.regionPath, rx, rz)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	hash := s.regionHash(rx, rz)

	s.regions[hash] = new(region.RegionFile)

	err = region.DecodeRegion(file, s.regions[hash])

	return s.regions[hash], err
}
package chunk

type HeightMap struct {
	HeightMaps struct {
		MotionBlocking []int64 `nbt:"MOTION_BLOCKING"`
		WorldSurface   []int64 `nbt:"WORLD_SURFACE"`
	}
}

package util

type Vec3 struct {
	X uint8
	Y uint8
	Z uint8
}

type Vec3Float struct {
	X float32
	Y float32
	Z float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewVec3(x uint8, y uint8, z uint8) Vec3 {
	return Vec3{x, y, z}
}

func NewVec3Float(x float32, y float32, z float32) Vec3Float {
	return Vec3Float{x, y, z}
}

func NewColor(r uint8, g uint8, b uint8) Color {
	return Color{r, g, b}
}

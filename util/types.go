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

func NewVec3() Vec3 {
	return Vec3{0, 0, 0}
}

func NewVec3Float() Vec3Float {
	return Vec3Float{0.0, 0.0, 0.0}
}

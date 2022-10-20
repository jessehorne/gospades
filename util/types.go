package util

type Vec3 struct {
	X uint8
	Y uint8
	Z uint8
}

func NewVec3() Vec3 {
	return Vec3{0, 0, 0}
}

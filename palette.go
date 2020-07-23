package spritesheet32x32x256

import (
	"image/color"
)

type Palette interface {
	color.Model
	Color(uint8) color.Color
	Index(color.Color) uint8
}

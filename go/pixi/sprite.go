package pixi

import (
	"syscall/js"
)

type SpriteModule struct {
	sp js.Value
}

func (pixi *PIXI) Sprite() *SpriteModule {
	return &SpriteModule{
		sp: pixi.Sprite_(),
	}
}

func (sp *SpriteModule) From(path string) *Sprite {
	return &Sprite{
		sp.sp.Call("from", path),
	}
}

type Sprite struct {
	js.Value
}

func (sp *Sprite) SetX(x int) {
	sp.Value.Set("x", x)
}
func (sp *Sprite) SetY(y int) {
	sp.Value.Set("y", y)
}

func (sp *Sprite) Anchor_() js.Value {
	return sp.Value.Get("anchor")
}

func (sp *Sprite) Anchor() *PointData {
	return &PointData{
		sp.Anchor_(),
	}
}

func (sp *Sprite) Rotation_() js.Value {
	return sp.Value.Get("rotation")
}

func (sp *Sprite) Rotation() float64 {
	return sp.Rotation_().Float()
}

func (sp *Sprite) SetRotation(radian float64) {
	sp.Value.Set("rotation", radian)
}

func (sp *Sprite) Scale_() js.Value {
	return sp.Value.Get("scale")
}

func (sp *Sprite) Scale() *PointData {
	return &PointData{
		sp.Scale_(),
	}
}

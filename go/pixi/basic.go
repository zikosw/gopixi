package pixi

import "syscall/js"

type PointData struct {
	js.Value // from .scale, .anchor
}

func (pd *PointData) SetXY(a float64) {
	pd.Value.Call("set", a)
}

func (pd *PointData) Set(x, y float64) {
	pd.Value.Call("set", x, y)
}

// type Number struct {
// 	js.Value // from .rotation
// }

// func (n *Number) Get() float64 {
// 	return n.Value.Get("")
// }

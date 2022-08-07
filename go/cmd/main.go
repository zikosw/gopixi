//go:build js
// +build js

package main

import (
	_ "crypto/sha512"
	"fmt"
	"syscall/js"

	"gopixi/pixi"

	dom "honnef.co/go/js/dom/v2"
)

func main() {
	done := make(chan struct{})
	js.Global().Set("bunnyApp", js.FuncOf(BunnyApp))

	<-done
}

func BunnyApp(this js.Value, args []js.Value) interface{} {

	fmt.Println("bunny app run")

	pixiJS := args[0]

	PX := pixi.PIXI{Value: pixiJS}
	app := PX.App().New()

	view := &dom.BasicNode{Value: app.View_()}
	dom.GetWindow().Document().GetElementsByTagName("body")[0].AppendChild(view)

	bunny := PX.Sprite().From("assets/bunny.png")

	container := PX.Container().New()
	container.AddChild(bunny.Value)

	bunny.Scale().SetXY(2)

	app.Stage().AddChild(container.Value)

	w := app.Screen().Width()
	h := app.Screen().Height()

	bunny.Anchor().SetXY(0.5)
	bunny.SetX(w / 2)
	bunny.SetY(h / 2)

	app.Ticker().Add(func(delta float64) {
		bunny.SetRotation(bunny.Rotation() + 0.01*delta)
	})
	return nil
}

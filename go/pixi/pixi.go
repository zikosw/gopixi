package pixi

import (
	"syscall/js"
)

type PIXI struct {
	js.Value
}

func (px *PIXI) Sprite_() js.Value {
	return px.Value.Get("Sprite")
}

type ApplicationModule struct {
	am js.Value
}

func (px *PIXI) App_() js.Value {
	return px.Value.Get("Application")
}

func (px *PIXI) App() *ApplicationModule {
	return &ApplicationModule{
		am: px.App_(),
	}
}

func (px *PIXI) Container_() js.Value {
	return px.Value.Get("Container")
}

func (px *PIXI) Container() *ContainerModule {
	return &ContainerModule{
		px.Container_(),
	}
}

type Application struct {
	app js.Value
}

func (a *ApplicationModule) New() *Application {
	return &Application{
		app: a.am.New(),
	}
}

func (a *Application) View_() js.Value {
	return a.app.Get("view")
}

func (a *Application) Stage_() js.Value {
	return a.app.Get("stage")
}

func (a *Application) Stage() *Container {
	return &Container{a.Stage_()}
}

type Rectangle struct {
	js.Value
}

func (r *Rectangle) Width() int {
	return r.Value.Get("width").Int()
}
func (r *Rectangle) Height() int {
	return r.Value.Get("height").Int()
}

func (a *Application) Screen_() js.Value {
	return a.app.Get("screen")
}

func (a *Application) Screen() *Rectangle {
	return &Rectangle{a.Screen_()}
}

func (a *Application) Ticker_() js.Value {
	return a.app.Get("ticker")
}

func (a *Application) Ticker() *Ticker {
	return &Ticker{a.Ticker_()}
}

type Ticker struct {
	js.Value
}

func (t *Ticker) Add(callback func(delta float64)) {
	wrapFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		delta := 0.0
		if len(args) > 0 {
			delta = args[0].Float()
		}
		callback(delta)
		return nil
	})
	t.Value.Call("add", wrapFunc)
}

package pixi

import "syscall/js"

type ContainerModule struct {
	js.Value
}

func (cm *ContainerModule) New() *Container {
	return &Container{
		cm.Value.New(),
	}
}

type Container struct {
	js.Value
}

func (c *Container) AddChild(child js.Value) {
	c.Value.Call("addChild", child)
}

func (c *Container) Scale_() js.Value {
	return c.Value.Get("scale")
}

func (c *Container) Scale() *PointData {
	return &PointData{
		c.Scale_(),
	}
}

func (c *Container) Width() int {
	return c.Value.Get("width").Int()
}
func (c *Container) Height() int {
	return c.Value.Get("height").Int()
}

package goback

type Component interface {
}

type Back struct {
	Components []Component
}

func Build() *Back {
	return &Back{}
}

type ArgsConfig struct {
}

func (b *Back) Config() *Back {

	return b
}

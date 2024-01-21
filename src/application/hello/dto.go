package hello

type HelloDto struct {
	Name string `form:"name" default:"User"`
}

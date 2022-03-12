package commandHandler

type Middleware interface {
	Exec(ctx *Context, cmd Command) (next bool, err error)
}

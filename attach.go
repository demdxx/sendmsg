package sendmsg

type Attach interface {
	Name() string
	ContentType() string
	Content() []byte
}

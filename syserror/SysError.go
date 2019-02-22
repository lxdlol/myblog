package syserror

type SysError interface {
	Code()int
	Error()string
	Ression()error
}


func New(msg string,ression error)SysError{
	return UnKnow{msg, ression}
}
package syserror

type Error404 struct {
	UnKnow
}

func(e404 Error404)Code()int{
	return 1001
}
func(e404  Error404)Error()string{
		return "非法访问"
}
func(e404 Error404)Ression()error{
	return e404.ression
}
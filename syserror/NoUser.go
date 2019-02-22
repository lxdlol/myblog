package syserror


type NoUser struct {
	UnKnow
}

func(uk NoUser)Code()int{
	return 1002
}
func(uk NoUser)Error()string{
	return "用户未登录"
}
func(uk NoUser)Ression()error{
	return uk.ression
}
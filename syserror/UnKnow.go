package syserror


type UnKnow struct {
	msg string
	ression error
}


func(uk UnKnow)Code()int{
	return 1000
}
func(uk UnKnow)Error()string{
	if len(uk.msg)==0{
		return "未知错误"
	}
	return uk.msg
}
func(uk UnKnow)Ression()error{
	return uk.ression
}
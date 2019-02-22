package syserror

type HasPriseError struct {
	UnKnow
}

func (e404 HasPriseError) Code() int {
	return 4444
}
func (e404 HasPriseError) Error() string {
	return "已经点赞"
}

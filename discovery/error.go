package discovery

type NotYetDialed struct{}

func (e *NotYetDialed) Error() string {
	return "Need to dial before sending"
}

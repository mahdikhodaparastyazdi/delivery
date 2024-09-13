package couriorproviders

type CouriorSender interface {
	SendCourior(receptor string) error
}

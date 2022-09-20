package httpservice

type Handler struct {
	done chan struct{}
}

package infrastructure

type HandlerLogger struct{}

func (handler *HandlerLogger) Log(message string) error {
	return nil
}

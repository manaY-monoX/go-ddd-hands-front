package errortype

// 内部エラー(サーバに接続できない)
type InternalError struct {
	message string
}

func (i *InternalError) Error() string {
	return i.message
}

func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}

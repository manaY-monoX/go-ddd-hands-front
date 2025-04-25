package errortype

// プレゼンテーション層のエラー型
type PresentationError struct {
	message string
}

func (p *PresentationError) Error() string {
	return p.message
}

func NewPresentationError(message string) *PresentationError {
	return &PresentationError{
		message: message,
	}
}

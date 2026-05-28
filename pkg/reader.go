package pkg

import (
	"io"
)

// Реализация простейшей структуры которая содержит
// полезные данные и метаданные, например длинну строки.
type SimpleNewReader struct {
	Data    string
	DataLen int
}

func NewSimpleReader(s string) *SimpleNewReader {
	return &SimpleNewReader{
		Data:    s,
		DataLen: 0,
	}
}

// Реализация метода интерфейса io.Reader
// Метод CountData(p []byte) читает данные в буфер 'p', продвигая указатель 'dataLen'.
// В результате в структуру записываются соответвующие метаданные.
// Когда данные кончатся — возвращает io.EOF.
func (r *SimpleNewReader) CountData(p []byte) (n int, err error) {
	if r.DataLen >= len(r.Data) {
		return 0, io.EOF
	}

	n = copy(p, []byte(r.Data[r.DataLen:]))
	r.DataLen += n
	return n, nil
}

package wrapper

import "errors"

var (
	w Wrapper
)


type WrapperEncoder func(interface{}) ([]byte, error)

type WrapperDecoder func([]byte, interface{}) error

type Wrapper interface {
	SetBase(interface{})
	GetBase() interface{}

	SetField(string, interface{})
	GetField(string) interface{}
	// Marshal(json.Marshaler)

	SetEncoder(encoder WrapperEncoder)
	SetDecoder(decoder WrapperDecoder)

	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type W struct {
	encoder WrapperEncoder
	decoder WrapperDecoder

	wrapped map[string]interface{}

	base interface{}
}

func (q *W) SetBase(base interface{}) {
	q.wrapped["base"] = base
	q.base = base
}

func (q *W) GetBase() interface{} {
	return q.base
}

func (q *W) SetField(key string, value interface{}) {
	q.wrapped[key] = value
}

func (q *W) GetField(name string) interface{} {
	return q.wrapped[name]
}

func (q *W) SetEncoder(encoder WrapperEncoder) {
	q.encoder = encoder
}

func (q *W) SetDecoder(decoder WrapperDecoder) {
	q.decoder = decoder
}

func (q *W) Marshal() ([]byte, error) {
	if q.encoder == nil {
		return []byte(""), errors.New("encoder is not set")
	}
	return q.encoder(q.wrapped)
}

func (q *W) Unmarshal(src []byte) error {
	if q.decoder == nil {
		return errors.New("decoder is not set")
	}
	if err := q.decoder(src, &q.wrapped); err != nil{
		return err
	}
	
	b, ok := q.wrapped["base"]
	if !ok{
		return errors.New("src does not contain base field")
	}
	q.base = b

	return nil
}

func UseWrapper() Wrapper {
	if w == nil {
		w = &W{wrapped: make(map[string]interface{})}
	}
	return w
}

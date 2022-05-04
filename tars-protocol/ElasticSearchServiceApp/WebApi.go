// Package ElasticSearchServiceApp comment
// This file was generated by tars2go 1.1.7
// Generated from WebApi.tars
package ElasticSearchServiceApp

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// Data struct implement
type Data struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (st *Data) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Data) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Type, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Value, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *Data) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Data, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *Data) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Type, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Value, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Data) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}
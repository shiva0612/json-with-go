// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeShivaTypes(in *jlexer.Lexer, out *Popup) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "menuitem":
			if in.IsNull() {
				in.Skip()
				out.MenuItem = nil
			} else {
				in.Delim('[')
				if out.MenuItem == nil {
					if !in.IsDelim(']') {
						out.MenuItem = make([]MenuItem, 0, 2)
					} else {
						out.MenuItem = []MenuItem{}
					}
				} else {
					out.MenuItem = (out.MenuItem)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MenuItem
					(v1).UnmarshalEasyJSON(in)
					out.MenuItem = append(out.MenuItem, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeShivaTypes(out *jwriter.Writer, in Popup) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"menuitem\":"
		out.RawString(prefix[1:])
		if in.MenuItem == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MenuItem {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Popup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeShivaTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Popup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeShivaTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Popup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeShivaTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Popup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeShivaTypes(l, v)
}
func easyjson6601e8cdDecodeShivaTypes1(in *jlexer.Lexer, out *MenuItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = string(in.String())
		case "onclick":
			out.OnClick = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeShivaTypes1(out *jwriter.Writer, in MenuItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"onclick\":"
		out.RawString(prefix)
		out.String(string(in.OnClick))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MenuItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeShivaTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MenuItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeShivaTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MenuItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeShivaTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MenuItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeShivaTypes1(l, v)
}
func easyjson6601e8cdDecodeShivaTypes2(in *jlexer.Lexer, out *Menu) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "popup":
			(out.Popup).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeShivaTypes2(out *jwriter.Writer, in Menu) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"popup\":"
		out.RawString(prefix)
		(in.Popup).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Menu) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeShivaTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Menu) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeShivaTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Menu) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeShivaTypes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Menu) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeShivaTypes2(l, v)
}

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package generate

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

func easyjsonA0f52c14DecodeGithubComMengdjGoctlRestClientGenerate(in *jlexer.Lexer, out *JSONListResult) {
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
		case "Path":
			out.Path = string(in.String())
		case "Main":
			out.Main = bool(in.Bool())
		case "Dir":
			out.Dir = string(in.String())
		case "GoMod":
			out.GoMod = string(in.String())
		case "GoVersion":
			out.GoVersion = string(in.String())
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
func easyjsonA0f52c14EncodeGithubComMengdjGoctlRestClientGenerate(out *jwriter.Writer, in JSONListResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Path\":"
		out.RawString(prefix[1:])
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"Main\":"
		out.RawString(prefix)
		out.Bool(bool(in.Main))
	}
	{
		const prefix string = ",\"Dir\":"
		out.RawString(prefix)
		out.String(string(in.Dir))
	}
	{
		const prefix string = ",\"GoMod\":"
		out.RawString(prefix)
		out.String(string(in.GoMod))
	}
	{
		const prefix string = ",\"GoVersion\":"
		out.RawString(prefix)
		out.String(string(in.GoVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JSONListResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA0f52c14EncodeGithubComMengdjGoctlRestClientGenerate(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JSONListResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA0f52c14EncodeGithubComMengdjGoctlRestClientGenerate(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JSONListResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA0f52c14DecodeGithubComMengdjGoctlRestClientGenerate(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JSONListResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA0f52c14DecodeGithubComMengdjGoctlRestClientGenerate(l, v)
}

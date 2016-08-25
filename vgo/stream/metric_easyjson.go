// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package stream

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_MetricData(in *jlexer.Lexer, out *MetricData) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "n":
			out.Name = string(in.String())
		case "ts":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tags = make(map[string]string)
				} else {
					out.Tags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Tags)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "f":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Fields = make(map[string]interface{})
				} else {
					out.Fields = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					v2 = in.Interface()
					(out.Fields)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_MetricData(out *jwriter.Writer, in MetricData) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"n\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ts\":")
	if in.Tags == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3_first := true
		for v3_name, v3_value := range in.Tags {
			if !v3_first {
				out.RawByte(',')
			}
			v3_first = false
			out.String(string(v3_name))
			out.RawByte(':')
			out.String(string(v3_value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"f\":")
	if in.Fields == nil {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4_first := true
		for v4_name, v4_value := range in.Fields {
			if !v4_first {
				out.RawByte(',')
			}
			v4_first = false
			out.String(string(v4_name))
			out.RawByte(':')
			out.Raw(json.Marshal(v4_value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"t\":")
	out.Raw((in.Time).MarshalJSON())
	out.RawByte('}')
}
func (v MetricData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_MetricData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v MetricData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_MetricData(w, v)
}
func (v *MetricData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_MetricData(&r, v)
	return r.Error()
}
func (v *MetricData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_MetricData(l, v)
}
func easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_Metrics(in *jlexer.Lexer, out *Metrics) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "d":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Data = make([]*MetricData, 0, 8)
			} else {
				out.Data = nil
			}
			for !in.IsDelim(']') {
				var v5 *MetricData
				if in.IsNull() {
					in.Skip()
					v5 = nil
				} else {
					v5 = new(MetricData)
					(*v5).UnmarshalEasyJSON(in)
				}
				out.Data = append(out.Data, v5)
				in.WantComma()
			}
			in.Delim(']')
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_Metrics(out *jwriter.Writer, in Metrics) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"d\":")
	out.RawByte('[')
	for v6, v7 := range in.Data {
		if v6 > 0 {
			out.RawByte(',')
		}
		if v7 == nil {
			out.RawString("null")
		} else {
			(*v7).MarshalEasyJSON(out)
		}
	}
	out.RawByte(']')
	out.RawByte('}')
}
func (v Metrics) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_Metrics(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v Metrics) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_9478868c_encode_github_com_corego_vgo_vgo_stream_Metrics(w, v)
}
func (v *Metrics) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_Metrics(&r, v)
	return r.Error()
}
func (v *Metrics) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_9478868c_decode_github_com_corego_vgo_vgo_stream_Metrics(l, v)
}

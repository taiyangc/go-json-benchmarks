package go_benchmark

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// encoding/json
func BenchmarkDecodeStdStructSmall(b *testing.B) {
	b.ReportAllocs()
	var data SmallPayload
	for i := 0; i < b.N; i++ {
		json.Unmarshal(smallFixture, &data)
	}
}

func BenchmarkEncodeStdStructSmall(b *testing.B) {
	var data SmallPayload
	json.Unmarshal(smallFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

// json-iterator
func BenchmarkDecodeJsoniterStructSmall(b *testing.B) {
	b.ReportAllocs()
	var data SmallPayload
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(smallFixture, &data)
	}
}

func BenchmarkEncodeJsoniterStructSmall(b *testing.B) {
	var data SmallPayload
	jsoniter.Unmarshal(smallFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(data)
	}
}

// easyjson
func BenchmarkDecodeEasyJsonSmall(b *testing.B) {
	b.ReportAllocs()
	var data SmallPayload
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: smallFixture}
		data.UnmarshalEasyJSON(lexer)
	}
}

func BenchmarkEncodeEasyJsonSmall(b *testing.B) {
	var data SmallPayload
	lexer := &jlexer.Lexer{Data: smallFixture}
	data.UnmarshalEasyJSON(lexer)
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		writer := &jwriter.Writer{}
		data.MarshalEasyJSON(writer)
		buf.Reset()
		writer.DumpTo(buf)
	}
}

// encoding/json
func BenchmarkDecodeStdStructMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		json.Unmarshal(mediumFixture, &data)
	}
}

func BenchmarkEncodeStdStructMedium(b *testing.B) {
	var data MediumPayload
	json.Unmarshal(mediumFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

// json-iterator
func BenchmarkDecodeJsoniterStructMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(mediumFixture, &data)
	}
}

func BenchmarkEncodeJsoniterStructMedium(b *testing.B) {
	var data MediumPayload
	jsoniter.Unmarshal(mediumFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(data)
	}
}

// easyjson
func BenchmarkDecodeEasyJsonMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: mediumFixture}
		data.UnmarshalEasyJSON(lexer)
	}
}

func BenchmarkEncodeEasyJsonMedium(b *testing.B) {
	var data MediumPayload
	lexer := &jlexer.Lexer{Data: mediumFixture}
	data.UnmarshalEasyJSON(lexer)
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		writer := &jwriter.Writer{}
		data.MarshalEasyJSON(writer)
		buf.Reset()
		writer.DumpTo(buf)
	}
}

// encoding/json
func BenchmarkDecodeStdStructLarge(b *testing.B) {
	b.ReportAllocs()
	var data LargePayload
	for i := 0; i < b.N; i++ {
		json.Unmarshal(largeFixture, &data)
	}
}

func BenchmarkEncodeStdStructLarge(b *testing.B) {
	var data LargePayload
	json.Unmarshal(largeFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

// json-iterator
func BenchmarkDecodeJsoniterStructLarge(b *testing.B) {
	b.ReportAllocs()
	var data LargePayload
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(largeFixture, &data)
	}
}

func BenchmarkEncodeJsoniterStructLarge(b *testing.B) {
	var data LargePayload
	jsoniter.Unmarshal(largeFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(data)
	}
}

// easyjson
/* Unsupported
func BenchmarkDecodeEasyJsonLarge(b *testing.B) {
	b.ReportAllocs()
	var data LargePayload
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: largeFixture}
		data.UnmarshalEasyJSON(lexer)
	}
}

func BenchmarkEncodeEasyJsonLarge(b *testing.B) {
	var data LargePayload
	lexer := &jlexer.Lexer{Data: largeFixture}
	data.UnmarshalEasyJSON(lexer)
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		writer := &jwriter.Writer{}
		data.MarshalEasyJSON(writer)
		buf.Reset()
		writer.DumpTo(buf)
	}
}
*/

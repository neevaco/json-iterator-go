package test

import (
	"bytes"
	"encoding/json"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"testing"
)

func Test_disallowUnknownFields(t *testing.T) {
	should := require.New(t)
	type TestObject struct{}
	var obj TestObject
	decoder := jsoniter.NewDecoder(bytes.NewBufferString(`{"field1":100}`))
	decoder.DisallowUnknownFields()
	should.Error(decoder.Decode(&obj))
}

func Test_new_decoder(t *testing.T) {
	should := require.New(t)
	{
		decoder := json.NewDecoder(bytes.NewBufferString(`[1][2]`))
		arr := []int{}
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{1}, arr)
		should.True(decoder.More())
		buffered, _ := ioutil.ReadAll(decoder.Buffered())
		should.Equal("[2]", string(buffered))
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{2}, arr)
		should.False(decoder.More())
	}
	{
		decoder := jsoniter.NewDecoder(bytes.NewBufferString(`[1][2]`))
		arr := []int{}
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{1}, arr)
		should.True(decoder.More())
		buffered, _ := ioutil.ReadAll(decoder.Buffered())
		should.Equal("[2]", string(buffered))
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{2}, arr)
		should.False(decoder.More())
	}
}

func Test_new_decoder_whitespace(t *testing.T) {
	should := require.New(t)
	{
		decoder := json.NewDecoder(bytes.NewBufferString(`  [1]  [2]  `))
		arr := []int{}
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{1}, arr)
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{2}, arr)
		should.Equal(io.EOF, decoder.Decode(&arr))
	}
	{
		decoder := jsoniter.NewDecoder(bytes.NewBufferString(`  [1]  [2]  `))
		arr := []int{}
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{1}, arr)
		should.Nil(decoder.Decode(&arr))
		should.Equal([]int{2}, arr)
		should.Equal(io.EOF, decoder.Decode(&arr))
	}
}

func Test_use_number(t *testing.T) {
	should := require.New(t)
	decoder1 := json.NewDecoder(bytes.NewBufferString(`123`))
	decoder1.UseNumber()
	decoder2 := jsoniter.NewDecoder(bytes.NewBufferString(`123`))
	decoder2.UseNumber()
	var obj1 interface{}
	should.Nil(decoder1.Decode(&obj1))
	should.Equal(json.Number("123"), obj1)
	var obj2 interface{}
	should.Nil(decoder2.Decode(&obj2))
	should.Equal(json.Number("123"), obj2)
}

func Test_decoder_more(t *testing.T) {
	should := require.New(t)
	decoder := jsoniter.NewDecoder(bytes.NewBufferString("abcde"))
	should.True(decoder.More())
}

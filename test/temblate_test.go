package temblate

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

type tdata struct {
	Nickname string
}

func TestGetMessage(t *testing.T) {
	type args struct {
		lang string
		key  string
		data interface{}
	}

	dat := &tdata{Nickname: "Devvy"}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "msg1en", args: args{lang: "en", key: "msg1", data: dat}, want: "Just a simple message."},
		{name: "msg2en", args: args{lang: "en", key: "msg2", data: dat}, want: `Your nickname is Devvy!

However, we now add some symbols:

Sunnycloud: 🌦
Rocket: 🚀
Devil: 😈
`,
		},
		{name: "msg1langnoexist", args: args{lang: "de", key: "msg1", data: dat}, want: "Just a simple message."},
		{name: "msg1nl", args: args{lang: "nl", key: "msg1", data: dat}, want: "Gewoon een simpele mededeling."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(tt.args.lang, tt.args.key, tt.args.data); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

// this is a kind of trivial test to check possible UTF encoding issues when handling []byte vs. string
// The observed issue (with an email sending lib) could not be reproduced though, and is probably caused by the email encoding in that lib
func TestFileEnc(t *testing.T) {
	dat := &tdata{Nickname: "Devvy"}
	fn := ".tstfile.txt"
	msg := GetMessage("en", "msg2", &dat)
	if err := ioutil.WriteFile(fn, []byte(msg), 0640); err != nil {
		t.Errorf("ERROR writing file: %q", fn)
	}
	t.Logf("msg: %q", msg)
	msgr, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Errorf("ERROR reading file: %q", fn)
	}
	if !bytes.Equal([]byte(msg), msgr) {
		t.Errorf("ERROR: conversion not identical: %q != %q", msg, msgr)
	}
	_ = os.Remove(fn)
}

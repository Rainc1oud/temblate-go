package temblate_go

import "testing"

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
		{name: "msg2en", args: args{lang: "en", key: "msg2", data: dat}, want: "Your nickname is Devvy!\n"},
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

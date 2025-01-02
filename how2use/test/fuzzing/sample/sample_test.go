package sample

import "testing"

func FuzzDoSomething(f *testing.F) {
	f.Add(int16(0), "")
	f.Add(int16(1), "test")
	f.Add(int16(1000), "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	f.Add(int16(10000), "てすと")
	f.Fuzz(func(t *testing.T, id int16, name string) {
		doSomething(id, name)
	})
}

//func Test_doSomething(t *testing.T) {
//	type args struct {
//		id   int
//		name string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{
//			name: "1",
//			args: args{id: 1, name: "test"},
//			want: "{id:1, name:test}",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := doSomething(tt.args.id, tt.args.name); got != tt.want {
//				t.Errorf("doSomething() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

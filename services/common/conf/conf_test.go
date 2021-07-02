package conf

import "testing"

func TestCONF(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success", args{"TEST"}, "TEST"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CONF(tt.args.key); got != tt.want {
				t.Errorf("CONF() = %v, want %v", got, tt.want)
			}
		})
	}
}

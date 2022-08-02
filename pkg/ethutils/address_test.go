package ethutils

import "testing"

func TestIsStringAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid address",
			args: args{address: "0x96216849c49358B10257cb55b28eA603c874b05E"},
			want: true,
		},
		{
			name: "invalid address",
			args: args{address: "0x96216849c49358B10257cb55b28eA603c874b05g"},
			want: false,
		},
		{
			name: "invalid address",
			args: args{address: "0x96216849c49358B10257cb55b28eA603c874b0"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringAddress(tt.args.address); got != tt.want {
				t.Errorf("IsStringAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

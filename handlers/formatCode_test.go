package handlers

import "testing"

func TestFormatCode(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				dir: "./files",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FormatCode(tt.args.dir)
		})
	}
}

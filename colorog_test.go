package colorog

import (
	"testing"

	"github.com/wander4747/colorog/color"
)

func TestColorog_All(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all colors",
			args: args{
				text: "test colorful",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Success(tt.args.text)
			Info(tt.args.text)
			Warning(tt.args.text)
			Light(tt.args.text)
			Danger(tt.args.text)
			WithColor(color.Green, tt.args.text)
			Unicorn(tt.args.text)
		})
	}
}

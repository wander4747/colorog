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
			c := New()
			c.Success(tt.args.text)
			c.Info(tt.args.text)
			c.Warning(tt.args.text)
			c.Light(tt.args.text)
			c.Danger(tt.args.text)
			c.WithColor(color.Green, tt.args.text)
		})
	}
}

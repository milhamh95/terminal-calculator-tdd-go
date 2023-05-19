package table_test

import (
	"github.com/stretchr/testify/require"
	"terminal-calculator-tdd-go/table"
	"testing"
)

func TestDivide(t *testing.T) {
	type args struct {
		x int8
		y int8
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "pos x, pos y",
			args: args{
				x: 8,
				y: 4,
			},
			want:    "2.00",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			x, y := int8(tt.args.x), int8(tt.args.y)
			got, err := table.Divide(x, y)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}

			require.Nil(t, err)
			require.Equal(t, tt.want, *got)
		})
	}
}

package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCode(t *testing.T) {
	type args struct {
		id      uint64
		options []func(*CodeOptions)
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				id: 1,
			},
			want: "H84VXB6J",
		},
		{
			name: "with-options",
			args: args{
				id: 1,
				options: []func(*CodeOptions){
					WithCodeChars([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}),
					WithCodeN1(9),
					WithCodeN2(3),
					WithCodeL(5),
					WithCodeSalt(56789),
				},
			},
			want: "07873",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewCode(tt.args.id, tt.args.options...))
		})
	}
}

func BenchmarkNewCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCode(1)
	}
}

func BenchmarkNewCodeTimeConsuming(b *testing.B) {
	b.StopTimer()

	id := NewCode(1)
	assert.Equal(b, "H84VXB6J", id)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		NewCode(1)
	}
}

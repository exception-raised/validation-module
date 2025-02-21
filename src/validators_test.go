package validator

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "test@test.com",
			},
			want: true,
		},
		{
			name: "invalid email",
			args: args{
				email: "test@test",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestIsPasswordValid(t *testing.T) {
	type args struct {
		password  string
		minLength int
		maxLength int
		regex     string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid password",
			args: args{
				password:  "testtest",
				minLength: 8,
				maxLength: 16,
				regex:     "",
			},
			want: true,
		},
		{
			name: "invalid password",
			args: args{
				password:  "",
				minLength: 8,
				maxLength: 16,
				regex:     "",
			},
			want: false,
		},
		{
			name: "valid password with symbols and digits",
			args: args{
				password:  "Test123!",
				minLength: 8,
				maxLength: 16,
				regex:     validateStrictPasswordPattern,
			},
			want: true,
		},
		{
			name: "invalid password with symbols and digits",
			args: args{
				password:  "test123",
				minLength: 8,
				maxLength: 16,
				regex:     validateStrictPasswordPattern,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPassword(tt.args.password, tt.args.minLength, tt.args.maxLength, tt.args.regex); got != tt.want {
				t.Errorf("IsValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

package google_cloud

import "testing"

func TestFirebaseAuthIDToken_Email(t *testing.T) {
	tests := []struct {
		name  string
		t     FirebaseAuthIDToken
		want  string
		want1 bool
	}{
		{
			name: "email claim が存在する => メールアドレスを返す",
			t: FirebaseAuthIDToken{
				Claims: map[string]interface{}{
					"email": "test@example.com",
				},
			},
			want:  "test@example.com",
			want1: true,
		},
		{
			name: "email claim が存在しない => 空文字を返す",
			t: FirebaseAuthIDToken{
				Claims: map[string]interface{}{},
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.t.Email()
			if got != tt.want {
				t.Errorf("Email() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Email() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

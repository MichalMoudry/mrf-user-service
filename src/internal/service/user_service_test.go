package service

import (
	"testing"
)

// Test fixture covering tests related to user ID validation.
func Test_UserIdValidation(t *testing.T) {
	type args struct {
		token   string
		wantErr bool
	}
	tests := []struct {
		name string
		args
		shouldSkip bool
	}{
		{
			name: "Basic token validation",
			args: args{
				token:   "",
				wantErr: false,
			},
			shouldSkip: true,
		},
	}

	/*creds := `
	`
	opt := option.WithCredentialsJSON([]byte(creds))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config.GetFirebaseConfig(), opt)
	if err != nil {
		t.Error(err)
		return
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	*/
	for _, test := range tests {
		if !test.shouldSkip {
			t.Run(test.name, func(t *testing.T) {
				/*token, err := auth.VerifyIDToken(ctx, test.token)
				if err != nil {
					t.Error(err)
					return
				}
				t.Log(token.Subject)*/
			})
		} else {
			t.Skip("Test is marked to be skipped.")
		}
	}
}

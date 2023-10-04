package service

import (
	"context"
	"os"
	"testing"
	"user-service/internal/config"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
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

	os.Setenv("FB_TYPE", "")
	os.Setenv("FB_PRIV_KEY_ID", "")
	os.Setenv("FB_PRIV_KEY", "")
	os.Setenv("FB_CLIENT_EMAIL", "")
	os.Setenv("FB_CLIENT_ID", "")
	os.Setenv("FB_AUTH_URI", "")
	os.Setenv("FB_TOKEN_URI", "")
	os.Setenv("FB_CLIENT_CERT_URL", "")
	creds, err := config.CreateFirebaseCredentials()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config.GetFirebaseConfig(), option.WithCredentialsJSON([]byte(creds)))
	if err != nil {
		//t.Error(err)
		return
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		//t.Error(err)
		return
	}

	for _, test := range tests {
		if !test.shouldSkip {
			t.Run(test.name, func(t *testing.T) {
				token, err := auth.VerifyIDToken(ctx, test.token)
				if err != nil {
					t.Error(err)
					return
				}
				t.Log(token.Subject)
			})
		} else {
			t.Skip("Test is marked to be skipped.")
		}
	}
}

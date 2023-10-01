package service

import (
	"context"
	"os"
	"testing"
	"user-service/internal/config"

	firebase "firebase.google.com/go/v4"
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
	}{
		{
			name: "Basic token validation",
			args: args{
				token:   "",
				wantErr: false,
			},
		},
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./ocr-microservice-project.json")
	ctx := context.Background()
	_, err := firebase.NewApp(ctx, config.GetFirebaseConfig())
	if err != nil {
		t.Error(err)
		return
	}
	/*_, err = app.Auth(ctx)
	if err != nil {
		t.Error(err)
		return
	}*/

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			/*token, err := auth.VerifyIDToken(ctx, test.token)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(token.Subject)*/
		})
	}
}

package config

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_PrivateKeyFormatting(t *testing.T) {
	type args struct {
		privKey string
		wantErr bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test basic string replacement in private key",
			args: args{
				privKey: "-----BEGIN PRIVATE KEY-----\\nwfjwjfowfkwepkf\\n",
				wantErr: false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Setenv("FB_PRIV_KEY", test.args.privKey)
			data, err := CreateFirebaseCredentials()
			if err != nil {
				t.Error(err)
				return
			}

			credentials := GoogleCredentials{}
			err = json.Unmarshal(data, &credentials)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(credentials)
		})
	}
}

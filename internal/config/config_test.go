package config

import (
	"reflect"
	"testing"
	"time"
)

// TODO смотри в config.go и main.yml
func TestInit(t *testing.T) {

	type args struct {
		path    string
		pathEnv string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path:    "../fixtures/test",
				pathEnv: "../fixtures/etest",
			},
			want: &Config{
				LoggerLevel: 7,
				CacheTTL:    time.Minute * 777,
				HTTP: HTTPConfig{
					Port:         "7777",
					ReadTimeout:  time.Second * 77,
					WriteTimeout: time.Second * 77,
				},
				Auth: AuthConfig{
					JWT: JWTConfig{
						AccessTokenTTL:  time.Minute * 7,
						RefreshTokenTTL: time.Minute * 777,
						SigningKey:      "test_key",
					},
					PasswordSalt: "test_salt",
				},
				Mongo: MongoConfig{
					DatabaseName: "testDatabase",
					URI:          "mongodb://localhost:27017/test",
					User:         "Tester",
					Password:     "test",
				},
				Email: EmailConfig{
					ClientSecret: "test_secret",
					ClientID:     "321",
					ListID:       "666",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.path, tt.args.pathEnv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, want %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}

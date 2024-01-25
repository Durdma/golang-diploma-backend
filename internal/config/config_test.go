package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// TODO смотри в config.go и main.yml
func TestInit(t *testing.T) {
	type env struct {
		mongoURI      string
		mongoUser     string
		mongoPass     string
		passwordSalt  string
		jwtSigningKey string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("MONGO_URI", env.mongoURI)
		os.Setenv("MONGO_USER", env.mongoUser)
		os.Setenv("MONGO_PASS", env.mongoPass)
		os.Setenv("PASSWORD_SALT", env.passwordSalt)
		os.Setenv("JWT_SIGNING_KEY", env.jwtSigningKey)
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
				path: "D:\\Projects\\University\\LAST_COURSE\\ДП\\go-saas\\internal\\fixtures\\test.yml",
				env: env{
					mongoURI:      "mongodb://localhost:27017",
					mongoUser:     "User",
					mongoPass:     "qwerty",
					passwordSalt:  "salt",
					jwtSigningKey: "key",
				},
			},
			want: &Config{
				LoggerLevel: 5,
				HTTP: HTTPConfig{
					Port:         "8080",
					ReadTimeout:  time.Second * 10,
					WriteTimeout: time.Second * 10,
				},
				Auth: AuthConfig{
					JWT: JWTConfig{
						AccessTokenTTL:  time.Minute * 10,
						RefreshTokenTTL: time.Minute * 15,
						SigningKey:      "key",
					},
					PasswordSalt: "salt",
				},
				Mongo: MongoConfig{
					Name:     "testDatabase",
					URI:      "mongodb://localhost:27017",
					User:     "User",
					Password: "qwerty",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := Init(tt.args.path)
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

package configuration

import (
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

func TestDBDonfig_GetPostgresDsn(t *testing.T) {
	type fields struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
	}
	tests := map[string]struct {
		fields fields
		want   string
	}{
		"success": {
			fields: fields{
				DBHost:     "test",
				DBPort:     "test",
				DBUser:     "test",
				DBPassword: "test",
				DBName:     "test",
			},
			want: "host=test port=test user=test password=test dbname=test sslmode=disable",
		},
	}
	for caseName, tt := range tests {
		t.Run(caseName, func(t *testing.T) {
			c := DBDonfig{
				DBHost:     tt.fields.DBHost,
				DBPort:     tt.fields.DBPort,
				DBUser:     tt.fields.DBUser,
				DBPassword: tt.fields.DBPassword,
				DBName:     tt.fields.DBName,
			}
			got := c.GetPostgresDsn()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNew(t *testing.T) {
	tests := map[string]struct {
		setEnvFf  func()
		dropEnvFn func()
		want      *Configuration
		wantErr   bool
	}{
		"error.parsing envs": {
			setEnvFf:  func() {},
			dropEnvFn: func() {},
			want:      nil,
			wantErr:   true,
		},
		"success": {
			setEnvFf: func() {
				os.Setenv("DB_HOST", "postgres")
				os.Setenv("DB_PORT", "5432")
				os.Setenv("DB_USER", "postgres")
				os.Setenv("DB_PASSWORD", "postgres")
				os.Setenv("DB_NAME", "test")
				os.Setenv("APPLICATION_PORT", "8080")
			},

			dropEnvFn: func() {
				os.Unsetenv("DB_HOST")
				os.Unsetenv("DB_PORT")
				os.Unsetenv("DB_USER")
				os.Unsetenv("DB_PASSWORD")
				os.Unsetenv("DB_NAME")
				os.Unsetenv("APPLICATION_PORT")
			},
			want: &Configuration{
				ApplicationPort: "8080",
				DBDonfig: DBDonfig{
					DBHost:     "postgres",
					DBPort:     "5432",
					DBUser:     "postgres",
					DBPassword: "postgres",
					DBName:     "test",
				},
			},
			wantErr: false,
		},
	}
	for caseName, tt := range tests {
		t.Run(caseName, func(t *testing.T) {
			tt.setEnvFf()

			got, err := New()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)

			tt.dropEnvFn()
		})
	}
}

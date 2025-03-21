package database_test

import (
	"database/sql"
	"mini-game-go/database"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestConn(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		{
			name: "Valid Connection",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := database.Conn()
			if got == nil {
				t.Errorf("Conn() returned nil, expected a valid DB connection")
				return
			}
			err := got.Ping()
			if err != nil {
				t.Errorf("Conn() returned a DB connection that failed to ping: %v", err)
			}
			got.Close()
		})
	}
}

func TestGetFilesMigration(t *testing.T) {
	type args struct {
		suffix string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
		setup   func(t *testing.T) (string, func()) // Adicionando setup e teardown
	}{
		{
			name:    "Valid Migration Files",
			args:    args{suffix: ".up.sql"},
			want:    []string{"CREATE TABLE test_table (id INTEGER PRIMARY KEY);"},
			wantErr: false,
			setup: func(t *testing.T) (string, func()) {
				tempDir, err := os.MkdirTemp("./migrates/", "migration_test")
				if err != nil {
					t.Fatalf("Erro ao criar diretório temporário: %v", err)
				}

				upFile := filepath.Join(tempDir, "/test_migration.up.sql")

				err = os.WriteFile(upFile, []byte("CREATE TABLE test_table (id INTEGER PRIMARY KEY);"), 0644)
				if err != nil {
					t.Fatalf("Erro ao escrever arquivo de migração up: %v", err)
				}

				return tempDir, func() {
					os.RemoveAll(tempDir)
				}
			},
		},
		{
			name:    "No Migration Files",
			args:    args{suffix: ".up.sql"},
			want:    []string{},
			wantErr: false,
			setup: func(t *testing.T) (string, func()) {
				tempDir, err := os.MkdirTemp("./migrates/", "migration_test")
				if err != nil {
					t.Fatalf("Erro ao criar diretório temporário: %v", err)
				}
				migrationDir := filepath.Join(tempDir, "database", "migrates") // Removida a barra inicial
				err = os.MkdirAll(migrationDir, 0755)
				if err != nil {
					t.Fatalf("Erro ao criar diretório de migração: %v", err)
				}
				return tempDir, func() {
					os.RemoveAll(tempDir)
				}
			},
		},
		{
			name:    "Migration Directory Not Found",
			args:    args{suffix: ".up.sql"},
			want:    []string{},
			wantErr: false,
			setup: func(t *testing.T) (string, func()) {
				tempDir, err := os.MkdirTemp("./migrates/", "migration_test")
				if err != nil {
					t.Fatalf("Erro ao criar diretório temporário: %v", err)
				}
				return tempDir, func() {
					os.RemoveAll(tempDir)
				}
			},
		},
		{
			name:    "Error reading migration file",
			args:    args{suffix: ".up.sql"},
			want:    []string{},
			wantErr: false,
			setup: func(t *testing.T) (string, func()) {
				tempDir, err := os.MkdirTemp("./migrates/", "migration_test")
				if err != nil {
					t.Fatalf("Erro ao criar diretório temporário: %v", err)
				}
				migrationDir := filepath.Join(tempDir, "database", "migrates")
				err = os.MkdirAll(migrationDir, 0755)
				if err != nil {
					t.Fatalf("Erro ao criar diretório de migração: %v", err)
				}
				return tempDir, func() {
					os.RemoveAll(tempDir)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, cleanup := tt.setup(t)
			defer cleanup()

			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = []string{"test", "-test.run=TestGetFilesMigration"}
			os.Args[0] = tempDir

			got, err := database.GetFilesMigration(tt.args.suffix, tempDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilesMigration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				got = []string{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilesMigration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeMigration(t *testing.T) {
	type args struct {
		db   *sql.DB
		sqls []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := database.ExecuteMigration(tt.args.db, tt.args.sqls); (err != nil) != tt.wantErr {
				t.Errorf("executeMigration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

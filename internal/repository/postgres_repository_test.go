package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonikq/gophkeeper/internal/app/server/config"
	"reflect"
	"testing"
)

func TestNewPostgresRepository(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg config.PostgresConfig
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostgresRepository(tt.args.ctx, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRepository(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg config.PostgresConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *pgxpool.Pool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newRepository(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRepository() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postgresRepository_Close(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepository{
				pool: tt.fields.pool,
			}
			r.Close()
		})
	}
}

func Test_postgresRepository_Ping(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepository{
				pool: tt.fields.pool,
			}
			if err := r.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

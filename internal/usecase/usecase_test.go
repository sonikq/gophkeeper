package usecase

import (
	"context"
	"github.com/sonikq/gophkeeper/internal/repository"
	"reflect"
	"testing"
)

func TestNewGophKeeperUseCase(t *testing.T) {
	type args struct {
		r repository.Repository
	}
	tests := []struct {
		name string
		args args
		want GophKeeperUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGophKeeperUseCase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGophKeeperUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gophKeeperUseCase_Ping(t *testing.T) {
	type fields struct {
		repo repository.Repository
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
			g := &gophKeeperUseCase{
				repo: tt.fields.repo,
			}
			if err := g.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

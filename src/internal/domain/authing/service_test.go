package authing_test

import (
	"context"
	"testing"
	"time"

	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/authing"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"coryptex.com/bot/vip-signal/internal/domain/mocks"
	"github.com/stretchr/testify/require"
)

//func TestNewService(t *testing.T) {
//	type args struct {
//		repo domain.AdminRepository
//	}
//	tests := []struct {
//		name string
//		args args
//		want Service
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

var (
	authRepoMock = new(mocks.AdminRepository)
)

func Test_service_Authorize(t *testing.T) {
	type fields struct {
		repo domain.AdminRepository
	}
	type args struct {
		ctx context.Context
		tid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test Authorize (Allowed)",
			fields: fields{
				repo: authRepoMock,
			},
			args: args{
				ctx: context.TODO(),
				tid: "21126556",
			},
			want:    true, // Is Authorized?
			wantErr: false,
		},
		{
			name: "Test Authorize (Not Allowed)",
			fields: fields{
				repo: authRepoMock,
			},
			args: args{
				ctx: context.TODO(),
				tid: "21126556",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		authRepoMock.On("GetByTID", tt.args.ctx, tt.args.tid).Return(
			func(ctx context.Context, tid string) *entities.Admin {
				return &entities.Admin{
					ID:        "",
					TID:       tt.args.tid,
					Active:    tt.want,
					CreatedAt: time.Now(),
				}
			},
			func(ctx context.Context, tid string) error {
				return nil
			})

		svc := authing.NewService(authRepoMock)
		got, err := svc.Authorize(tt.args.ctx, tt.args.tid)
		if err != nil {
			require.EqualValues(t, tt.wantErr, err)
		}
		require.EqualValues(t, tt.want, got)
	}
}

func Test_service_NewAdmin(t *testing.T) {
	type fields struct {
		repo domain.AdminRepository
	}
	type args struct {
		ctx context.Context
		tid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test New Admin",
			fields: fields{
				repo: authRepoMock,
			},
			args: args{
				ctx: context.TODO(),
				tid: "2354446",
			},
			wantErr: false,
		},
	}
	// FIXME: test failed, fix then refactor (use assert/require pkg)
	for _, tt := range tests {
		authRepoMock.On("Add", tt.args.ctx, entities.Admin{
			ID:        "",
			TID:       tt.args.tid,
			CreatedAt: time.Now(),
			Active:    true,
		}).Return("", nil)
		svc := authing.NewService(authRepoMock)
		err := svc.NewAdmin(tt.args.ctx, tt.args.tid)
		if err != nil {
			require.EqualValues(t, tt.wantErr, err)
		}
	}
}

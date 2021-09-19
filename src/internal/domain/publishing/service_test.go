package publishing_test

import (
	"context"
	"coryptex.com/bot/vip-signal/internal/domain"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"coryptex.com/bot/vip-signal/internal/domain/mocks"
	"coryptex.com/bot/vip-signal/internal/domain/publishing"
	"github.com/pkg/errors"
	"testing"
)

//func TestNewService(t *testing.T) {
//	type args struct {
//		repo domain.SignalRepository
//		pub  domain.SignalPublisher
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
//			if got := NewService(tt.args.repo, tt.args.pub); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

var (
	signalRepoMock      = new(mocks.SignalRepository)
	signalPublisherMock = new(mocks.SignalPublisher)
)

func Test_service_Publish(t *testing.T) {
	//asrt := assert.New(t)
	type fields struct {
		repo domain.SignalRepository
		pub  domain.SignalPublisher
	}
	type args struct {
		ctx context.Context
		ss  []entities.Signal
	}
	ss := []entities.Signal{
		{
			ID:          "22",
			Pair:        "BTC/USDT",
			Date:        "25/2/2021",
			ImageURL:    "http://something.sth",
			TimeFrame:   "1H",
			EntryPrice:  "30000",
			TargetPrice: "20000",
			StopLoss:    "31000",
			Risk2Reward: "10%",
			TradeVolume: "30000",
		},
		//{
		//	ID:          "23",
		//	Pair:        "ETH/USDT",
		//	Date:        "17/5/2021",
		//	ImageURL:    "http://something.sth",
		//	TimeFrame:   "1M",
		//	EntryPrice:  "1500",
		//	TargetPrice: "4000",
		//	StopLoss:    "1000",
		//	Risk2Reward: "29%",
		//	TradeVolume: "63000",
		//},
		//{
		//	ID:          "56",
		//	Pair:        "TRX/USDT",
		//	Date:        "6/4/2020",
		//	ImageURL:    "http://something.sth",
		//	TimeFrame:   "1m",
		//	EntryPrice:  "0.055",
		//	TargetPrice: "0.9",
		//	StopLoss:    "0.05",
		//	Risk2Reward: "77%",
		//	TradeVolume: "8700",
		//},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "service_publish",
			fields: fields{
				repo: signalRepoMock,
				pub:  signalPublisherMock,
			},
			args: args{
				ctx: context.Background(),
				ss:  ss,
			},
			wantErr: false,
		},
		//{
		//	name: "service_publish_error",
		//	fields: fields{
		//		repo: signalRepoMock,
		//		pub:  signalPublisherMock,
		//	},
		//	args: args{
		//		ctx: context.Background(),
		//		ss:  ss,
		//	},
		//	wantErr: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rErr error
			if tt.wantErr {
				rErr = errors.New("Error")
			} else {
				rErr = nil
			}
			signalRepoMock.On("AddMany", tt.args.ctx, tt.args.ss).Return("", rErr)
			for _, s := range tt.args.ss {
				signalPublisherMock.On("Publish", tt.args.ctx, s).Return(rErr)
			}
			svc := publishing.NewService(tt.fields.repo, tt.fields.pub)
			if err := svc.Publish(tt.args.ctx, tt.args.ss); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

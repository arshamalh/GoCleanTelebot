package mongo_test

import (
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"context"
	"coryptex.com/bot/vip-signal/internal/data/repositories/mocks"
	"testing"
)

var signalDBMock = new(mocks.SignalDB)

func Test_mongo_InsertMany(t *testing.T) {
	type fields struct {
		db *mocks.SignalDB
	}
	type args struct {
		ctx  context.Context
		sigs []entities.Signal
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
		{
			ID:          "23",
			Pair:        "ETH/USDT",
			Date:        "17/5/2021",
			ImageURL:    "http://something.sth",
			TimeFrame:   "1M",
			EntryPrice:  "1500",
			TargetPrice: "4000",
			StopLoss:    "1000",
			Risk2Reward: "29%",
			TradeVolume: "63000",
		},
		{
			ID:          "56",
			Pair:        "TRX/USDT",
			Date:        "6/4/2020",
			ImageURL:    "http://something.sth",
			TimeFrame:   "1m",
			EntryPrice:  "0.055",
			TargetPrice: "0.9",
			StopLoss:    "0.05",
			Risk2Reward: "77%",
			TradeVolume: "8700",
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test Insert Many Items",
			fields:  fields{
				db: signalDBMock,
			},
			args:    args{
				ctx:  context.Background(),
				sigs: ss,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signalDBMock.On("InsertMany", tt.args.ctx, tt.args.sigs).Return("", nil)
			got, err := tt.fields.db.InsertMany(tt.args.ctx, tt.args.sigs)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertMany() got = %v, want %v", got, tt.want)
			}
		})
	}
}
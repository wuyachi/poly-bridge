package models

import (
	"poly-bridge/conf"
	"testing"
)

func TestGetL1BlockNumberOfArbitrumTx(t *testing.T) {
	type args struct {
		hash string
	}
	test := struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		name: "TestGetL1BlockNumberOfArbitrumTx",
		args: args{
			"ec72ffe0758dc05c03a825b0c78d9e9a22a777dfd1e035adfbdeec25bb62f026",
		},
		want: 13408988,
	}
	config := conf.NewConfig("./../conf/config_mainnet.json")
	conf.GlobalConfig = config

	t.Run(test.name, func(t *testing.T) {
		got, err := GetL1BlockNumberOfArbitrumTx(test.args.hash)
		if (err != nil) != test.wantErr {
			t.Errorf("GetL1BlockNumberOfArbitrumTx() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if got != test.want {
			t.Errorf("GetL1BlockNumberOfArbitrumTx() got = %v, want %v", got, test.want)
		}
	})
}

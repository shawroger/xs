package config

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "文件测试1",
			args: args{filename: "../testdata/config.parse_test.01.json"},
			want: &Config{Port: 8080, Debug: true, GinDebug: false, Files: []FileConfig{
				{
					File:   "./testdata/xlsx/test-01.xlsx",
					Prefix: "test",
					Sheets: nil,
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

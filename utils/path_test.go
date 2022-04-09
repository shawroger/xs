package utils

import "testing"

func TestUnifyPathFormat(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "混合路径1",
			args: args{filename: "A:\\B/C\\D///E"},
			want: "A:/B/C/D/E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnifyPathFormat(tt.args.filename); got != tt.want {
				t.Errorf("UnifyPathFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPureName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "纯英文路径1",
			args: args{filename: "./ABC/D//EF/G.txt"},
			want: "G",
		},
		{
			name: "纯英文路径2",
			args: args{filename: "./ABC/D//EF/TEST_IT_IS.txt"},
			want: "TEST_IT_IS",
		},
		{
			name: "纯英文路径3",
			args: args{filename: "./ABC/D//EF/MAD_NUMBER123"},
			want: "MAD_NUMBER123",
		},
		{
			name: "混合路径1",
			args: args{filename: "不知道/是什么.exe"},
			want: "是什么",
		},
		{
			name: "混合路径2",
			args: args{filename: "不知道/是什么hello123.exe"},
			want: "是什么hello123",
		},
		{
			name: "双斜杠混合路径1",
			args: args{filename: "D:\\不知道\\是什么hello123.exe"},
			want: "是什么hello123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPureName(tt.args.filename); got != tt.want {
				t.Errorf("GetPureName() = %v, want %v", got, tt.want)
			}
		})
	}
}

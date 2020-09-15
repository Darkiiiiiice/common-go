package strings

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsEmpty_EmptyString",
			args: args{str: ""},
			want: true,
		},
		{
			name: "TestIsEmpty_EmptyStringAtPrefix",
			args: args{str: "   apple"},
			want: false,
		},
		{
			name: "TestIsEmpty_NotEmptyString",
			args: args{str: "apple"},
			want: false,
		},
		{
			name: "TestIsEmpty_EmptyStringAtSuffix",
			args: args{str: "apple   "},
			want: false,
		},
		{
			name: "TestIsEmpty_Blank",
			args: args{str: "    "},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.str); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.args.str); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbbreviate2(t *testing.T) {
	type args struct {
		str      string
		offset   int
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abbreviate2(tt.args.str, tt.args.offset, tt.args.maxWidth); got != tt.want {
				t.Errorf("Abbreviate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbbreviate(t *testing.T) {
	type args struct {
		str          string
		abbrevMarker string
		offset       int
		maxWidth     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "",
				abbrevMarker: "...",
				offset:       0,
				maxWidth:     100,
			},
			want: "",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "---",
				offset:       -1,
				maxWidth:     10,
			},
			want: "abcdefg---",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: ",",
				offset:       0,
				maxWidth:     10,
			},
			want: "abcdefghi,",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: ",",
				offset:       1,
				maxWidth:     10,
			},
			want: "abcdefghi,",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: ",",
				offset:       2,
				maxWidth:     10,
			},
			want: "abcdefghi,",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "::",
				offset:       4,
				maxWidth:     10,
			},
			want: "::efghij::",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "...",
				offset:       6,
				maxWidth:     10,
			},
			want: "...ghij...",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "*",
				offset:       9,
				maxWidth:     10,
			},
			want: "*ghijklmno",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "'",
				offset:       10,
				maxWidth:     10,
			},
			want: "'ghijklmno",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghijklmno",
				abbrevMarker: "!",
				offset:       12,
				maxWidth:     10,
			},
			want: "!ghijklmno",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghij",
				abbrevMarker: "abra",
				offset:       0,
				maxWidth:     4,
			},
			want: "abcdefghij",
		},
		{
			name: "TestAbbreviate_Empty_String",
			args: args{
				str:          "abcdefghij",
				abbrevMarker: "...",
				offset:       5,
				maxWidth:     6,
			},
			want: "abcdefghij",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abbreviate(tt.args.str, tt.args.abbrevMarker, tt.args.offset, tt.args.maxWidth); got != tt.want {
				t.Errorf("Abbreviate4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsEmpty_EmptyString",
			args: args{str: ""},
			want: false,
		},
		{
			name: "TestIsEmpty_EmptyStringAtPrefix",
			args: args{str: "   apple"},
			want: true,
		},
		{
			name: "TestIsEmpty_NotEmptyString",
			args: args{str: "apple"},
			want: true,
		},
		{
			name: "TestIsEmpty_EmptyStringAtSuffix",
			args: args{str: "apple   "},
			want: true,
		},
		{
			name: "TestIsEmpty_Blank",
			args: args{str: "    "},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.args.str); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

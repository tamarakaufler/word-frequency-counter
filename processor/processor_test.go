package processor

import (
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"
)

func Test_preprocess(t *testing.T) {
	type args struct {
		sep  *regexp.Regexp
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// test 1
		{
			name: "Processing of a regular file with no extra empty lines",
			args: args{
				sep:  ContentSeparatorRegex,
				file: "test/file1.txt",
			},
			want: []string{
				"One two three four five",
				"One three five",
				"Six seven eight nine ten",
				"Six eight ten",
				"Six eight ten",
				"One three five six eight ten",
				"Six eight ten",
				"Five six eight ten",
				"Six eight ten",
			},
		},
		{
			name: "Processing of a regular file with extra empty lines",
			args: args{
				sep:  ContentSeparatorRegex,
				file: "test/file2.txt",
			},
			want: []string{
				"One two three four five",
				"One three five",
				"Six seven eight nine ten",
				"Six eight ten",
				"Six eight ten",
				"One three five six eight ten",
				"Six eight ten",
				"Five six eight ten",
				"Six eight ten",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := ioutil.ReadFile(tt.args.file)
			got := preprocess(tt.args.sep, string(b))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preprocess() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		sep     *regexp.Regexp
		segment string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// test 1
		{
			name: "Processing sentence containing only lowercase letters",
			args: args{
				sep:     WordRegex,
				segment: "I am here just now, just now",
			},
			want: map[string]int{
				"am":   1,
				"here": 1,
				"just": 2,
				"now":  2,
			},
		},
		// test 2
		{
			name: "Processing sentence containing lower and uppercase letters",
			args: args{
				sep:     WordRegex,
				segment: "I am here just now, Just Now",
			},
			want: map[string]int{
				"am":   1,
				"here": 1,
				"just": 2,
				"now":  2,
			},
		},
		// test 3
		{
			name: "Processing sentence containing letters and numbers",
			args: args{
				sep:     WordRegex,
				segment: "I am 123 here just5 now, Just5 Now",
			},
			want: map[string]int{
				"am": 1,
				//"123":   1,
				"here":  1,
				"just5": 2,
				"now":   2,
			},
		},
		// test 4
		{
			name: "Processing sentence containing letters, numbers and non word characters",
			args: args{
				sep:     WordRegex,
				segment: "I am% 123 - here just5 :now, Just5 'Now",
			},
			want: map[string]int{
				"am": 1,
				//"123":   1,
				"here":  1,
				"just5": 2,
				"now":   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.sep, tt.args.segment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

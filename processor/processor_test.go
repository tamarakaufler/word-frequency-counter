package processor

import (
	"io/ioutil"
	"reflect"
	"regexp"
	"sync"
	"testing"
)

func TestProcessor_Run(t *testing.T) {
	type fields struct {
		File    string
		Workers int
		jobs    chan job
		results chan result
		done    chan struct{}
		index   map[string]int
		wg      sync.WaitGroup
		mu      sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				File:    tt.fields.File,
				Workers: tt.fields.Workers,
				jobs:    tt.fields.jobs,
				results: tt.fields.results,
				done:    tt.fields.done,
				index:   tt.fields.index,
				wg:      tt.fields.wg,
				mu:      tt.fields.mu,
			}
			if err := p.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Processor.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProcessor_processorInit(t *testing.T) {
	type fields struct {
		File    string
		Workers int
		jobs    chan job
		results chan result
		done    chan struct{}
		index   map[string]int
		wg      sync.WaitGroup
		mu      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				File:    tt.fields.File,
				Workers: tt.fields.Workers,
				jobs:    tt.fields.jobs,
				results: tt.fields.results,
				done:    tt.fields.done,
				index:   tt.fields.index,
				wg:      tt.fields.wg,
				mu:      tt.fields.mu,
			}
			p.processorInit()
		})
	}
}

func TestProcessor_sendjob(t *testing.T) {
	type fields struct {
		File    string
		Workers int
		jobs    chan job
		results chan result
		done    chan struct{}
		index   map[string]int
		wg      sync.WaitGroup
		mu      sync.Mutex
	}
	type args struct {
		segments []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				File:    tt.fields.File,
				Workers: tt.fields.Workers,
				jobs:    tt.fields.jobs,
				results: tt.fields.results,
				done:    tt.fields.done,
				index:   tt.fields.index,
				wg:      tt.fields.wg,
				mu:      tt.fields.mu,
			}
			p.sendjob(tt.args.segments)
		})
	}
}

func TestProcessor_worker(t *testing.T) {
	type fields struct {
		File    string
		Workers int
		jobs    chan job
		results chan result
		done    chan struct{}
		index   map[string]int
		wg      sync.WaitGroup
		mu      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				File:    tt.fields.File,
				Workers: tt.fields.Workers,
				jobs:    tt.fields.jobs,
				results: tt.fields.results,
				done:    tt.fields.done,
				index:   tt.fields.index,
				wg:      tt.fields.wg,
				mu:      tt.fields.mu,
			}
			p.worker()
		})
	}
}

func TestProcessor_reducer(t *testing.T) {
	type fields struct {
		File    string
		Workers int
		jobs    chan job
		results chan result
		done    chan struct{}
		index   map[string]int
		wg      sync.WaitGroup
		mu      sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				File:    tt.fields.File,
				Workers: tt.fields.Workers,
				jobs:    tt.fields.jobs,
				results: tt.fields.results,
				done:    tt.fields.done,
				index:   tt.fields.index,
				wg:      tt.fields.wg,
				mu:      tt.fields.mu,
			}
			p.reducer()
		})
	}
}

func Test_checkInput(t *testing.T) {
	type args struct {
		p *Processor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// test 1
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkInput(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("checkInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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

func Test_postprocess(t *testing.T) {
	type args struct {
		index map[string]int
		part  map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postprocess(tt.args.index, tt.args.part)
		})
	}
}

func Test_display(t *testing.T) {
	type args struct {
		index map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			display(tt.args.index)
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		sep     *regexp.Regexp
		clean   *regexp.Regexp
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
				sep:     SegmentSeparatorRegex,
				clean:   SegmentCleanupRegex,
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
				sep:     SegmentSeparatorRegex,
				clean:   SegmentCleanupRegex,
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
				sep:     SegmentSeparatorRegex,
				clean:   SegmentCleanupRegex,
				segment: "I am 123 here just5 now, Just5 Now",
			},
			want: map[string]int{
				"am":    1,
				"123":   1,
				"here":  1,
				"just5": 2,
				"now":   2,
			},
		},
		// test 4
		{
			name: "Processing sentence containing letters, numbers and non word characters",
			args: args{
				sep:     SegmentSeparatorRegex,
				clean:   SegmentCleanupRegex,
				segment: "I am% 123 - here just5 :now, Just5 'Now",
			},
			want: map[string]int{
				"am":    1,
				"123":   1,
				"here":  1,
				"just5": 2,
				"now":   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.sep, tt.args.clean, tt.args.segment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

package processor

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	ErrIncorrectInput     = "Incorrect input: %s"
	ContentSeparatorRegex = regexp.MustCompile(`\s*\.\s*`)
	SegmentSeparatorRegex = regexp.MustCompile(`\s+`)
	SegmentCleanupRegex   = regexp.MustCompile(`.*?([a-zA-Z0-9]{2,}).*`)
)

type job struct {
	id      int
	segment string
	start   time.Time
}

type result struct {
	job
	part map[string]int
}

type Processor struct {
	File    string
	Workers int

	jobs    chan job
	results chan result
	done    chan struct{}
	index   map[string]int
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func (p *Processor) Run() error {
	if err := checkInput(p); err != nil {
		return err
	}
	// initialize the Processor instance
	p.processorInit()

	// turn the file content into segments
	// that can be concurretly processed
	b, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}
	segments := preprocess(ContentSeparatorRegex, string(b))
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.sendjob(segments)
	}()

	go p.reducer()

	// set up the worker pool
	i := 0
	p.wg.Add(p.Workers)
	for i < p.Workers {
		go func(i int) {
			defer p.wg.Done()
			wi := i + 1
			p.worker(wi)
		}(i)
		i++
	}

	p.wg.Wait()
	p.done <- struct{}{}

	display(p.index)

	return nil
}

func (p *Processor) processorInit() {
	p.jobs = make(chan job, 20)
	p.results = make(chan result)
	p.done = make(chan struct{})

	p.index = make(map[string]int)
}

func (p *Processor) sendjob(segments []string) {
	defer close(p.jobs)

	for i, s := range segments {
		j := job{
			id:      i,
			segment: s,
			start:   time.Now(),
		}

		p.jobs <- j
	}
}

func (p *Processor) worker(i int) {
	for j := range p.jobs {
		sr := process(SegmentSeparatorRegex, SegmentCleanupRegex, j.segment)
		if len(sr) == 0 {
			continue
		}

		elapsed := time.Since(j.start)

		log.Printf("Worker %d - job [%s] took %s\n", i, j.segment, elapsed)

		r := result{
			job:  j,
			part: sr,
		}
		p.results <- r
	}
}

func (p *Processor) reducer() {
	for {
		select {
		case r := <-p.results:
			p.mu.Lock()
			postprocess(p.index, r.part)
			p.mu.Unlock()
		case <-p.done:
			fmt.Println("Finished")
			break
		}
	}
}

func checkInput(p *Processor) error {
	if p.File == "" {
		return fmt.Errorf(ErrIncorrectInput, "file path cannot be empty")
	}
	if _, err := os.Stat(p.File); err != nil {
		return fmt.Errorf(ErrIncorrectInput, err)
	}
	if p.Workers <= 0 {
		p.Workers = 2
	}
	return nil
}

// preprocess takes the file content and turns it into
// a list of segments that can be processed concurrently
func preprocess(re *regexp.Regexp, content string) []string {
	content = strings.Replace(content, "(\n|:)", " ", -1)
	list := re.Split(content, -1)

	last := len(list) - 1
	if list[last] == "" {
		list = list[:last]
	}

	return list
}

// process processes a segment
func process(sep *regexp.Regexp, clean *regexp.Regexp, segment string) map[string]int {
	processed := make(map[string]int)
	parts := sep.Split(segment, -1)
	if len(parts) == 0 {
		return nil
	}
	for _, p := range parts {
		if len(p) < 2 {
			continue
		}
		p = strings.ToLower(p)

		p = clean.ReplaceAllString(p, "$1")
		if p != "" {
			processed[p]++
		}
	}

	return processed
}

// postprocess incorporates a part result into the final index
func postprocess(index, part map[string]int) {
	for pk, pv := range part {
		index[pk] = index[pk] + pv
	}
}

func display(index map[string]int) {
	var keys []string
	for k := range index {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, index[k])
	}
}

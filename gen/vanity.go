package gen

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	// Number of seconds to wait between each update to the keys per second
	// monitor.
	kpsUpdateSecs = 10
	// Number of seconds each worker waits in between notifying the speedMonitor
	// with how many keys have been generated while they are actively searching
	// for one that matches the pattern.
	keyUpdateSecs = 15
)

var (
	defaultNumWorkers = uint32(runtime.NumCPU())
)

// SoVane (couldn't resist a random music reference somewhere)
type SoVane struct {
	Addr string
	Wif  string

	ctype    string
	vanity   string
	compress bool

	sync.Mutex

	numWorkers uint32
	started    bool
	Finished   bool

	Wg               sync.WaitGroup
	workerWg         sync.WaitGroup
	updateNumWorkers chan struct{}
	queryKeysPerSec  chan float64
	updateKeys       chan uint64
	speedMonitorQuit chan struct{}
	Quit             chan struct{}
}

func (s *SoVane) speedMonitor() {
	var keysPerSec float64
	var totalKeys uint64

	ticker := time.NewTicker(time.Second * kpsUpdateSecs)
	defer ticker.Stop()

out:
	for {
		select {
		case numKeys := <-s.updateKeys:
			totalKeys += numKeys

		case <-ticker.C:
			curKeysPerSec := float64(totalKeys) / kpsUpdateSecs
			if keysPerSec == 0 {
				keysPerSec = curKeysPerSec
			}
			keysPerSec = (keysPerSec + curKeysPerSec) / 2
			totalKeys = 0
			//if keysPerSec != 0 {
			fmt.Printf("Generator speed: %6.0f keys/s\n", keysPerSec)
			//}

		case s.queryKeysPerSec <- keysPerSec:
			// nothing

		case <-s.speedMonitorQuit:
			break out
		}
	}

	s.Wg.Done()
}

func (s *SoVane) generateAddress(quit chan struct{}) {
	ticker := time.NewTicker(time.Second * keyUpdateSecs)
	defer ticker.Stop()

out:
	for {
		select {
		case <-quit:
			break out
		default:
			//
		}

		wif, addr, err := GenerateAddress(s.ctype, s.compress)
		if err != nil {
			fmt.Println(err)
			break
		}

		if strings.HasPrefix(addr, s.vanity) {
			s.Wif = wif
			s.Addr = addr
			break
		}

		s.updateKeys <- 1
	}

	s.workerWg.Done()
	s.Stop()
}

func (s *SoVane) generateWorkerController() {
	var runningWorkers []chan struct{}
	launchWorkers := func(numWorkers uint32) {
		for i := uint32(0); i < numWorkers; i++ {
			quit := make(chan struct{})
			runningWorkers = append(runningWorkers, quit)

			s.workerWg.Add(1)
			go s.generateAddress(quit)
		}
	}

	runningWorkers = make([]chan struct{}, 0, s.numWorkers)
	launchWorkers(s.numWorkers)

out:
	for {
		select {
		// Update the number of running workers
		case <-s.updateNumWorkers:
			numRunning := uint32(len(runningWorkers))
			if s.numWorkers == numRunning {
				continue
			}

			// Add new workers
			if s.numWorkers > numRunning {
				launchWorkers(s.numWorkers - numRunning)
				continue
			}

			// Signal the most recently created goroutines to exit
			for i := numRunning - 1; i >= s.numWorkers; i-- {
				close(runningWorkers[i])
				runningWorkers[i] = nil
				runningWorkers = runningWorkers[:i]
			}
		case <-s.Quit:
			for _, quit := range runningWorkers {
				close(quit)
			}
			break out
		}
	}

	s.workerWg.Wait()
	close(s.speedMonitorQuit)
	s.Wg.Done()
}

func (s *SoVane) Start() {
	s.Lock()
	defer s.Unlock()

	if s.started {
		return
	}

	s.Quit = make(chan struct{})
	s.speedMonitorQuit = make(chan struct{})
	s.Wg.Add(2)
	go s.speedMonitor()
	go s.generateWorkerController()

	s.started = true
}

func (s *SoVane) Stop() {
	s.Lock()
	defer s.Unlock()

	if !s.started {
		return
	}

	close(s.Quit)

	s.started = false
	s.Finished = true
}

func NewVanityGen(coin, vane string, press bool) *SoVane {
	return &SoVane{
		ctype:            coin,
		vanity:           vane,
		compress:         press,
		numWorkers:       defaultNumWorkers,
		updateNumWorkers: make(chan struct{}),
		queryKeysPerSec:  make(chan float64),
		updateKeys:       make(chan uint64),
	}
}

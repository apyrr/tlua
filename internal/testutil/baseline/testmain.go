package baseline

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/apyrr/tlua/internal/collections"
)

var (
	// recordedBaselines tracks all baseline file paths that were written during the test run.
	recordedBaselines collections.SyncSet[string]

	// trackingInitialized is set to true when Track() is called.
	trackingInitialized bool
)

// trackingDirEnv names the directory where tracking files should be written; baseline
// tracking is enabled when it is set. Herebyfile.mjs sets it, to a fresh temporary
// directory, for full test runs.
//
// It must be read while tests run, never at package init or in TestMain before m.Run:
// the go test cache keys a package's result on the environment variables its tests
// read, and only reads made during the run are recorded. A read outside the run would
// let a package that writes baselines replay a cached result that records nothing,
// and every baseline it owns would then be reported unused.
const trackingDirEnv = "TLUA_BASELINE_TRACKING_DIR"

// Track sets up baseline tracking and returns a cleanup function that writes the tracking file.
// It should be called from TestMain using defer:
//
//	func TestMain(m *testing.M) {
//	    defer baseline.Track()()
//	    m.Run()
//	}
func Track() func() {
	trackingInitialized = true

	// Hash the entire call stack to create a unique filename per calling package.
	// This must be done in Track(), not in the deferred cleanup, because
	// the deferred function's call stack won't include the caller's info.
	var pcs [32]uintptr
	n := runtime.Callers(2, pcs[:]) // Skip Track and runtime.Callers
	h := fnv.New64a()
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		h.Write([]byte(frame.File))
		if !more {
			break
		}
	}
	trackingFile := fmt.Sprintf("%016x.txt", h.Sum64())

	return func() {
		// After tests complete, write the recorded baselines. Nothing is recorded unless
		// recordBaseline saw the tracking directory set during the run.
		if trackingDir := os.Getenv(trackingDirEnv); trackingDir != "" {
			writeRecordedBaselines(filepath.Join(trackingDir, trackingFile))
		}
	}
}

// recordBaseline adds a baseline file path to the recorded set.
// The path should be relative to the baselines/reference directory.
func recordBaseline(t testing.TB, relativePath string) {
	if os.Getenv(trackingDirEnv) != "" {
		if !trackingInitialized {
			t.Error("baseline: package uses baselines but TestMain did not call baseline.Track(). " +
				"Please add a TestMain function with: defer baseline.Track()()")
			return
		}
		recordedBaselines.Add(relativePath)
	}
}

// writeRecordedBaselines writes the list of recorded baseline files to a tracking file.
func writeRecordedBaselines(trackingPath string) {
	if recordedBaselines.Size() == 0 {
		return
	}

	if err := doWriteRecordedBaselines(trackingPath); err != nil {
		fmt.Fprintf(os.Stderr, "baseline: failed to write tracking file %s: %v\n", trackingPath, err)
		os.Exit(1)
	}
}

func doWriteRecordedBaselines(trackingPath string) error {
	f, err := os.Create(trackingPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for baseline := range recordedBaselines.Keys() {
		if _, err := fmt.Fprintln(w, baseline); err != nil {
			return err
		}
	}
	return w.Flush()
}

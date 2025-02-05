package schedule

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/EdgeLordKirito/wallpapersetter/internal/config"
	"github.com/EdgeLordKirito/wallpapersetter/internal/filevalidator"
	"github.com/EdgeLordKirito/wallpapersetter/platforms/independent"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	dur, err := time.ParseDuration(interval)
	if err != nil {
		return fmt.Errorf("Could not parse time string reason '%v'", err)
	}
	conf, err := config.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Unable to read Config %v", err)
	}
	strategy := independent.GetBackendStrategy(conf)
	sysConf := conf.GetOSConfig()
	var dirs []string = sysConf.Dirs()
	_, _ = strategy, dirs
	if path != "" {
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return fmt.Errorf("Need to provide an Directory to schedule")
		}
		dirs = []string{path}
	}

	ctx, cancel, wg, errChan, signalChan := setupTicker()
	wg.Add(1)
	defer wg.Wait()
	go startTicker(ctx, wg, errChan, dur, dirs, strategy)

	select {
	case sig := <-signalChan:
		fmt.Println("Received signal:", sig)
		cancel()
	case err := <-errChan:
		fmt.Println("Ticker stopped due to error:", err)
		cancel()
		return err
	}
	wg.Wait()
	fmt.Println("Program exiting...")
	return nil
}

func RunIter(dirs []string, strategy independent.WallpaperSetter) error {
	files, _, err := filevalidator.CollectImageFiles(dirs)
	if err != nil {
		return fmt.Errorf("Could not collect images from specified directories reason '%v'", err)
	}
	file, err := filevalidator.PickRandomFile(files)
	if err != nil {
		return fmt.Errorf("Could not pick random file from list reason '%v'", err)
	}
	return strategy.Set(file)
}

func startTicker(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error, pause time.Duration,
	dirs []string, strat independent.WallpaperSetter) {
	defer wg.Done()
	ticker := time.NewTicker(pause)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := RunIter(dirs, strat); err != nil {
				errChan <- err
				return
			}
		case <-ctx.Done():
			fmt.Println("Ticker stopping...")
			return
		}
		if ctx.Err() != nil {
			break
		}
	}
}

func setupTicker() (context.Context, context.CancelFunc, *sync.WaitGroup, chan error, chan os.Signal) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	signalChan := make(chan os.Signal, 1)

	// Handle OS signals
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	return ctx, cancel, &wg, errChan, signalChan
}

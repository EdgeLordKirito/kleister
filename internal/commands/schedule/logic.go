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
	"github.com/EdgeLordKirito/wallpapersetter/package/statusserver"
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
	dirs, err := getDirs(sysConf, path)
	if err != nil {
		return fmt.Errorf("Could not get list of directories reason '%v'", err)
	}
	tickSync := setupTicker()
	cBundle, err := newContextBundle(tickSync)
	settings := statusserver.ServerSettings{
		Adress: statusserver.DefaultAdress,
		Auth:   statusserver.TruthyAuth{}}
	tickSync.Waiter.Add(1)
	go statusserver.SetupStatusServer(*cBundle, settings)
	tickSync.Waiter.Add(1)
	go startTicker(tickSync, dur, dirs, strategy)
	defer tickSync.Waiter.Wait()

	select {
	case sig := <-tickSync.SignalChannel:
		fmt.Println("Received signal:", sig)
		tickSync.Cancel()
	case err := <-tickSync.ErrChannel:
		fmt.Println("Ticker stopped due to error:", err)
		tickSync.Cancel()
		return err
	case <-tickSync.Context.Done(): // Listen for context cancellation
		fmt.Println("Context canceled. Shutting down...")
	}
	tickSync.Waiter.Wait()
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

func startTicker(tickSync *tickerSync, pause time.Duration,
	dirs []string, strat independent.WallpaperSetter) {
	defer tickSync.Waiter.Done()
	ticker := time.NewTicker(pause)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := RunIter(dirs, strat); err != nil {
				tickSync.ErrChannel <- err
				return
			}
		case <-tickSync.Context.Done():
			fmt.Println("Ticker stopping...")
			return
		}
		if tickSync.Context.Err() != nil {
			break
		}
	}
}

func setupTicker() *tickerSync {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	signalChan := make(chan os.Signal, 1)

	// Handle OS signals
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := tickerSync{
		Context:       ctx,
		Cancel:        cancel,
		Waiter:        &wg,
		ErrChannel:    errChan,
		SignalChannel: signalChan,
	}

	return &ticker
}

func getDirs(conf config.Config, p string) ([]string, error) {
	dirs := conf.Dirs()
	if p != "" {
		info, err := os.Stat(p)
		if err != nil {
			return []string{}, err
		}
		if !info.IsDir() {
			return []string{}, fmt.Errorf("Need to provide an Directory to schedule")
		}
		dirs = []string{p}
	}
	if len(dirs) == 0 {
		return []string{}, fmt.Errorf("There has to be atleast one Directory" +
			"specified in the config for the current OS")
	}
	return dirs, nil
}

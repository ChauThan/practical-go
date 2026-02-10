package player

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"

	"youtube-tui/internal/config"
	"youtube-tui/internal/logging"
)

// Player implements the interfaces.MediaPlayer interface for playing audio/video content.
// It uses MPV (media player) for playback.
type Player struct {
	executable string
	autoStop   bool
	currentCmd *exec.Cmd
	mu         sync.Mutex
}

// NewPlayer creates a new Player instance with the given configuration.
// If cfg is nil, default configuration will be used.
func NewPlayer(cfg *config.Provider) *Player {
	if cfg == nil {
		cfg = config.NewProvider()
	}

	return &Player{
		executable: cfg.GetPlayerExecutable(),
		autoStop:   cfg.GetPlayerAutoStop(),
	}
}

// Play starts playback for the given video ID.
// It constructs a YouTube URL for the video and launches MPV to play it.
// If another video is currently playing, it will be stopped first.
func (p *Player) Play(videoID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	logging.Info("starting playback", "videoID", videoID)

	// Stop any currently playing instance
	if p.currentCmd != nil {
		logging.Debug("stopping current playback before starting new one")
		p.stopInternal()
	}

	// Construct the direct YouTube URL
	// TODO: In the future, get the actual video URL from an API
	// for now, use the platform-specific method to play YouTube videos
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin", "linux":
		// Use MPV for macOS and Linux
		cmd = exec.Command(p.executable, url)
	case "windows":
		// Use MPV for Windows
		cmd = exec.Command(p.executable, url)
	default:
		logging.Error("unsupported OS for playback", "os", runtime.GOOS)
		return fmt.Errorf("player: unsupported operating system '%s' for MPV playback", runtime.GOOS)
	}

	logging.Debug("executing player command", "executable", p.executable, "url", url)

	// Start the player
	if err := cmd.Start(); err != nil {
		logging.Error("failed to start playback", "videoID", videoID, "error", err)
		return fmt.Errorf("player: failed to start playback for video ID '%s': %w", videoID, err)
	}

	p.currentCmd = cmd
	logging.Info("playback started successfully", "videoID", videoID)

	// If auto-stop is enabled, wait for the command to finish
	// and clean up
	if p.autoStop {
		logging.Debug("auto-stop enabled, waiting for playback to complete")
		go func() {
			cmd.Wait()
			p.mu.Lock()
			defer p.mu.Unlock()
			p.currentCmd = nil
			logging.Info("playback completed", "videoID", videoID)
		}()
	}

	return nil
}

// Pause pauses the current playback.
// Note: MPV might not support pause on all platforms.
func (p *Player) Pause() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.currentCmd == nil {
		return fmt.Errorf("player: no video is currently playing (cannot pause)")
	}

	// TODO: Implement pause functionality
	// This requires MPV's IPC (Inter-Process Communication) to send pause commands
	// For now, pause is not supported via direct command execution
	return fmt.Errorf("player: pause functionality not yet implemented")
}

// Stop stops the current playback and cleans up resources.
func (p *Player) Stop() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.stopInternal()
}

// stopInternal stops the current playback without locking.
// Internal method that assumes the caller holds the mutex.
func (p *Player) stopInternal() error {
	if p.currentCmd == nil {
		return fmt.Errorf("player: no video is currently playing (cannot stop)")
	}

	if err := p.currentCmd.Process.Kill(); err != nil {
		return fmt.Errorf("player: failed to stop playback: %w", err)
	}

	p.currentCmd = nil
	return nil
}

// IsPlaying returns true if a video is currently playing.
func (p *Player) IsPlaying() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.currentCmd != nil
}

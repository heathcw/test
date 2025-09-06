package main

import (
	"fmt"
	"os"
	"io"
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func loadWav(path string) (*audio.Player, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer f.Close()

    // Read entire file into memory
    data, err := io.ReadAll(f)
    if err != nil {
        return nil, fmt.Errorf("error reading file: %v", err)
    }

    // Decode from in-memory buffer
    d, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
    if err != nil {
        return nil, fmt.Errorf("error decoding: %v", err)
    }

    // Create a new audio player
    p, err := audioContext.NewPlayer(d)
    if err != nil {
        return nil, fmt.Errorf("error creating player: %v", err)
    }

    return p, nil
}
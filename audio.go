package pacman

import (
	"io"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/skatiyar/pacman/assets"
)

type Audio struct {
	ctx     *audio.Context
	players *AudioPlayers
}

const (
	// This sample rate doesn't match with wav/vorbis's sample rate,
	// but decoders adjust them.
	sampleRate = 48000
)

func NewAudio() (*Audio, error) {
	audioCtx, err := audio.NewContext(sampleRate)
	if err != nil {
		return nil, err
	}

	sounds, err := assets.LoadSounds(audioCtx)
	if err != nil {
		return nil, err
	}

	beginning, err := newAudioPlayer(audioCtx, sounds.Beginning)
	if err != nil {
		return nil, err
	}

	chomp, err := newAudioPlayer(audioCtx, sounds.Chomp)
	if err != nil {
		return nil, err
	}

	death, err := newAudioPlayer(audioCtx, sounds.Death)
	if err != nil {
		return nil, err
	}

	eatFlask, err := newAudioPlayer(audioCtx, sounds.EatFlask)
	if err != nil {
		return nil, err
	}

	eatGhost, err := newAudioPlayer(audioCtx, sounds.EatGhost)
	if err != nil {
		return nil, err
	}

	extraPac, err := newAudioPlayer(audioCtx, sounds.ExtraPac)
	if err != nil {
		return nil, err
	}

	return &Audio{
		ctx: audioCtx,
		players: &AudioPlayers{
			Beginning: beginning,
			Chomp:     chomp,
			Death:     death,
			EatFlask:  eatFlask,
			EatGhost:  eatGhost,
			ExtraPac:  extraPac,
		},
	}, nil
}

type AudioPlayers struct {
	Beginning *audio.Player
	Chomp     *audio.Player
	Death     *audio.Player
	EatFlask  *audio.Player
	EatGhost  *audio.Player
	ExtraPac  *audio.Player
}

func newAudioPlayer(ctx *audio.Context, src io.ReadCloser) (*audio.Player, error) {
	player, err := audio.NewPlayer(ctx, src)
	if err != nil {
		return nil, err
	}

	player.SetVolume(0.3)

	return player, nil
}

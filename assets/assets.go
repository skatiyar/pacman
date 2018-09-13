// Package assets contains font, image & sound resources needed by the game
package assets

import (
	"bytes"
	"image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/skatiyar/pacman/assets/fonts"
	"github.com/skatiyar/pacman/assets/images"
	"github.com/skatiyar/pacman/assets/sounds"
	"github.com/skatiyar/pacman/spritetools"
)

type Characters struct {
	Pacman *ebiten.Image
	Ghost1 *ebiten.Image
	Ghost2 *ebiten.Image
	Ghost3 *ebiten.Image
	Ghost4 *ebiten.Image
}

type Powers struct {
	Life          *ebiten.Image
	Invincibility *ebiten.Image
}

type Walls struct {
	ActiveCorner   *ebiten.Image
	ActiveSide     *ebiten.Image
	InActiveCorner *ebiten.Image
	InActiveSide   *ebiten.Image
}

type Assets struct {
	ArcadeFont *truetype.Font
	Skin       *ebiten.Image
	Characters *Characters
	Powers     *Powers
	Walls      *Walls
}

// LoadAssets converts the character images(png, jpg, ...) to
// ebiten image format and loads fonts.
func LoadAssets() (*Assets, error) {
	skin, skinErr := loadSkin()
	if skinErr != nil {
		return nil, skinErr
	}

	font, fontErr := loadArcadeFont()
	if fontErr != nil {
		return nil, fontErr
	}

	characters, charactersErr := loadCharacters()
	if charactersErr != nil {
		return nil, charactersErr
	}

	powers, powersErr := loadPowers()
	if powersErr != nil {
		return nil, powersErr
	}

	walls, wallsErr := loadWalls()
	if wallsErr != nil {
		return nil, wallsErr
	}

	return &Assets{
		ArcadeFont: font,
		Skin:       skin,
		Characters: characters,
		Powers:     powers,
		Walls:      walls,
	}, nil
}

func loadSkin() (*ebiten.Image, error) {
	sImage, sImageErr := png.Decode(bytes.NewReader(images.SkinPng))
	if sImageErr != nil {
		return nil, sImageErr
	}

	skin, skinErr := ebiten.NewImageFromImage(sImage, ebiten.FilterDefault)
	if skinErr != nil {
		return nil, skinErr
	}

	return skin, nil
}

func loadArcadeFont() (*truetype.Font, error) {
	return truetype.Parse(fonts.ArcadeTTF)
}

func loadCharacters() (*Characters, error) {
	cImage, cImageErr := png.Decode(bytes.NewReader(images.CharactersPng))
	if cImageErr != nil {
		return nil, cImageErr
	}

	characters, charactersErr := ebiten.NewImageFromImage(cImage, ebiten.FilterDefault)
	if charactersErr != nil {
		return nil, charactersErr
	}

	pacman, pacmanErr := spritetools.GetSprite(61, 64, 0, 0, characters)
	if pacmanErr != nil {
		return nil, pacmanErr
	}

	ghost1, ghost1Err := spritetools.GetSprite(56, 64, 66, 0, characters)
	if ghost1Err != nil {
		return nil, ghost1Err
	}

	ghost2, ghost2Err := spritetools.GetSprite(56, 64, 125, 0, characters)
	if ghost2Err != nil {
		return nil, ghost2Err
	}

	ghost3, ghost3Err := spritetools.GetSprite(56, 64, 185, 0, characters)
	if ghost3Err != nil {
		return nil, ghost3Err
	}

	ghost4, ghost4Err := spritetools.GetSprite(56, 64, 244, 0, characters)
	if ghost4Err != nil {
		return nil, ghost4Err
	}

	return &Characters{
		Pacman: pacman,
		Ghost1: ghost1,
		Ghost2: ghost2,
		Ghost3: ghost3,
		Ghost4: ghost4,
	}, nil
}

func loadPowers() (*Powers, error) {
	pImage, pImageErr := png.Decode(bytes.NewReader(images.PowersPng))
	if pImageErr != nil {
		return nil, pImageErr
	}

	powers, powersErr := ebiten.NewImageFromImage(pImage, ebiten.FilterDefault)
	if powersErr != nil {
		return nil, powersErr
	}

	life, lifeErr := spritetools.GetSprite(64, 64, 0, 0, powers)
	if lifeErr != nil {
		return nil, lifeErr
	}

	invinc, invincErr := spritetools.GetSprite(64, 64, 67, 0, powers)
	if invincErr != nil {
		return nil, invincErr
	}

	return &Powers{
		Life:          life,
		Invincibility: invinc,
	}, nil
}

func loadWalls() (*Walls, error) {
	wImage, wImageErr := png.Decode(bytes.NewReader(images.WallsPng))
	if wImageErr != nil {
		return nil, wImageErr
	}

	walls, wallsErr := ebiten.NewImageFromImage(wImage, ebiten.FilterDefault)
	if wallsErr != nil {
		return nil, wallsErr
	}

	inactiveCorner, inactiveCornerErr := spritetools.GetSprite(12, 12, 0, 0, walls)
	if inactiveCornerErr != nil {
		return nil, inactiveCornerErr
	}

	inactiveSide, inactiveSideErr := spritetools.GetSprite(40, 12, 12, 0, walls)
	if inactiveSideErr != nil {
		return nil, inactiveSideErr
	}

	activeCorner, activeCornerErr := spritetools.GetSprite(12, 12, 52, 0, walls)
	if activeCornerErr != nil {
		return nil, activeCornerErr
	}

	activeSide, activeSideErr := spritetools.GetSprite(40, 12, 64, 0, walls)
	if activeSideErr != nil {
		return nil, activeSideErr
	}

	return &Walls{
		ActiveCorner:   activeCorner,
		ActiveSide:     activeSide,
		InActiveCorner: inactiveCorner,
		InActiveSide:   inactiveSide,
	}, nil
}

type Sounds struct {
	Beginning *wav.Stream
	Chomp     *wav.Stream
	Death     *wav.Stream
	EatFlask  *wav.Stream
	EatGhost  *wav.Stream
	ExtraPac  *wav.Stream
}

// LoadSounds returns a struct with wav files decoded
// for the provided audio context.
func LoadSounds(ctx *audio.Context) (*Sounds, error) {
	beginning, beginningErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.BeginningWav))
	if beginningErr != nil {
		return nil, beginningErr
	}

	chomp, chompErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.ChompWav))
	if chompErr != nil {
		return nil, chompErr
	}

	death, deathErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.DeathWav))
	if deathErr != nil {
		return nil, deathErr
	}

	eatFlask, eatFlaskErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.EatFlaskWav))
	if eatFlaskErr != nil {
		return nil, eatFlaskErr
	}

	eatGhost, eatGhostErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.EatGhostWav))
	if eatGhostErr != nil {
		return nil, eatGhostErr
	}

	extraPac, extraPacErr := wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.ExtraPacWav))
	if extraPacErr != nil {
		return nil, extraPacErr
	}

	return &Sounds{
		Beginning: beginning,
		Chomp:     chomp,
		Death:     death,
		EatFlask:  eatFlask,
		EatGhost:  eatGhost,
		ExtraPac:  extraPac,
	}, nil
}

// Images
//go:generate file2byteslice -package=images -input=./images/skin.png -output=./images/skin.go -var=SkinPng
//go:generate file2byteslice -package=images -input=./images/characters.png -output=./images/characters.go -var=CharactersPng
//go:generate file2byteslice -package=images -input=./images/powers.png -output=./images/powers.go -var=PowersPng
//go:generate file2byteslice -package=images -input=./images/walls.png -output=./images/walls.go -var=WallsPng

// Fonts
//go:generate file2byteslice -package=fonts -input=./fonts/arcade-n.ttf -output=./fonts/arcade-n.go -var=ArcadeTTF

// Sounds
//go:generate file2byteslice -package=sounds -input=./sounds/beginning.wav -output=./sounds/beginning.go -var=BeginningWav
//go:generate file2byteslice -package=sounds -input=./sounds/chomp.wav -output=./sounds/chomp.go -var=ChompWav
//go:generate file2byteslice -package=sounds -input=./sounds/death.wav -output=./sounds/death.go -var=DeathWav
//go:generate file2byteslice -package=sounds -input=./sounds/eatfruit.wav -output=./sounds/eatflask.go -var=EatFlaskWav
//go:generate file2byteslice -package=sounds -input=./sounds/eatghost.wav -output=./sounds/eatghost.go -var=EatGhostWav
//go:generate file2byteslice -package=sounds -input=./sounds/extrapac.wav -output=./sounds/extrapac.go -var=ExtraPacWav

package assets

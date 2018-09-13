// Images
//go:generate file2byteslice -package=images -input=./images/skin.png -output=./images/skin.go -var=SkinPng
//go:generate file2byteslice -package=images -input=./images/characters.png -output=./images/characters.go -var=CharactersPng
//go:generate file2byteslice -package=images -input=./images/powers.png -output=./images/powers.go -var=PowersPng
//go:generate file2byteslice -package=images -input=./images/walls.png -output=./images/walls.go -var=WallsPng

// Fonts
//go:generate file2byteslice -package=fonts -input=./fonts/arcade-n.ttf -output=./fonts/arcade-n.go -var=ArcadeTTF

// Sounds
//go:generate file2byteslice -package=sounds -input=./sounds/beginning.mp3 -output=./sounds/beginning.go -var=BeginningMp3
//go:generate file2byteslice -package=sounds -input=./sounds/chomp.mp3 -output=./sounds/chomp.go -var=ChompMp3
//go:generate file2byteslice -package=sounds -input=./sounds/death.mp3 -output=./sounds/death.go -var=DeathMp3
//go:generate file2byteslice -package=sounds -input=./sounds/eatflask.mp3 -output=./sounds/eatflask.go -var=EatFlaskMp3
//go:generate file2byteslice -package=sounds -input=./sounds/eatghost.mp3 -output=./sounds/eatghost.go -var=EatGhostMp3
//go:generate file2byteslice -package=sounds -input=./sounds/extrapac.mp3 -output=./sounds/extrapac.go -var=ExtraPacMp3

package assets

//go:generate file2byteslice -package=images -input=./images/skin.png -output=./images/skin.go -var=SkinPng
//go:generate file2byteslice -package=images -input=./images/characters.png -output=./images/characters.go -var=CharactersPng
//go:generate file2byteslice -package=images -input=./images/powers.png -output=./images/powers.go -var=PowersPng
//go:generate file2byteslice -package=images -input=./images/walls.png -output=./images/walls.go -var=WallsPng
//go:generate file2byteslice -package=fonts -input=./fonts/arcade-n.ttf -output=./fonts/arcade-n.go -var=ArcadeTTF

package assets

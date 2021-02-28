package png

import (
	"io/ioutil"

	"gowasmgame/resources/images"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetPNGFromFile slurps file into memory
func GetPNGFromFile(filepath string) []byte {
	// slurp a file into memory
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return dat
}

// Get gets the png in byteslice form
func Get() []byte {
	return images.Tiles
}

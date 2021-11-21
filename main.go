package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phR0ze/n/pkg/sys"
)

var unique = map[string]bool{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the target directory to mix the filenames of")
		return
	}
	target := os.Args[1]
	files := sys.Files(target)

	// Add existing names to unique map so we don't collide
	for _, x := range files {
		name := strings.TrimSuffix(filepath.Base(x), filepath.Ext(x))
		if _, exists := unique[x]; exists {
			fmt.Printf("Error: %s already exists!\n", filepath.Base(x))
			return
		}
		unique[name] = true
	}

	// Now rename the existing files to new random names
	for _, x := range files {
		ext := filepath.Ext(x)
		newname := fmt.Sprintf("%s%s", gen_unique_5_char_str(), ext)
		newpath := filepath.Join(filepath.Dir(x), newname)
		fmt.Printf("%v => %v\n", x, newpath)
		sys.Move(x, newpath)
	}
}

func gen_unique_5_char_str() (i string) {
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)
	for {
		i = fmt.Sprintf("%005d", gen.Intn(10000))
		if _, exists := unique[i]; !exists {
			unique[i] = true
			break
		}
	}
	return
}

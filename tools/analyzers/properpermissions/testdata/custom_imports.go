package testdata

import (
	"crypto/rand"
	"fmt"
	ioAlias "io/ioutil"
	"math/big"
	osAlias "os"
	"path/filepath"
)

func UseAliasedPackages() {
	randPath, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	p := filepath.Join(tempDir(), fmt.Sprintf("/%d", randPath))
	_ = osAlias.MkdirAll(p, osAlias.ModePerm) // want "os and ioutil dir and file writing functions are not permissions-safe, use shared/fileutil"
	someFile := filepath.Join(p, "some.txt")
	_ = ioAlias.WriteFile(someFile, []byte("hello"), osAlias.ModePerm) // want "os and ioutil dir and file writing functions are not permissions-safe, use shared/fileutil"
}

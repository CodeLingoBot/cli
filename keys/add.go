package keys

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/errgo.v1"
	"github.com/Scalingo/cli/config"
)

func Add(name string, path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return errgo.Mask(err)
	}
	if stat.Mode().IsDir() {
		return errgo.Newf("%s: is a directory", path)
	}
	if stat.Size() > 10*1024*1024 {
		return errgo.Newf("%s: is too large (%v bytes)", stat.Size())
	}

	keyContent, err := ioutil.ReadFile(path)
	if err != nil {
		return errgo.Mask(err)
	}

	c := config.ScalingoClient()
	_, err = c.KeysAdd(name, string(keyContent))
	if err != nil {
		return errgo.Mask(err)
	}

	fmt.Printf("Key '%s' has been added.\n", name)
	return nil
}

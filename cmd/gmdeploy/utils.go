package main

import "io/ioutil"

func Save(name, data string) error {
	err := ioutil.WriteFile(name, []byte(data), 0644)
	return err
}

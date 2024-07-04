package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/FrMnJ/encryptor/encryption"
	"github.com/FrMnJ/encryptor/files"
)

func main() {
	fromFile := flag.String("from", "", "Path to the file to be encrypted")
	toFile := flag.String("to", "", "File name of the new encryted file, created if not exists in the current directory")
	mode := flag.String("mode", "enc", "dec for decryption and enc for encryption(default enc)")
	key := flag.String("key", "", "Key uses to encrypt or decrypt files")
	flag.Parse()
	if key != nil && *key != "" {
		fmt.Printf("Key length: %d\nKey: %s \n", len(*key), *key)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if *fromFile == "" || *toFile == "" {
		panic(flag.ErrHelp.Error())
	}

	switch *mode {
	case "enc":
		if !files.IsFile(*fromFile) {
			panic("from file is not a valid file path")
		}
		data, err := os.ReadFile(*fromFile)
		if err != nil {
			panic(err)
		}
		key := encryption.Generate32Key()
		fmt.Println("Share this key with the person you want to share the file with:", string(key))
		encFile, err := os.Create(filepath.Join(cwd, *toFile))
		defer encFile.Close()
		if err != err {
			panic(err)
		}
		encData, err := encryption.Encrypt(key, data)
		if err != nil {
			panic(err)
		}
		numEncBytes, err := encFile.Write(encData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Number of bytes encrypted %d to %s \n", numEncBytes, *toFile)
	case "dec":
		if *key == "" {
			panic(flag.ErrHelp.Error())
		}
		if !files.IsFile(*fromFile) {
			panic("from file is not a valid file path")
		}
		encData, err := os.ReadFile(*fromFile)
		if err != nil {
			panic(err)
		}
		decFile, err := os.Create(filepath.Join(cwd, *toFile))
		if err != nil {
			panic(err)
		}
		defer decFile.Close()
		decData, err := encryption.Decrypt([]byte(*key), encData)
		if err != nil {
			panic(err)
		}
		numDecBytes, err := decFile.Write(decData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Number of bytes decrypted %d to %s \n", numDecBytes, *toFile)
	default:
		panic("Invalid mode")
	}

	os.Exit(0)
}

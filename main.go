package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"inc"
	"io/ioutil"
	"os"
	"parser"
	"path/filepath"
	"strings"
	"utils"
)

func main() {
	var index int = 1

	if dat, err := ioutil.ReadFile("./BUILDER"); err == nil {
		if err := parser.Parse_builder(string(dat)); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	} else {
		panic("Could't open BUILDER, exits")
	}

	filepath.Walk("../MasterThesis/code", utils.Walk())
	if err := utils.DFS(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if strings.Compare(os.Args[1], "clean") == 0 {
		if _, err := os.Create(".state"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		index = 2
	}

	if err := parser.Parse_state(); err != nil {
		fmt.Println(err)
	}

	for _, v := range os.Args[index:] {
		err := utils.Execute(v)
		if err != nil {
			fmt.Println(err)
		}
	}
	if f, err := os.Create(".state"); err == nil {
		var buffer bytes.Buffer
		var enc *gob.Encoder = gob.NewEncoder(&buffer)
		for _, v := range inc.State {
			if err := enc.Encode(v); err != nil {
				fmt.Println("Error encoding state, exits")
				os.Exit(1)
			}
			if _, err := f.Write(buffer.Bytes()); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			buffer.Reset()
		}
	} else {
		fmt.Println("Couldn't create .state, exits")
		fmt.Println(err)
	}

}

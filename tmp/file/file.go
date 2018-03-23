package file

import (
	"io/ioutil"
	"os"
)
type File struct{
	Fi *os.File
}

func (fi *File)Read()string{

	f,err:=os.Open("d:/test.txt")
	if err != nil {
		panic(err)
	}

	b,err:=ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fi.Fi=f
	return string(b)

}
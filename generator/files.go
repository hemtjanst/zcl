package codegen

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type File struct {
	name   string
	parent *Dir
	info   os.FileInfo
}

func (f *File) ReadCluster() (*Cluster, error) {
	b, err := ioutil.ReadFile(f.Path())
	if err != nil {
		return nil, err
	}
	ret := &Cluster{}
	err = yaml.Unmarshal(b, ret)
	return ret, err

}

func (f *File) SubPath() string {
	return path.Join(f.parent.SubPath(), f.name)
}

func (f *File) Path() string {
	return path.Join(f.parent.Path(), f.name)
}

type Dir struct {
	name   string
	parent *Dir
	dirs   []*Dir
	files  []*File
}

func (d *Dir) SubPath() string {
	if d.parent == nil {
		return ""
	}
	return path.Join(d.parent.SubPath(), d.name)
}

func (d *Dir) Path() string {
	n := d.name
	if d.parent != nil {
		n = path.Join(d.parent.Path(), n)
	}
	return n
}

func (d *Dir) traverse() {
	files, _ := ioutil.ReadDir(d.Path())
	for _, f := range files {
		if f.IsDir() {
			dir := &Dir{
				name:   f.Name(),
				parent: d,
			}
			dir.traverse()
			d.dirs = append(d.dirs, dir)
		} else if strings.HasSuffix(f.Name(), ".yaml") {
			d.files = append(d.files, &File{
				name:   f.Name(),
				parent: d,
				info:   f,
			})
		} else {
			log.Printf("Unknown file: %+v", f)
		}
	}
}

func (d *Dir) Files() (files []*File) {
	for _, dir := range d.dirs {
		files = append(files, dir.Files()...)
	}
	files = append(files, d.files...)
	return
}

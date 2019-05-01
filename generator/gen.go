package codegen

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type BufWriter []byte

func (b *BufWriter) Write(bytes []byte) (int, error) {
	*b = append(*b, bytes...)
	return len(bytes), nil
}

func writeFile(f string, b []byte) error {
	b2, err2 := format.Source(b)
	if err2 == nil {
		b = b2
	}

	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		return err
	}
	return err2
}

func GenerateCluster(base, defPath, tplPath, outPath string) error {

	d := &Dir{name: filepath.Join(defPath, "clusters")}
	d.traverse()

	loadTpl := func(tpl string, funcs template.FuncMap) (*template.Template, error) {
		b, err := ioutil.ReadFile(filepath.Join(tplPath, tpl+".go.tmpl"))
		if err != nil {
			return nil, fmt.Errorf("loading template %s: %v", tpl, err)
		}
		return template.New(tpl).
			Funcs(funcs).
			Parse(string(b))
	}

	for _, v := range d.Files() {
		p := v.SubPath()
		dp := filepath.Join(filepath.Join(outPath, filepath.Dir(p)))
		_ = os.MkdirAll(dp, 0755)
		clusterTpl, err := loadTpl("clusters",
			template.FuncMap{
				"base": path.Base,
				"path": func(pkg ...string) string { return path.Join(append([]string{base}, pkg...)...) },
				"fmt":  fmt.Sprintf,
				"package": func() string {
					return filepath.Base(dp)
				},
				"joinCond": func(attrList []Attr, cond []Condition) string {
					var vstr []string
				next:
					for _, c := range cond {
						mpre := ""
						mpost := ""
						if c.Mask != "" {
							mpre = "("
							mpost = fmt.Sprintf(" & %s)", c.Mask)
						}
						for _, a := range attrList {
							if a.ID.Uint() == c.Attr.Uint() {
								vstr = append(vstr, fmt.Sprintf(
									"%sv.%s%s == %s",
									mpre,
									a.Name.Fmt(),
									mpost,
									c.Value,
								))

								continue next
							}
						}
						log.Fatalf("Cannot find attribute related to condition %+v", c)
					}
					return strings.Join(vstr, " && ")
				},
			})

		if err != nil {
			return fmt.Errorf("parsing clusters template: %v", err)
		}

		cl, err := v.ReadCluster()
		if err != nil {
			return fmt.Errorf("while reading %s: %v", v.name, err)
		}

		wr := &BufWriter{}
		err = clusterTpl.Execute(wr, cl)
		if err != nil {
			return fmt.Errorf("while executing template for %s: %v", v.name, err)
		}
		outFile := filepath.Join(dp, fmt.Sprintf("%04x_%s.go", cl.ID.Uint(), cl.Name.FmtUs()))

		if err = writeFile(outFile, []byte(*wr)); err != nil {
			return fmt.Errorf("trying to write %s: %v", outFile, err)
		}

	}
	return nil

}

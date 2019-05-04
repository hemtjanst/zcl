package generator

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
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

	clusterIndex := map[string][]string{}
	mfClusterIndex := map[uint16]map[string][]string{}

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

		if cl.MfCode.Valid() {
			mf := uint16(cl.MfCode.Uint())
			if _, ok := mfClusterIndex[mf]; !ok {
				mfClusterIndex[mf] = map[string][]string{}
			}
			if _, ok := mfClusterIndex[mf][filepath.Base(dp)]; !ok {
				mfClusterIndex[mf][filepath.Base(dp)] = []string{}
			}
			mfClusterIndex[mf][filepath.Base(dp)] = append(mfClusterIndex[mf][filepath.Base(dp)], cl.MfCode.Stn(cl.Name))

		} else {
			if _, ok := clusterIndex[filepath.Base(dp)]; !ok {
				clusterIndex[filepath.Base(dp)] = []string{}
			}
			clusterIndex[filepath.Base(dp)] = append(clusterIndex[filepath.Base(dp)], cl.MfCode.Stn(cl.Name))
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

	wr := &BufWriter{}
	_, _ = wr.Write([]byte(`
package cluster

import (
`))

	_, _ = wr.Write([]byte(fmt.Sprintf("    \"%s\"\n", base)))

	var domains []string

	for n := range clusterIndex {
		domains = append(domains, n)
	}
	sort.Strings(domains)
	for _, n := range domains {
		_, _ = wr.Write([]byte(fmt.Sprintf("    \"%s/cluster/%s\"\n", base, n)))
	}

	_, _ = wr.Write([]byte(`
)

var Clusters = map[zcl.ClusterID]zcl.Cluster{
`))

	for _, n := range domains {
		clusters := clusterIndex[n]
		sort.Strings(clusters)

		for _, v := range clusters {
			_, _ = wr.Write([]byte(fmt.Sprintf("    %s.%sID: %s.%sCluster,\n", n, v, n, v)))
		}
	}

	_, _ = wr.Write([]byte("}\n"))

	_ = writeFile(filepath.Join(outPath, "all.go"), []byte(*wr))

	return nil

}

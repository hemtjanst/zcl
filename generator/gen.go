package generator

import (
	"fmt"
	"go/format"
	"gopkg.in/yaml.v2"
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

func genHelpers(clusterID, pkg, tplPath, base string, mfc MfCode) template.FuncMap {
	// clusterID()
	return template.FuncMap{
		"base": path.Base,
		"path": func(pkg ...string) string { return path.Join(append([]string{base}, pkg...)...) },
		"fmt":  fmt.Sprintf,
		"package": func() string {
			return filepath.Base(pkg)
		},
		"clusterID": func() string {
			return clusterID
		},
		"mfc": func() MfCode {
			return mfc
		},
		"renderType": func(attr Attr) string {
			attrTpl, err := loadTpl(tplPath, "attr", genHelpers(clusterID, pkg, tplPath, base, mfc))
			if err != nil {
				log.Fatalf("Cannot load attribute template: %+v", err)
			}
			wr := &BufWriter{}
			err = attrTpl.Execute(wr, attr)
			if err != nil {
				log.Fatalf("Cannot execute attribute template: %+v", err)
			}
			return string(*wr)
		},
		"renderCluster": func(cl Cluster) string {
			attrTpl, err := loadTpl(tplPath, "cluster", genHelpers(cl.ID.Hex4(), pkg, tplPath, base, cl.MfCode))
			if err != nil {
				log.Fatalf("Cannot load attribute template: %+v", err)
			}
			wr := &BufWriter{}
			err = attrTpl.Execute(wr, cl)
			if err != nil {
				log.Fatalf("Cannot execute attribute template: %+v", err)
			}
			return string(*wr)
		},
		"renderCommand": func(cmd Command, newClusterID ...string) string {
			clid := clusterID
			if len(newClusterID) > 0 {
				clid = newClusterID[0]
			}
			cmdTpl, err := loadTpl(tplPath, "command", genHelpers(clid, pkg, tplPath, base, mfc))
			if err != nil {
				log.Fatalf("Cannot load command template: %+v", err)
			}
			wr := &BufWriter{}
			err = cmdTpl.Execute(wr, cmd)
			if err != nil {
				log.Fatalf("Cannot execute command template: %+v", err)
			}
			return string(*wr)
		},
		"condDesc": func(name string, cond []Condition) string {
			var desc []string
			for _, c := range cond {
				d := c.Desc.Trim()
				if d != "" {
					desc = append(desc, d)
				}
			}
			if len(desc) == 0 {
				return ""
			}
			return Desc(strings.Join(desc, " and ")).Comment(name + " is only included if")
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
				if c.Name != "" {
					vstr = append(vstr, fmt.Sprintf(
						"%sv.%s%s == %s",
						mpre,
						c.Name.Fmt(),
						mpost,
						c.Value,
					))
					continue next
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
	}

}

func GenerateZdo(base, defPath, tplPath, outPath string) error {
	file := filepath.Join(defPath, "zdo.yaml")

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	ret := &Zdo{}
	err = yaml.Unmarshal(b, ret)
	if err != nil {
		return err
	}
	zdoTpl, err := loadTpl(
		tplPath,
		"zdo",
		genHelpers("", "cluster", tplPath, base, MfCode("")),
	)

	if err != nil {
		return fmt.Errorf("while parsing template for ZDO: %v", err)
	}
	wr := &BufWriter{}
	err = zdoTpl.Execute(wr, ret)
	if err != nil {
		return fmt.Errorf("while executing template for ZDO: %v", err)
	}
	outFile := filepath.Join(outPath, "zdo.go")

	if err = writeFile(outFile, []byte(*wr)); err != nil {
		return fmt.Errorf("trying to write %s: %v", outFile, err)
	}

	return nil
}

func loadTpl(tplPath, tpl string, funcs template.FuncMap) (*template.Template, error) {
	b, err := ioutil.ReadFile(filepath.Join(tplPath, tpl+".go.tmpl"))
	if err != nil {
		return nil, fmt.Errorf("loading template %s: %v", tpl, err)
	}
	return template.New(tpl).
		Funcs(funcs).
		Parse(string(b))
}

func GenerateStruct(base, defPath, tplPath, outPath string) error {

	d := &Dir{name: filepath.Join(defPath, "clusters")}
	d.traverse()

	clusterIndex := map[string][]string{}
	mfClusterIndex := map[uint16]map[string][]string{}

	for _, v := range d.Files() {
		p := v.SubPath()
		log.Printf("Parsing %s", v.Path())
		clRoot, err := v.ReadCluster()
		if err != nil {
			return fmt.Errorf("while reading %s: %v", v.name, err)
		}

		if len(clRoot.Clusters) == 0 {
			continue
		}

		dp := filepath.Join(filepath.Join(outPath, filepath.Base(p)))
		dp = strings.TrimSuffix(dp, ".yaml")
		_ = os.MkdirAll(dp, 0755)

		structTpl, err := loadTpl(tplPath, "struct", genHelpers("", dp, tplPath, base, ""))
		if err != nil {
			return fmt.Errorf("parsing struct template: %v", err)
		}

		for typeName, typeVal := range clRoot.TypeMap {
			if typeVal.Name == "" {
				log.Printf("[WARN:%s] Type %s is missing name", p, typeName)
			}
			if typeVal.Name.Fmt() != typeName {
				log.Printf("[WARN:%s] Type %s doesn't have the same key as formatted name: %s", p, typeName, typeVal.Name.Fmt())
			}
		}

		for _, cl := range clRoot.Clusters {
			clusterTpl, err := loadTpl(tplPath, "cluster", genHelpers(cl.Name.Fmt()+"ID", dp, tplPath, base, cl.MfCode))
			if err != nil {
				return fmt.Errorf("parsing cluster template: %v", err)
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
		err = structTpl.Execute(wr, clRoot)
		if err != nil {
			return fmt.Errorf("while executing template for %s: %v", v.name, err)
		}
		outFile := filepath.Join(dp, "struct.go")
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

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
	if strings.HasSuffix(f, ".go") {
		b2, err2 := format.Source(b)
		if err2 == nil {
			b = b2
		}
	}

	err := ioutil.WriteFile(f, b, 0644)
	return err
}

func genHelpers(cluster *Cluster, zdo *Zdo, server bool, clusterID, pkg, tplPath, base string, mfc MfCode) template.FuncMap {
	// clusterID()
	return template.FuncMap{
		"base": path.Base,
		"path": func(pkg ...string) string { return path.Join(append([]string{base}, pkg...)...) },
		"fmt":  fmt.Sprintf,
		"zdo": func() bool {
			return zdo != nil
		},
		"withPath": func(path string, iPath string, cmdPath string, val interface{}) map[string]interface{} {
			return map[string]interface{}{
				"Path":    path,
				"CmdPath": cmdPath,
				"IPath":   iPath,
				"Val":     val,
			}
		},
		"strEsc": func(str interface{}) string {
			var s string
			if v, ok := str.(string); ok {
				s = v
			} else if v, ok := str.(fmt.Stringer); ok {
				s = v.String()
			} else {
				s = fmt.Sprintf("%v", str)
			}

			s = strings.TrimSpace(strings.ReplaceAll(s, "\r", ""))
			rows := strings.Split(s, "\n")
			var out []string
			if strings.Contains(s, "`") {
				for _, ss := range rows {
					ss = strings.ReplaceAll(
						strings.ReplaceAll(ss, `\`, `\\`),
						`"`, `\"`)
					out = append(out, `"`+ss+`"`) // Use quotes if string contains `
				}
			} else {
				for _, ss := range rows {
					out = append(out, "`"+ss+"`")
				}
				return fmt.Sprintf("`%s`", s)
			}

			return strings.Join(out, " + \"\\n\" + \n")
		},
		"package": func() string {
			return filepath.Base(pkg)
		},
		"clusterID": func() string {
			return clusterID
		},
		"mfc": func() MfCode {
			return mfc
		},
		"direction": func(c Command) string {
			if cluster != nil {
				for _, cmd := range cluster.Server.Command {
					if cmd.ID == c.ID {
						return "ClientToServer"
					}
				}
				for _, cmd := range cluster.Client.Command {
					if cmd.ID == c.ID {
						return "ServerToClient"
					}
				}
			}
			return ""
		},
		"findResponse": func(c Command) *Command {
			if c.Response == nil {
				return nil
			}
			rsp := *c.Response
			if cluster != nil && server {
				for _, cmd := range cluster.Client.Command {
					if cmd.ID.AsDecimal() == rsp.AsDecimal() {
						return &cmd
					}
				}
			} else if cluster != nil {
				for _, cmd := range cluster.Server.Command {
					if cmd.ID.AsDecimal() == rsp.AsDecimal() {
						return &cmd
					}
				}
			} else if zdo != nil {
				for _, cmd := range zdo.Commands {
					if cmd.ID.AsDecimal() == rsp.AsDecimal() {
						return &cmd
					}
				}
			}
			return nil
		},
		"renderType": func(attr Attr) string {
			attrTpl, err := loadTpl(tplPath, "attr.go", genHelpers(cluster, zdo, server, clusterID, pkg, tplPath, base, mfc))
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
			attrTpl, err := loadTpl(tplPath, "cluster.go", genHelpers(cluster, zdo, server, cl.ID.Hex4(), pkg, tplPath, base, cl.MfCode))
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
		"renderCommand": func(cmd Command, server bool, newClusterID ...string) string {
			clid := clusterID
			if len(newClusterID) > 0 {
				clid = newClusterID[0]
			}
			cmdTpl, err := loadTpl(tplPath, "command.go", genHelpers(cluster, zdo, server, clid, pkg, tplPath, base, mfc))
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
					mpre = mpre + "("
					mpost = fmt.Sprintf(" & %s)%s", c.Mask, mpost)
				}
				eq := "=="
				if c.Invert {
					eq = "!="
				}
				if c.Operator != "" {
					eq = c.Operator
				}
				if c.Name != "" {
					vstr = append(vstr, fmt.Sprintf(
						"%sv.%s%s %s %s",
						mpre,
						c.Name.Fmt(),
						mpost,
						eq,
						c.Value,
					))
					continue next
				}
				for _, a := range attrList {
					if a.ID.Uint() == c.Attr.Uint() {
						vstr = append(vstr, fmt.Sprintf(
							"%sv.%s%s %s %s",
							mpre,
							a.Name.Fmt(),
							mpost,
							eq,
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
		"zdo.go",
		genHelpers(nil, ret, false, "", "cluster", tplPath, base, MfCode("")),
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

func GenerateTypescript(base, defPath, tplPath, outPath string) error {

	data := map[string]interface{}{}

	file := filepath.Join(defPath, "zdo.yaml")
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	zdp := &Zdo{}
	err = yaml.Unmarshal(b, zdp)
	if err != nil {
		return err
	}
	data["zdp"] = zdp

	d := &Dir{name: filepath.Join(defPath, "clusters")}
	d.traverse()

	var cluster []Cluster
	for _, v := range d.Files() {
		clRoot, err := v.ReadCluster()
		if err != nil {
			return fmt.Errorf("while reading %s: %v", v.name, err)
		}

		if len(clRoot.Clusters) == 0 {
			continue
		}
		cluster = append(cluster, *clRoot)
	}

	data["cluster"] = cluster

	tpl, err := loadTpl(
		tplPath,
		"zyaml.ts",
		genHelpers(nil, zdp, false, "", "ZigBee", tplPath, base, MfCode("")),
	)

	if err != nil {
		return fmt.Errorf("while parsing template for ZDO: %v", err)
	}
	wr := &BufWriter{}
	err = tpl.ExecuteTemplate(wr, "main", data)
	if err != nil {
		return fmt.Errorf("while executing template for ZDO: %v", err)
	}
	if err = writeFile(outPath, []byte(*wr)); err != nil {
		return fmt.Errorf("trying to write %s: %v", outPath, err)
	}

	return nil
}

func loadTpl(tplPath, tpl string, funcs template.FuncMap) (*template.Template, error) {
	b, err := ioutil.ReadFile(filepath.Join(tplPath, tpl+".tmpl"))
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

		structTpl, err := loadTpl(tplPath, "struct.go", genHelpers(nil, nil, false, "", dp, tplPath, base, ""))
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
			clusterTpl, err := loadTpl(tplPath, "cluster.go", genHelpers(&cl, nil, false, cl.Name.Fmt()+"ID", dp, tplPath, base, cl.MfCode))
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

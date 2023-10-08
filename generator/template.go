package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

type TemplateDef struct {
	Values          []any           `xml:"values" toml:"values,omitempty" yaml:"values,omitempty" json:"values,omitempty"`
	TypeMap         map[string]Attr `xml:"types" toml:"types,omitempty" yaml:"types,omitempty" json:"types,omitempty"`
	ClusterPerValue *Cluster        `xml:"cluster_per_value" toml:"cluster_per_value,omitempty" yaml:"cluster_per_value,omitempty" json:"cluster_per_value,omitempty"`
}

func (c *Cluster) ExpandTemplates() error {
	for _, tpl := range c.Templates {
		types, err := tpl.Types()
		if err != nil {
			return err
		}
		for k, v := range types {
			if c.TypeMap == nil {
				c.TypeMap = map[string]Attr{}
			} else {
				if _, ok := c.TypeMap[k]; ok {
					return fmt.Errorf("duplicate type: %s", k)
				}
			}
			c.TypeMap[k] = v
		}
		cl, err := tpl.Clusters()
		if err != nil {
			return err
		}
		c.Clusters = append(c.Clusters, cl...)
	}
	// Forget about templates once they're merged
	c.Templates = nil
	return nil
}

func mustStrTemplate[T ~string](str T, v any) T {
	tpl, err := template.New("str").Parse(string(str))
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	if err = tpl.Execute(buf, v); err != nil {
		panic(err)
	}
	return T(buf.String())
}

func (a Attr) FromTemplate(v any) (o Attr, err error) {
	defer func() {
		f := recover()
		if f != nil {
			err = f.(error)
		}
	}()
	o = Attr{
		ID:            mustStrTemplate(a.ID, v),
		Name:          mustStrTemplate(a.Name, v),
		ArgName:       mustStrTemplate(a.ArgName, v),
		Type:          mustStrTemplate(a.Type, v),
		ArrayType:     mustStrTemplate(a.ArrayType, v),
		Access:        mustStrTemplate(a.Access, v),
		Default:       mustStrTemplate(a.Default, v),
		Required:      mustStrTemplate(a.Required, v),
		Report:        a.Report,
		Scene:         mustStrTemplate(a.Scene, v),
		Range:         mustStrTemplate(a.Range, v),
		MfCode:        mustStrTemplate(a.MfCode, v),
		ShowAs:        mustStrTemplate(a.ShowAs, v),
		ListSize:      mustStrTemplate(a.ListSize, v),
		Enumeration:   mustStrTemplate(a.Enumeration, v),
		Unit:          mustStrTemplate(a.Unit, v),
		Multiplier:    mustStrTemplate(a.Multiplier, v),
		Desc:          mustStrTemplate(a.Desc, v),
		Marshal:       mustStrTemplate(a.Marshal, v),
		Unmarshal:     mustStrTemplate(a.Unmarshal, v),
		MarshalNoType: a.MarshalNoType,
	}

	if a.Values != nil {
		o.Values = map[string]Name{}
		for vk, vv := range a.Values {
			o.Values[mustStrTemplate(vk, v)] = mustStrTemplate(vv, v)
		}
	}
	if a.Bits != nil {
		o.Bits = map[string]Name{}
		for vk, vv := range a.Values {
			o.Bits[mustStrTemplate(vk, v)] = mustStrTemplate(vv, v)
		}
	}
	for _, c := range a.Cond {
		o.Cond = append(o.Cond, Condition{
			Desc:     mustStrTemplate(c.Desc, v),
			Name:     mustStrTemplate(c.Name, v),
			Attr:     mustStrTemplate(c.Attr, v),
			Value:    mustStrTemplate(c.Value, v),
			Mask:     mustStrTemplate(c.Mask, v),
			Invert:   c.Invert,
			Operator: mustStrTemplate(c.Operator, v),
		})
	}
	for _, aa := range a.StructAttr {
		if av, err := aa.FromTemplate(v); err != nil {
			return o, err
		} else {
			o.StructAttr = append(o.StructAttr, av)
		}
	}
	return
}

func (c Command) FromTemplate(v any) (o Command, err error) {
	defer func() {
		f := recover()
		if f != nil {
			err = f.(error)
		}
	}()
	o = Command{
		ID:       mustStrTemplate(c.ID, v),
		Name:     mustStrTemplate(c.Name, v),
		Desc:     mustStrTemplate(c.Desc, v),
		Dir:      mustStrTemplate(c.Dir, v),
		Required: mustStrTemplate(c.Required, v),
		ShowAs:   mustStrTemplate(c.ShowAs, v),
		Vendor:   mustStrTemplate(c.Vendor, v),
		MfCode:   mustStrTemplate(c.MfCode, v),
	}
	if c.Response != nil {
		rs := mustStrTemplate(*c.Response, v)
		o.Response = &rs
	}

	for _, attr := range c.PayloadAttr {
		av, err := attr.FromTemplate(v)
		if err != nil {
			return o, err
		}
		o.PayloadAttr = append(o.PayloadAttr, av)
	}
	return
}

func (c CmdAttr) FromTemplate(v any) (o CmdAttr, err error) {
	for _, attr := range c.Attr {
		av, err := attr.FromTemplate(v)
		if err != nil {
			return o, err
		}
		o.Attr = append(o.Attr, av)
	}
	for _, as := range c.AttrSet {
		ao := AttrSet{
			ID:     "",
			Desc:   "",
			MfCode: "",
		}
		for _, attr := range as.Attr {
			av, err := attr.FromTemplate(v)
			if err != nil {
				return o, err
			}
			ao.Attr = append(ao.Attr, av)
		}
		o.AttrSet = append(o.AttrSet, ao)
	}
	for _, cmd := range c.Command {
		oc, err := cmd.FromTemplate(v)
		if err != nil {
			return o, err
		}
		o.Command = append(o.Command, oc)
	}
	return o, nil
}

func (t TemplateDef) Types() (out map[string]Attr, err error) {
	out = map[string]Attr{}
	defer func() {
		f := recover()
		if f != nil {
			err = f.(error)
		}
	}()
	for _, v := range t.Values {
		for k, vv := range t.TypeMap {
			kk := mustStrTemplate(k, v)
			out[kk], err = vv.FromTemplate(v)
			if err != nil {
				return
			}
		}
	}
	return
}

func (t TemplateDef) Clusters() (out []Cluster, err error) {
	if t.ClusterPerValue == nil {
		return
	}
	src := t.ClusterPerValue
	defer func() {
		f := recover()
		if f != nil {
			err = f.(error)
		}
	}()

	for _, v := range t.Values {
		var cl, srv CmdAttr
		if cl, err = src.Client.FromTemplate(v); err != nil {
			return
		}
		if srv, err = src.Server.FromTemplate(v); err != nil {
			return
		}

		out = append(out, Cluster{
			ID:     mustStrTemplate(src.ID, v),
			Name:   mustStrTemplate(src.Name, v),
			MfCode: mustStrTemplate(src.MfCode, v),
			Desc:   mustStrTemplate(src.Desc, v),
			Server: srv,
			Client: cl,
		})

	}
	return
}

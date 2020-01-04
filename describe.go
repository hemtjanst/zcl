package zcl

import (
	"reflect"
)

type ValueDescription struct {
	Index        int
	Name         string
	Type         TypeID
	List         bool
	ListType     *ValueDescription
	EnumValues   []Option
	BitmapValues []Option
}

type CommandDescription struct {
	Name      string
	ClusterID *ClusterID
	Required  bool
	MnfCode   *uint16
	List      bool
	Arguments []ValueDescription
}

func DescribeAttr(attr Argument) (desc ValueDescription) {
	ft := reflect.TypeOf(attr)

	desc.List = ft.Kind() == reflect.Slice
	if so, ok := ft.MethodByName("SingleOptions"); ok {
		opts := so.Func.Call([]reflect.Value{})
		for _, optVal := range opts {
			opt := optVal.Interface()
			if optOpt, ok := opt.(Option); ok {
				desc.EnumValues = append(desc.EnumValues, optOpt)
			}
		}
	}
	if so, ok := ft.MethodByName("MultiOptions"); ok {
		opts := so.Func.Call([]reflect.Value{})
		for _, optVal := range opts {
			opt := optVal.Interface()
			if optOpt, ok := opt.(Option); ok {
				desc.BitmapValues = append(desc.BitmapValues, optOpt)
			}
		}
	}
	return
}

func DescribeGeneral(cmd General) (desc CommandDescription) {
	rv := reflect.ValueOf(cmd)
	rf := reflect.TypeOf(cmd)
	if rf.Kind() == reflect.Ptr {
		rf = rf.Elem()
		rv = rv.Elem()
	}
	desc.Name = rf.Name()

	switch rf.Kind() {
	case reflect.Slice:
		//t := rf.Elem()
		desc.List = true
		desc.Arguments = []ValueDescription{{
			Index: -1,
			Name:  rf.Name(),
		}}
	case reflect.Struct:
		for i := 0; i < rf.NumField(); i++ {
			f := rf.Field(i)
			fv := rv.Field(i)
			//ft := f.Type

			fd := ValueDescription{}
			if attr, ok := fv.Interface().(Argument); ok {
				fd = DescribeAttr(attr)
			} else {

			}
			fd.Index = i
			fd.Name = f.Name
			desc.Arguments = append(desc.Arguments, fd)
		}
	}

	return
}

func DescribeCommand(cmd Command) (desc CommandDescription) {
	desc = DescribeGeneral(cmd)

	c := cmd.Cluster()
	desc.ClusterID = &c

	m := cmd.MnfCode()
	if len(m) == 2 {
		mnf := uint16(m[0]<<8 | m[1])
		desc.MnfCode = &mnf
	}
	desc.Required = cmd.Required()
	return
}

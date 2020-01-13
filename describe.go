package zcl

import (
	"log"
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
	Default      Val
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
	desc.Name = ft.Elem().Name()

	if ea, ok := attr.(EnumArg); ok {
		desc.EnumValues = ea.SingleOptions()
	} else if ea, ok := attr.(BitmapArg); ok {
		desc.BitmapValues = ea.MultiOptions()
	}
	desc.Type = attr.TypeID()
	desc.List = ft.Kind() == reflect.Slice
	return
}

func DescribeArgs(cmd interface{}) (v []ValueDescription) {
	rv := reflect.ValueOf(cmd)
	rf := reflect.TypeOf(cmd)
	if rf.Kind() == reflect.Ptr {
		rf = rf.Elem()
		rv = rv.Elem()
	}
	if rf.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < rf.NumField(); i++ {
		f := rf.Field(i)
		fv := rv.Field(i)
		//ft := f.Type

		fd := ValueDescription{}
		if attr, ok := fv.Interface().(Argument); ok {
			fd = DescribeAttr(attr)
		} else if attr, ok := fv.Addr().Interface().(Argument); ok {
			fd = DescribeAttr(attr)
		} else {
			log.Printf("Unable to describe argument %s of %s: %s %#v", f.Name, rf.Name(), fv.Type(), fv)
		}
		fd.Index = i
		fd.Name = f.Name
		v = append(v, fd)
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
		desc.Arguments = DescribeArgs(cmd)
	}

	return
}

func DescribeCommand(cmd Command) (desc CommandDescription) {
	desc = DescribeGeneral(cmd)

	c := cmd.Cluster()
	desc.ClusterID = &c

	m := cmd.MnfCode()
	if m > 0 {
		desc.MnfCode = &m
	}
	desc.Required = cmd.Required()
	return
}

func DescribeZdoCommand(cmd ZdoCommand) (desc CommandDescription) {
	rf := reflect.TypeOf(cmd)
	if rf.Kind() == reflect.Ptr {
		rf = rf.Elem()
	}

	c := cmd.Cluster()
	desc = CommandDescription{
		Name:      rf.Name(),
		Arguments: DescribeArgs(cmd),
		ClusterID: &c,
	}
	return
}

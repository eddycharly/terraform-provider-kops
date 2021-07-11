package main

import (
	"reflect"

	"k8s.io/apimachinery/pkg/util/sets"
)

type optionsDoc struct {
	header string
	footer string
}

type options struct {
	noSchema     bool
	version      *int
	exclude      sets.String
	rename       map[string]string
	nullable     sets.String
	required     sets.String
	computed     sets.String
	computedOnly sets.String
	forceNew     sets.String
	sensitive    sets.String
	suppressDiff sets.String
	doc          *optionsDoc
}

func newOptions() *options {
	return &options{
		exclude:      sets.NewString(),
		rename:       make(map[string]string),
		nullable:     sets.NewString(),
		required:     sets.NewString(),
		computed:     sets.NewString(),
		computedOnly: sets.NewString(),
		forceNew:     sets.NewString(),
		sensitive:    sets.NewString(),
		suppressDiff: sets.NewString(),
	}
}

func noSchema() func(o *options) {
	return func(o *options) {
		o.noSchema = true
	}
}

func version(v int) func(o *options) {
	return func(o *options) {
		o.version = &v
	}
}

func exclude(excluded ...string) func(o *options) {
	return func(o *options) {
		o.exclude.Insert(excluded...)
	}
}

func rename(old, new string) func(o *options) {
	return func(o *options) {
		o.rename[old] = new
	}
}

func required(required ...string) func(o *options) {
	return func(o *options) {
		o.required.Insert(required...)
	}
}

func nullable(required ...string) func(o *options) {
	return func(o *options) {
		o.nullable.Insert(required...)
	}
}

func computed(computed ...string) func(o *options) {
	return func(o *options) {
		o.computed.Insert(computed...)
	}
}

func computedOnly(computedOnly ...string) func(o *options) {
	return func(o *options) {
		o.computedOnly.Insert(computedOnly...)
	}
}

func forceNew(forceNew ...string) func(o *options) {
	return func(o *options) {
		o.forceNew.Insert(forceNew...)
	}
}

func sensitive(sensitive ...string) func(o *options) {
	return func(o *options) {
		o.sensitive.Insert(sensitive...)
	}
}

func suppressDiff(suppressDiff ...string) func(o *options) {
	return func(o *options) {
		o.suppressDiff.Insert(suppressDiff...)
	}
}

func doc(header, footer string) func(o *options) {
	return func(o *options) {
		o.doc = &optionsDoc{
			header: header,
			footer: footer,
		}
	}
}

func (o *options) verify(t reflect.Type) error {
	if err := verifyFields(t, o.exclude.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.nullable.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.required.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.computed.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.computedOnly.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.forceNew.List()...); err != nil {
		return err
	}
	if err := verifyFields(t, o.suppressDiff.List()...); err != nil {
		return err
	}
	for k := range o.rename {
		if err := verifyFields(t, k); err != nil {
			return err
		}
	}
	return nil
}

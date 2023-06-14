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
	exclude      sets.Set[string]
	rename       map[string]string
	nullable     sets.Set[string]
	required     sets.Set[string]
	computed     sets.Set[string]
	computedOnly sets.Set[string]
	forceNew     sets.Set[string]
	sensitive    sets.Set[string]
	suppressDiff sets.Set[string]
	doc          *optionsDoc
}

func newOptions() *options {
	return &options{
		exclude:      sets.New[string](),
		rename:       make(map[string]string),
		nullable:     sets.New[string](),
		required:     sets.New[string](),
		computed:     sets.New[string](),
		computedOnly: sets.New[string](),
		forceNew:     sets.New[string](),
		sensitive:    sets.New[string](),
		suppressDiff: sets.New[string](),
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

func doc(header, footer string) func(o *options) {
	return func(o *options) {
		o.doc = &optionsDoc{
			header: header,
			footer: footer,
		}
	}
}

func (o *options) verify(t reflect.Type) error {
	if err := verifyFields(t, sets.List(o.sensitive)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.exclude)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.nullable)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.required)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.computed)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.computedOnly)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.forceNew)...); err != nil {
		return err
	}
	if err := verifyFields(t, sets.List(o.suppressDiff)...); err != nil {
		return err
	}
	for k := range o.rename {
		if err := verifyFields(t, k); err != nil {
			return err
		}
	}
	return nil
}

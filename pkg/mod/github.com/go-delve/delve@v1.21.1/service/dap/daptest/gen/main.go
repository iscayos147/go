// Binary gen generates service/dap/daptest/responses.go.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/types"
	"os"
	"text/template"

	"golang.org/x/tools/go/packages"

	_ "github.com/google/go-dap"
)

var oFlag = flag.String("o", "", "output file name")

var tmpl = template.Must(template.New("assert").Parse(`package daptest

// Code generated by go generate; DO NOT EDIT.
// The code generator program is in ./gen directory.

import (
	"testing"

	"github.com/google/go-dap"
)
{{range .}}
// Expect{{.}} reads a protocol message from the connection
// and fails the test if the read message is not *{{.}}.
func (c *Client) Expect{{.}}(t *testing.T) *dap.{{.}} {
	t.Helper()
	m := c.ExpectMessage(t)
	return c.Check{{.}}(t, m)
}

// Check{{.}} fails the test if m is not *{{.}}.
func (c *Client) Check{{.}}(t *testing.T, m dap.Message) *dap.{{.}} {
	t.Helper(){{if or (or (eq . "StepInResponse") (eq . "StepOutResponse")) (eq . "NextResponse") }}
	_, ok := m.(*dap.ContinuedEvent)
	if !ok {
		t.Fatalf("got %#v, want *dap.ContinuedEvent", m)
	}
	m = c.ExpectMessage(t){{else}}{{if (eq . "ConfigurationDoneResponse") }}
	oe, ok := m.(*dap.OutputEvent)
	if !ok {
		t.Fatalf("got %#v, want *dap.OutputEvent", m)
	}
	if oe.Body.Output != "Type 'dlv help' for list of commands.\n" {
		t.Fatalf("got %#v, want Output=%q", m, "Type 'dlv help' for list of commands.\n")
	}
	m = c.ExpectMessage(t){{end}}{{end}}
	r, ok := m.(*dap.{{.}})
	if !ok {
		t.Fatalf("got %#v, want *dap.{{.}}", m)
	}
	return r
}{{end}}
`))

func main() {
	flag.Parse()

	if *oFlag == "" {
		fmt.Fprintf(os.Stderr, "-o must be provided\n")
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes,
	}, "github.com/google/go-dap")
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if len(pkgs) != 1 || pkgs[0].Types == nil {
		fmt.Fprintf(os.Stderr, "invalid package was loaded: %#v\n", pkgs)
		os.Exit(1)
	}

	messages := []string{}
	scope := pkgs[0].Types.Scope()
	for _, name := range scope.Names() {
		// Find only types that are embedding go-dap.Response message.
		obj := scope.Lookup(name)
		if !obj.Exported() {
			continue // skip unexported
		}
		u, ok := obj.Type().Underlying().(*types.Struct)
		if !ok {
			continue
		}
		for i := 0; i < u.NumFields(); i++ {
			if f := u.Field(i); f.Embedded() && (f.Type().String() == "github.com/google/go-dap.Response" ||
				f.Type().String() == "github.com/google/go-dap.Event") {
				messages = append(messages, obj.Name())
				break
			}
		}
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, messages); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate: %v\n", err)
		os.Exit(1)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Generated invalid go code: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*oFlag, formatted, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write: %v\n", err)
		os.Exit(1)
	}
}

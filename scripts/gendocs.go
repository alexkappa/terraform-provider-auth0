// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/alexkappa/terraform-provider-auth0/auth0"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var args = struct {
	provider,
	providerName,
	resource string
}{}

func init() {
	flag.StringVar(&args.provider, "provider", "auth0", "Provider key")
	flag.StringVar(&args.providerName, "provider-name", "Auth0", "Provider friendly name")
	flag.StringVar(&args.resource, "resource", "", "Resource key")
	flag.Parse()
}

func main() {
	buf := bytes.NewBuffer([]byte{})
	p := auth0.Provider()
	r := &Resource{
		ProviderKey:    args.provider,
		ProviderName:   args.providerName,
		ResourceKey:    args.resource,
		ResourceSchema: p.ResourcesMap[args.resource],
	}
	r.GenerateResourceMarkdown(buf)
	fmt.Print(buf.String())
}

type Resource struct {
	ProviderKey  string
	ProviderName string

	ResourceKey    string
	ResourceSchema *schema.Resource
}

func (r *Resource) GenerateResourceMarkdown(wr io.Writer) error {
	rd := r.resourceDocsFromSchema(r.ResourceSchema, nil, false)
	return resourceDocsTemplate.Execute(wr, rd)
}

func (r *Resource) resourceDocsFromSchema(res *schema.Resource, docs *ResourceDocs, isNested bool) *ResourceDocs {
	if docs == nil {
		docs = &ResourceDocs{
			ProviderKey:        r.ProviderKey,
			ProviderName:       r.ProviderName,
			ResourceKey:        r.ResourceKey,
			MarkdownHeaderFunc: markdownHeader,
			Fields:             make(map[string]*schema.Schema),
			NestedFields:       make(map[string]map[string]*schema.Schema),
		}
	}

	for name, s := range res.Schema {
		if v, isResource := s.Elem.(*schema.Resource); isResource {
			docs.NestedFields[name] = v.Schema
			log.Printf("Processing nested field: %q", name)
			r.resourceDocsFromSchema(v, docs, true)
		}
		if _, isSchema := s.Elem.(*schema.Schema); isSchema {
			log.Printf("Nested Schema is not implemented (yet) - SKIPPING %q", name)
		}

		if !isNested {
			log.Printf("Processing primitive field: %q", name)
			docs.Fields[name] = s
		}
	}

	return docs
}

func markdownHeader(header string) string {
	return strings.Replace(header, "_", "\\_", 0)
}

type ResourceDocs struct {
	ProviderKey  string
	ProviderName string
	ResourceKey  string

	Fields       map[string]*schema.Schema
	NestedFields map[string]map[string]*schema.Schema

	MarkdownHeaderFunc func(s string) string
}

var resourceDocsTemplate = template.Must(template.New("resource-docs").Parse(`---
layout: "{{.ProviderKey}}"
page_title: "{{.ProviderName}}: {{.ResourceKey}}"
description: |-
  TODO
---

# {{call .MarkdownHeaderFunc .ResourceKey}}

TODO

## Example Usage

` + "```hcl" + `
resource "{{.ResourceKey}}" "example" {
  // TODO
}
` + "```" + `

## Argument Reference

The following arguments are supported:

{{range $key, $schema := .Fields}}{{if or $schema.Optional $schema.Required }}
* ` + "`{{ $key }}`" + ` - {{if $schema.Required}}(Required){{else}}(Optional){{end}} {{ $schema.Description }}
{{- end}}{{end}}
{{- if gt (len .NestedFields) 0}}
{{- range $fieldName, $nestedFields := .NestedFields}}
### ` + "`{{ $fieldName }}`" + `
{{range $key, $schema := $nestedFields}}{{if or $schema.Optional $schema.Required }}
* ` + "`{{ $key }}`" + ` - {{if $schema.Required}}(Required){{else}}(Optional){{end}} {{ $schema.Description }}.

{{- end}}{{- end}}
{{range $key, $schema := $nestedFields}}{{if and $schema.Computed (not $schema.Optional)}}

* ` + "`{{ $key }}`" + ` - {{ $schema.Description }}
{{- end}}{{- end -}}
{{end}}
{{end}}
## Attributes Reference
In addition to the arguments listed above, the following computed attributes are
exported:
{{range $key, $schema := .Fields}}{{if and $schema.Computed (not $schema.Optional)}}
* ` + "`{{ $key }}`" + ` - {{ $schema.Description }}
{{end}}{{end}}
## Import
{{.ResourceKey}} can be imported using the , e.g.
` + "```" + `
$ terraform import {{.ResourceKey}}.example ...
` + "```" + `
`))

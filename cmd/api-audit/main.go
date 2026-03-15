// Package main provides an API surface audit tool for GopherFrame.
// It analyzes the public API and generates a report for community review.
//
// Usage:
//
//	go run cmd/api-audit/main.go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fmt.Println("# GopherFrame API Surface Audit Report")
	fmt.Println()
	fmt.Println("Generated for community feedback. Please review the public API")
	fmt.Println("and submit feedback via GitHub Discussions.")
	fmt.Println()

	// Scan root package
	rootExports := scanPackage(".")
	fmt.Println("## Root Package (`gopherframe`)")
	fmt.Println()
	printExports(rootExports)

	// Scan expr package
	exprExports := scanPackage("./pkg/expr")
	fmt.Println("## Expression Package (`pkg/expr`)")
	fmt.Println()
	printExports(exprExports)

	// Scan core package
	coreExports := scanPackage("./pkg/core")
	fmt.Println("## Core Package (`pkg/core`)")
	fmt.Println()
	printExports(coreExports)

	// Summary
	total := len(rootExports) + len(exprExports) + len(coreExports)
	fmt.Printf("\n## Summary\n\n")
	fmt.Printf("- **Root package**: %d exported symbols\n", len(rootExports))
	fmt.Printf("- **expr package**: %d exported symbols\n", len(exprExports))
	fmt.Printf("- **core package**: %d exported symbols\n", len(coreExports))
	fmt.Printf("- **Total public API surface**: %d symbols\n", total)
	fmt.Println()
	fmt.Println("## Feedback Template")
	fmt.Println()
	fmt.Println("Please copy this template to a GitHub Discussion:")
	fmt.Println()
	fmt.Println("```")
	fmt.Println("### API Feedback")
	fmt.Println()
	fmt.Println("**Symbol**: [function/type name]")
	fmt.Println("**Category**: [naming|signature|missing|redundant|documentation]")
	fmt.Println("**Suggestion**: [your suggestion]")
	fmt.Println("**Rationale**: [why this change would improve the API]")
	fmt.Println("```")
}

type export struct {
	Name    string
	Kind    string // "func", "type", "method", "const", "var"
	Pkg     string
	Comment string
}

func scanPackage(dir string) []export {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not parse %s: %v\n", dir, err)
		return nil
	}

	var exports []export
	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			_ = filename
			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.FuncDecl:
					if d.Name.IsExported() {
						e := export{
							Name: d.Name.Name,
							Kind: "func",
						}
						if d.Recv != nil && len(d.Recv.List) > 0 {
							e.Kind = "method"
							recvType := exprToString(d.Recv.List[0].Type)
							e.Name = recvType + "." + d.Name.Name
						}
						if d.Doc != nil {
							e.Comment = strings.TrimSpace(d.Doc.Text())
							if len(e.Comment) > 80 {
								e.Comment = e.Comment[:80] + "..."
							}
						}
						exports = append(exports, e)
					}
				case *ast.GenDecl:
					for _, spec := range d.Specs {
						switch s := spec.(type) {
						case *ast.TypeSpec:
							if s.Name.IsExported() {
								exports = append(exports, export{
									Name: s.Name.Name,
									Kind: "type",
								})
							}
						case *ast.ValueSpec:
							for _, name := range s.Names {
								if name.IsExported() {
									kind := "var"
									if d.Tok == token.CONST {
										kind = "const"
									}
									exports = append(exports, export{
										Name: name.Name,
										Kind: kind,
									})
								}
							}
						}
					}
				}
			}
		}
	}

	sort.Slice(exports, func(i, j int) bool {
		if exports[i].Kind != exports[j].Kind {
			return exports[i].Kind < exports[j].Kind
		}
		return exports[i].Name < exports[j].Name
	})
	return exports
}

func printExports(exports []export) {
	currentKind := ""
	for _, e := range exports {
		if e.Kind != currentKind {
			currentKind = e.Kind
			fmt.Printf("### %ss\n\n", strings.ToUpper(currentKind[:1])+currentKind[1:])
		}
		if e.Comment != "" {
			fmt.Printf("- `%s` — %s\n", e.Name, e.Comment)
		} else {
			fmt.Printf("- `%s`\n", e.Name)
		}
	}
	fmt.Println()
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return exprToString(e.X)
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	default:
		return "?"
	}
}

// CreateAPIFeedbackIssueTemplate generates a GitHub issue template for API feedback.
func CreateAPIFeedbackIssueTemplate() string {
	return `---
name: API Feedback
about: Provide feedback on GopherFrame's public API
title: "[API] "
labels: api-feedback
---

## API Symbol
<!-- Which function, type, or method are you providing feedback on? -->

## Category
<!-- Choose one: naming | signature | missing | redundant | documentation -->

## Current Behavior
<!-- How does the API currently work? -->

## Suggested Change
<!-- What would you change? -->

## Rationale
<!-- Why would this change improve the API? -->

## Example
<!-- Show before/after code if applicable -->
`
}

func init() {
	// Write issue template if running with --write-template
	for _, arg := range os.Args[1:] {
		if arg == "--write-template" {
			tmpl := CreateAPIFeedbackIssueTemplate()
			dir := filepath.Join(".github", "ISSUE_TEMPLATE")
			_ = os.MkdirAll(dir, 0750)
			_ = os.WriteFile(filepath.Join(dir, "api_feedback.md"), []byte(tmpl), 0600)
			fmt.Println("Wrote .github/ISSUE_TEMPLATE/api_feedback.md")
		}
	}
}

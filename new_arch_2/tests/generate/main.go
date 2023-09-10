package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const (
	GenerateDecorator = "@generate-test"
)

func main() {
	dirName, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	//dirName := "/Users/alonmuroch/Projects/ssv-experiments/new_arch_2/tests/spec/asgard/qbft/full_flow/"
	fmt.Printf("Generating tests for dir: %s\n", dirName)
	packageName, funcNames, instanceNames := parseFuncNamesToGenerate(dirName)

	if len(funcNames) > 0 {
		newF, err := os.Create(filepath.Join(dirName, "generated_test.go"))
		if err != nil {
			panic(err.Error())
		}
		f := createTestFunc(packageName, funcNames, instanceNames)
		fset := token.NewFileSet() // positions are relative to fset
		if err := printer.Fprint(newF, fset, f); err != nil {
			panic(err.Error())
		}
	}
}

func parseFuncNamesToGenerate(dirPath string) (string, []string, []string) {
	entires, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err.Error())
	}

	funcNames := make([]string, 0)
	instanceNames := make([]string, 0)
	packageName := ""
	for _, e := range entires {
		if !e.IsDir() && e.Name() != "generate.go" {
			file := filepath.Join(dirPath, e.Name())
			fset := token.NewFileSet() // positions are relative to fset
			f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
			if err != nil {
				panic(err.Error())
			}

			for _, decl := range f.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					if funcDecl.Doc != nil {
						for _, c := range funcDecl.Doc.List {
							if strings.Contains(c.Text, GenerateDecorator) {
								funcNames = append(funcNames, funcDecl.Name.Name)

								// find return type
								for _, l1 := range funcDecl.Body.List {
									if l2, ok := l1.(*ast.ReturnStmt); ok {
										name := "*" + l2.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit).Type.(*ast.SelectorExpr).X.(*ast.Ident).Name
										name += "."
										name += l2.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit).Type.(*ast.SelectorExpr).Sel.Name
										instanceNames = append(instanceNames, name)
									}
								}
							}

							// set package name
							packageName = f.Name.Name
						}
					}
				}
			}
		}
	}
	return packageName, funcNames, instanceNames
}

func createTestFunc(packageName string, funcNames, testInstances []string) *ast.File {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", template, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	f.Name.Name = packageName

	for i, funcName := range funcNames {
		fmt.Printf("    Adding func: %s\n", funcName)
		addTest(f, funcName, testInstances[i])
	}

	addAllTestsSlice(f, funcNames)

	fmt.Printf("Found %d tests\n", len(funcNames))

	return f
}

func addTest(input *ast.File, funcName, testInstance string) {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", singleTestTemplate, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	// change Test function name
	f.Decls[0].(*ast.FuncDecl).
		Name.Name = "Test" + funcName

	// change test generate function name
	f.Decls[0].(*ast.FuncDecl).
		Body.List[0].(*ast.AssignStmt).
		Rhs[0].(*ast.CallExpr).
		Args[0].(*ast.CallExpr).
		Fun.(*ast.Ident).Name = funcName

	f.Decls[0].(*ast.FuncDecl).
		Body.List[0].(*ast.AssignStmt).
		Rhs[0].(*ast.CallExpr).
		Fun.(*ast.IndexExpr).
		Index.(*ast.Ident).Name = testInstance

	f.Decls[0].(*ast.FuncDecl).
		Body.List[0].(*ast.AssignStmt).
		Rhs[0].(*ast.CallExpr).
		Args[0].(*ast.CallExpr).
		Fun.(*ast.Ident).Obj.Name = testInstance

	input.Decls = append(input.Decls, f.Decls[0])
}

func addAllTestsSlice(input *ast.File, funcName []string) {
	// prepare test func names
	toReplace := ""
	for _, n := range funcName {
		toReplace += n + "(),"
	}
	tmplat := strings.Replace(strings.Clone(testsArrayTemplate), "__to_replace__", toReplace, -1)

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", tmplat, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	input.Decls = append(input.Decls, f.Decls[1])
}

var template = `
package xxx
import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/qbft"
	"ssv-experiments/new_arch_2/tests/spec/asgard/ssv"
	"testing"
)
`
var singleTestTemplate = `
package xxx

func XXX(t *testing.T) {
	tst, err := tests.NewSpecTest[XXX](XXX())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}
`
var testsArrayTemplate = `
package XXX 

import "ssv-experiments/new_arch_2/tests"

var AllTests = []tests.TestObject{
	__to_replace__
}
`

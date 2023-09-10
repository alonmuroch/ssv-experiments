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
	packageName, funcNames := parseFuncNamesToGenerate(dirName)

	if len(funcNames) > 0 {
		newF, err := os.Create("/Users/alonmuroch/Projects/ssv-experiments/new_arch_2/tests/spec/asgard/qbft/full_flow/generated_test.go")
		if err != nil {
			panic(err.Error())
		}
		f := createTestFunc(packageName, funcNames)
		fset := token.NewFileSet() // positions are relative to fset
		if err := printer.Fprint(newF, fset, f); err != nil {
			panic(err.Error())
		}
	}
}

func parseFuncNamesToGenerate(dirPath string) (string, []string) {
	entires, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err.Error())
	}

	funcNames := make([]string, 0)
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
							}

							// set package name
							packageName = f.Name.Name
						}
					}
				}
			}
		}
	}
	return packageName, funcNames
}

func createTestFunc(packageName string, funcNames []string) *ast.File {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", template, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	f.Name.Name = packageName

	for _, funcName := range funcNames {
		fmt.Printf("    Adding func: %s\n", funcName)
		addTest(f, funcName)
	}

	addAllTestsSlice(f, funcNames)

	fmt.Printf("Found %d tests\n", len(funcNames))

	return f
}

func addTest(input *ast.File, funcName string) {
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
	"testing"
)
`
var singleTestTemplate = `
package xxx

func XXX(t *testing.T) {
	tst, err := tests.NewSpecTest[*qbft.ProcessMessageTest](XXX())
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

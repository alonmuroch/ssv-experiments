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

type parsedTest struct {
	funcName         string
	testInstanceName string
}

type parsedPackage struct {
	packageName string
	tests       []*parsedTest
	imports     map[string]*ast.ImportSpec
}

// GenerateTests will generate tests file for all tests in working directory
func GenerateTests() {
	dirName, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	//dirName := "/Users/alonmuroch/Projects/ssv-experiments/new_arch_2/tests/spec/asgard/qbft/full_flow/"

	fmt.Printf("Generating tests for dir: %s\n", dirName)
	parsedPkg := parseFuncNamesToGenerate(dirName)

	// write generated test
	newF, err := os.Create(filepath.Join(dirName, "generated_test.go"))
	if err != nil {
		panic(err.Error())
	}
	f := createTestFunctions(parsedPkg)
	fset := token.NewFileSet() // positions are relative to fset
	if err := printer.Fprint(newF, fset, f); err != nil {
		panic(err.Error())
	}

	// write all tests
	newF, err = os.Create(filepath.Join(dirName, "all_tests.go"))
	if err != nil {
		panic(err.Error())
	}
	f = createAllTests(parsedPkg)
	fset = token.NewFileSet() // positions are relative to fset
	if err := printer.Fprint(newF, fset, f); err != nil {
		panic(err.Error())
	}
}

func parseFuncNamesToGenerate(dirPath string) *parsedPackage {
	entires, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err.Error())
	}

	ret := &parsedPackage{
		tests:   []*parsedTest{},
		imports: map[string]*ast.ImportSpec{},
	}
	for _, e := range entires {
		if !e.IsDir() && e.Name() != "generate.go" {
			file := filepath.Join(dirPath, e.Name())
			fset := token.NewFileSet() // positions are relative to fset
			f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
			if err != nil {
				continue // some non golang file
			}

			for _, decl := range f.Decls {
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					if funcDecl.Doc != nil {
						for _, c := range funcDecl.Doc.List {
							if strings.Contains(c.Text, GenerateDecorator) {
								parsedTst := &parsedTest{
									funcName: funcDecl.Name.Name,
								}

								// find return type
								for _, l1 := range funcDecl.Body.List {
									if l2, ok := l1.(*ast.ReturnStmt); ok {
										parsedTst.testInstanceName = "*" + l2.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit).Type.(*ast.SelectorExpr).X.(*ast.Ident).Name
										parsedTst.testInstanceName += "."
										parsedTst.testInstanceName += l2.Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit).Type.(*ast.SelectorExpr).Sel.Name
									}
								}
								ret.tests = append(ret.tests, parsedTst)
							}

							// add imports
							for _, imprt := range f.Imports {
								ret.imports[imprt.Path.Value] = imprt
							}

							// set package name
							ret.packageName = f.Name.Name
						}
					}
				}
			}
		}
	}

	fmt.Printf("Found %d tests\n", len(ret.tests))
	return ret
}

// createAllTests creates an AllTests var for all the tests in the package
func createAllTests(parsedPackage *parsedPackage) *ast.File {
	toReplace := ""
	for i := range parsedPackage.tests {
		toReplace += parsedPackage.tests[i].funcName + "(),"
	}
	tmplat := strings.Replace(strings.Clone(testsArrayTemplate), "__to_replace__", toReplace, -1)

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", tmplat, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	f.Name.Name = parsedPackage.packageName

	return f
}

// createTestFunctions creates runnable unit tests for each spec test
func createTestFunctions(parsedPackage *parsedPackage) *ast.File {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", template, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	// add imports
	for _, imprt := range parsedPackage.imports {
		f.Imports = append(f.Imports, imprt)
		f.Decls[0].(*ast.GenDecl).Specs = append(f.Decls[0].(*ast.GenDecl).Specs, imprt)
	}

	// change packcage name
	f.Name.Name = parsedPackage.packageName

	// add tests
	for i := range parsedPackage.tests {
		fmt.Printf("    Adding func: %s\n", parsedPackage.tests[i].funcName)
		addTest(f, parsedPackage.tests[i].funcName, parsedPackage.tests[i].testInstanceName)
	}

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

var template = `
package xxx
import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
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

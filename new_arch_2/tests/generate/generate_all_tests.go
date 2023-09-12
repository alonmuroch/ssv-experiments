package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const (
	FirstImportPos            = 30 + 34 + StandardSpacingPostImport
	StandardSpacingPostImport = 2
)

type generateTestInfo struct {
	importPath  *ast.ImportSpec
	packageName string
}

// GenerateAllTestsFile generates a unified file with all "AllTests" from all packages, used to generate ssz tests
func GenerateAllTestsFile() {
	currentDirPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	testsMap := map[string]*generateTestInfo{}
	basePath := filepath.Join(currentDirPath, "../")
	startingPackagePath := "ssv-experiments/new_arch_2/tests/spec"
	processDirectory(basePath, startingPackagePath, "asgard", FirstImportPos, testsMap)

	// prepare template
	toReplace := ""
	for pkgName, imprt := range testsMap {
		toReplace += "\"" + pkgName + "\":" + imprt.importPath.Name.Name + ".AllTests,"
	}
	tmplat := strings.Replace(strings.Clone(allTestsTemplate), "__to_replace__", toReplace, -1)

	// read template
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", tmplat, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	// change package name
	f.Name.Name = "main"

	// add imports
	for _, imprt := range testsMap {
		f.Imports = append(f.Imports, imprt.importPath)
		f.Decls[0].(*ast.GenDecl).Specs = append(f.Decls[0].(*ast.GenDecl).Specs, imprt.importPath)
	}

	// write
	newF, err := os.Create(filepath.Join("/Users/alonmuroch/Projects/ssv-experiments/new_arch_2/tests/generate", "all_tests.go"))
	if err != nil {
		panic(err.Error())
	}
	fset = token.NewFileSet() // positions are relative to fset
	if err := printer.Fprint(newF, fset, f); err != nil {
		panic(err.Error())
	}
}

func processGeneratedTest(basePath, startingPackagePath, packagePath string, pos int) (*generateTestInfo, int) {
	importName := strings.Replace(packagePath, "/", "_", -1)
	return &generateTestInfo{
		importPath: &ast.ImportSpec{
			Name: &ast.Ident{
				NamePos: token.Pos(pos),
				Name:    importName,
			},
			Path: &ast.BasicLit{
				ValuePos: token.Pos(pos + len(importName) + 1),
				Value:    "\"" + filepath.Join(startingPackagePath, packagePath) + "\"",
			},
		},
		packageName: importName,
	}, pos + len(importName) + 1 + StandardSpacingPostImport
}

// processDirectory is a recursive function that iterates on all directories (starting from packagePath). Returns a map between package path and generated test info
func processDirectory(basePath, startingPackagePath, packagePath string, pos int, testsMap map[string]*generateTestInfo) {
	entires, err := os.ReadDir(filepath.Join(basePath, startingPackagePath, packagePath))
	if err != nil {
		panic(err.Error())
	}

	for _, e := range entires {
		if e.IsDir() {
			processDirectory(basePath, startingPackagePath, filepath.Join(packagePath, e.Name()), pos, testsMap)
		} else if e.Name() == "all_tests.go" {
			testInfo, nextPos := processGeneratedTest(basePath, startingPackagePath, packagePath, pos)
			testsMap[testInfo.packageName] = testInfo
			pos = nextPos
		}
	}
}

var allTestsTemplate = `
package XXX 

import "ssv-experiments/new_arch_2/tests"

// AllTests maps package name to tests
var AllTests = map[string][]tests.TestObject{
	__to_replace__
}`

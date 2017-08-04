package cprofile

import (
	"errors"
)

// Program is a mass of code
type Program struct {
	pkgs map[string]*Package
	pkg  *Package
}

func newProgram(pkgs map[string]*Package, pkg *Package) *Program {
	return &Program{pkgs, pkg}
}

// Imports returns an array of all packages imported by the program
func (p *Program) Imports() []*Package {
	if p == nil || len(p.pkgs) == 0 {
		return []*Package{}
	}

	list := p.flatten()
	return list
}

// Package returns the top-level package for this program
func (p *Program) Package() (*Package, error) {
	if p == nil || p.pkg == nil {
		return nil, errors.New("Program is not loaded")
	}

	return p.pkg, nil
}

func (p *Program) flatten() []*Package {
	l := len(p.pkgs)
	if l == 0 {
		return []*Package{}
	}
	pkgs := make([]*Package, l)
	i := 0
	for _, v := range p.pkgs {
		pkgs[i] = v
		i++
	}

	return pkgs
}

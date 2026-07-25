package ast_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
)

func TestDetachedIdentifierHasNoAssignmentDeclaration(t *testing.T) {
	t.Parallel()
	identifier := ast.NewNodeFactory(ast.NodeFactoryHooks{}).NewIdentifier("value")
	if target := ast.GetAssignmentTarget(identifier); target != nil {
		t.Fatalf("GetAssignmentTarget() = %v, want nil", target)
	}
	if name := ast.GetNameOfDeclaration(identifier); name != nil {
		t.Fatalf("GetNameOfDeclaration() = %v, want nil", name)
	}
}

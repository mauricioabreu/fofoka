package node_test

import (
	"testing"

	"github.com/mauricioabreu/fofoka/node"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}

var _ = Describe("Node", func() {
	Describe("Adding a neighbor", func() {
		It("adds the node in the list", func() {
			n := node.New(1)
			n.AddNeighbor(&node.Node{ID: 2})
			Expect(len(n.Neighbors)).To(Equal(1))
		})
	})
})

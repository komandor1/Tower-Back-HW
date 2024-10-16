package main
import "fmt"

func main() {
	var bst BinarySearchTree
    bst.Add(5)
	bst.Add(2)
	bst.Add(51)
	bst.Add(0)
	bst.Add(8)

    fmt.Println(bst.IsExist(5))
	bst.Delete(5)
	fmt.Println(bst.IsExist(5))
	fmt.Println(bst.root)
	fmt.Println(bst.IsExist(51))
	bst.Delete(51)
	fmt.Println(bst.IsExist(51))
	fmt.Println(bst.root.rightNode)
}

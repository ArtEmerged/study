package main

import (
	"fmt"
)

type TreeNode struct {
	parent     int
	priceAll   int
	companyAll map[string]bool
	index      int
	children   []*TreeNode
}

// var arr []*TreeNode

func main() {
	var root TreeNode
	isSetRoot := true

	var numberOfTree int
	var numberOfCompany int

	fmt.Scanf("%d %d", &numberOfTree, &numberOfCompany)

	var companies []string
	for i := 0; i < numberOfCompany; i++ {
		var temp string
		fmt.Scan(&temp)
		companies = append(companies, temp)
	}
	for i := 0; i < numberOfTree; i++ {
		var parent int
		var price int
		var company string

		fmt.Scanf("%d %d %s", &parent, &price, &company)
		companyAll := map[string]bool{
			company: true,
		}
		node := TreeNode{
			priceAll:   price,
			parent:     parent,
			companyAll: companyAll,
			index:      i + 1,
		}

		if isSetRoot {
			root = node
			isSetRoot = false
		} else {
			root.addNode(&node)
		}
	}

	fmt.Println(companies)

	printTree(&root, 0)

	// findNodeWithChildren(&root)
	// for _, v := range arr {
	// 	fmt.Println("O:", v.companyAll, "index:", v.index)
	// }
	// var allPrices []int
	// for _, node := range arr {
	// 	temp := make([]string, len(arr))
	// 	copy(temp, companies)

	// 	if haveTheCompaniesInNode(node, &temp) {
	// 		allPrices = append(allPrices, getPriceOfNode(node))
	// 	}
	// }
	// fmt.Println(allPrices)
}

func (r *TreeNode)


func (r *TreeNode) addNode(node *TreeNode) {
	if r.index == node.parent {
		r.children = append(r.children, node)
	}

	for _, child := range r.children {
		child.addNode(node)
	}
}

func printTree(node *TreeNode, depth int) {
	// Print the current node
	fmt.Printf("%s- %s: $%d\n", getIndentation(depth), node.companyAll, node.price)

	// Recursively print children
	for _, child := range node.children {
		printTree(child, depth+1)
	}
}

func getIndentation(depth int) string {
	indentation := ""
	for i := 0; i < depth; i++ {
		indentation += "  " // You can adjust the indentation as needed
	}
	return indentation
}




// func getPriceOfNode(root *TreeNode) int {
// 	if len(root.children) == 0 {
// 		return root.price
// 	}
// 	sum := 0
// 	for _, node := range root.children {
// 		sum += getPriceOfNode(node)
// 	}

// 	return sum + root.price
// }

// func haveTheCompaniesInNode(root *TreeNode, arr *[]string) bool {
// 	index := inArr(root.company, *arr)
// 	if index != -1 {
// 		temp := *arr
// 		temp = append(temp[:index], temp[index+1:]...)
// 		*arr = temp
// 		if len(*arr) == 0 {
// 			return true
// 		}
// 	}
// 	flag := false
// 	for _, node := range root.children {
// 		flag = haveTheCompaniesInNode(node, arr)
// 		if flag {
// 			return true
// 		}
// 	}
// 	return false
// }

// func inArr(match string, arr []string) int {
// 	for i, str := range arr {
// 		if str == match {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func findNodeWithChildren(r *TreeNode) {
// 	if len(r.children) != 0 {
// 		arr = append(arr, r)
// 		for i := range r.children {
// 			findNodeWithChildren(r.children[i])
// 		}
// 	}
// }
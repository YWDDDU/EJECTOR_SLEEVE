package main
import "fmt"

func mole_choose() int {
	// 选择类型并返回int值
	var choose int
	for 1 {
		fmt.Printf("\t0:自身司筒\t\n")
		fmt.Printf("\t1:有托司筒\t\n")
		fmt.Printf("\t2:扁顶针\t\n")
		fmt.Printf("\t3:Exit\t\n")
		fmt.Printf("\t选择类型: ")
		fmt.Scanln(&choose)
		if choose != 0 || choose != 1 || choose != 2 || choose != 3 {break}
	}
	return choose
}

func material_choose() int {
	// 选择材质并返回int值
	var choose int
	for 1 {
		fmt.Printf("\t0:SKD_61\t\n")
		fmt.Printf("\t1:H13\t\n")
		fmt.Printf("\t2:普通\t\n")
		fmt.Printf("\t选择类型: ")
		fmt.Scanln(&choose)
		if choose != 0 || choose != 1 || choose != 2 {break}
	}
	return choose
}

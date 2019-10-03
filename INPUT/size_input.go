package main
import "fmt"

func size_input(i int) string {
	var size string
	switch i {
	case 0:
		fmt.Printf("\t你的选择是自身司筒\t\n")
		fmt.Printf("\t请按照此形式输入尺寸:内径*外径*外筒长度*内针长度\t\n\t")
	case 1:
		fmt.Printf("\t你的选择是有托司筒\t\n")
		fmt.Printf("\t请按照此形式输入尺寸:内径*外径*外筒长度*内针长度*托位直径*托位长度\t\n\t")
	case 2:
		fmt.Printf("\t你的选择是扁顶针\t\n")
		fmt.Printf("\t请按照此形式输入尺寸:方位*扁位*总长*托位直径*托位长度\t\n\t")
	default:
		fmt.Printf("\tERROR\t\n")
	}
	fmt.Scanln(&size)
	return size
}

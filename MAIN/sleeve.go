// 
package main
import (
	"fmt"
)

func main() {
	var molec int
	var straight_ejector_sleeve = 0		//自身司筒
	var stepped_ejector_sleeve = 1		//有托司筒
	var flat_ejector_pin = 2			//扁顶针

	molec = mole_choose()		//类型选择
	if molec == straight_ejector_sleeve {
			fmt.Println(size_input0())	//规格输入
			//normal()		//普通材质
	} else if molec == stepped_ejector_sleeve {
			fmt.Println(size_input1())	//规格输入
			//skd_61()		//SKD_61材质
	} else if molec == flat_ejector_pin {
			fmt.Println(size_input2())	//规格输入
			//h13()			//H13材质
	} else if molec == 3 {
			return
	}
}

func mole_choose() int {
	var choose int
	// 选择类型并返回int值
	fmt.Printf("\t0:自身司筒\t\n")
	fmt.Printf("\t1:有托司筒\t\n")
	fmt.Printf("\t2:扁顶针\t\n")
	fmt.Printf("\t3:exit\t\n")
	// 输入数字并选择, 默认值为0, 类型转换
	fmt.Printf("\t选择类型: ")
	fmt.Scanln(&choose)
	return choose
}

func size_input0() string {
	//自身司筒规格
	//var size [4]float32
	var size0 string
	//var inside_diameter float32		//司筒内径
	//var outside_diameter float32	//司筒外径
	//var length0 float32				//内针长度
	//var length1 float32				//外筒长度
	//var format = "%f*%f*%f*%f"
	fmt.Printf("\t你的选择是自身司筒\t\n")
	fmt.Printf("\t请按照此形式输入尺寸:内径*外径*外筒长度*内针长度\t\n\t")
	fmt.Scanln(&size0)
	//fmt.Sscanf(size, format, &inside_diameter, &outside_diameter, &length1, &length0)
	return size0
}

func size_input1() string {
	//有托司筒规格
	var size1 string
	fmt.Printf("\t你的选择是有托司筒\t\n")
	fmt.Printf("\t请按照此形式输入尺寸:内径*外径*外筒长度*内针长度*托位直径*托位长度\t\n\t")
	fmt.Scanln(&size1)
	return size1
}

func size_input2() string {
	//扁顶针规格
	var size2 string
	fmt.Printf("\t你的选择是扁顶针\t\n")
	fmt.Printf("\t请按照此形式输入尺寸:方位*扁位*总长*托位直径*托位长度\t\n\t")
	fmt.Scanln(&size2)
	return size2
}
















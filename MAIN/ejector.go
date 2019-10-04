package main
import (
	"fmt"
)

func main() {
	var molec int						// 类型
	var material int					// 材质

	/*var (								//类型选择值
		straight_ejector_sleeve = 0,	//自身司筒
		stepped_ejector_sleeve = 1,		//有托司筒
		flat_ejector_pin = 2,			//扁顶针
	)

	var (								//材质选择值
		normail = 2,					//普通材质
		skd61 = 0,						//国产SKD61
		h13 = 1,						//进口SKD61
	)*/

	var ejector_size string				//司筒尺寸
	var (
		money_in float32				//进货价
	)

	for true {
		molec = mole_choose()				//类型选择
		if molec == 3 {break}
		material = material_choose()		//材质选择

		ejector_size = size_input(molec)	//返回尺寸

		money_in = ejector_math_in(ejector_size, material)
		fmt.Printf("\t%s\t\n\t进货价: %.2f\t\n", ejector_size, money_in)
	}

}

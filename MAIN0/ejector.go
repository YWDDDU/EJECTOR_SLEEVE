package main
import (
	"fmt"
)

func main() {
	var molec int						// 类型
	var material int					// 材质
	var straight_ejector_sleeve = 0		//自身司筒
	var stepped_ejector_sleeve = 1		//有托司筒
	var flat_ejector_pin = 2			//扁顶针
	var normal = 0						//普通材质
	var skd61 = 1						//国产SKD61
	var h13 = 3							//进口SKD61
	var ejector_size string				//司筒尺寸

	molec = mole_choose()				//类型选择
	material = material_choose()		//材质选择

	ejector_size = size_input(molec)	//返回尺寸

}
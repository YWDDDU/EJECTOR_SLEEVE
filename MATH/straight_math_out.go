package main
import (
	"fmt"
)

func straight_math_out (size string, material int) float32 float32 {
	var normail_num_out [12][2] float32			//普通材质
	var skd61_num_out [12][2] float32			//国产SKD61
	var choose_num_out [][] float32				//材质选择
	normal_num_out = {
		{0.08, 0.11},
		{0.05, 0.08},
		{0.06, 0.09},
		{0.07, 0.11},
		{0.09, 0.13},
		{0.09, 0.15},
		{0.13, 0.16},
		{0.18, 0.25},
		{0.21, 0.30},
		{0.27, 0.40},
		{0.35, 0.50},
		{0.43, 0.57},
	}
	skd61_num_out = {
		{0.11, 0.16},
		{0.07, 0.11},
		{0.08, 0.12},
		{0.09, 0.14},
		{0.11, 0.16},
		{0.12, 0.19},
		{0.16, 0.20},
		{0.23, 0.32},
		{0.26, 0.37},
		{0.33, 0.47},
		{0.40, 0.57},
		{0.48, 0.67},
	}

	var straight_format = "%f*%f*%f*%f"			//自身司筒格式
	var money_out0 float32						//建议售价0
	var money_out1 float32						//建议售价1

	var (
		inside_diameter float32					//内径
		outside_diameter float32				//外径
		length0 float32							//外筒长度
		length1 float32							//内针长度
	)


	//从输入读取规格
	fmt.Sscanf(size, straight_format, &inside_diameter, &outside_diameter, &length0, &length1)

	// 材质 --> 单价
	if material == 0 {
		//材质为国产61
		choose_num_out = skd61_num_out
	} else if material == 2 {
		//材质为普通
		choose_num_out = normail_num_out
	}

	i := (outside_diameter - inside_diameter) / 2	//司筒单边

	if inside_



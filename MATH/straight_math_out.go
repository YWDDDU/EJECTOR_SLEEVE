package main
import (
	"fmt"
)

func straight_math_out (size string, material int) float32{
	//var normail_num_out [12][2] float32			//普通材质
	//var skd61_num_out [12][2] float32				//国产SKD61
	var choose_num_out [][] float32					//材质选择
	//var choose_num_out [12][2]*float32				//材质选择--指针

	var normal_num_out = [12][2]float32{			//普通材质
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
	var skd61_num_out = [12][2]float32{				//国产SKD61
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
	//var money_out1 float32						//建议售价1

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
		//choose_num_out = skd61_num_out
		for i := 0; i < 12; i++ {
			for j := 0; j < 2; j++ {
				choose_num_out[i][j] = skd61_num_out[i][j]
			}
		}
	} else if material == 2 {
		//材质为普通
		//choose_num_out = normal_num_out
		for i := 0; i < 12; i++ {
			for j := 0; j < 2; j++ {
				choose_num_out[i][j] = normal_num_out[i][j]
			}
		}
	}

	i := (outside_diameter - inside_diameter) / 2	//司筒单边


	//建议零售价0
	if inside_diameter >= 1.0 && inside_diameter < 1.5 && outside_diameter <= 5.0 {
		money_out0 = money_out_math(choose_num_out[0][0], choose_num_out[0][1], length0, i)
	}

	if inside_diameter >= 1.5 && inside_diameter < 4.0 && outside_diameter <= 7.0 {
		money_out0 = money_out_math(choose_num_out[1][0], choose_num_out[1][1], length0, i)
	}

	if inside_diameter >= 4.0 && inside_diameter < 5.0 && outside_diameter <= 9.0 {
		money_out0 = money_out_math(choose_num_out[2][0], choose_num_out[2][1], length0, i)
	}

	if inside_diameter >= 5.0 && inside_diameter < 6.0 && outside_diameter <= 9.0 {
		money_out0 = money_out_math(choose_num_out[3][0], choose_num_out[3][1], length0, i)
	}

	if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 10.0 {
		money_out0 = money_out_math(choose_num_out[4][0], choose_num_out[4][1], length0, i)
	} else if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 12.0 {
		money_out0 = money_out_math(choose_num_out[5][0], choose_num_out[5][1], length0, i)
	}

	if inside_diameter >= 8.0 && inside_diameter < 10.0 && outside_diameter <= 15.0 {
		money_out0 = money_out_math(choose_num_out[6][0], choose_num_out[6][1], length0, i)
	}

	if inside_diameter >= 10.0 && inside_diameter < 12.0 && outside_diameter <= 18.0 {
		money_out0 = money_out_math(choose_num_out[7][0], choose_num_out[7][1], length0, i)
	}

	if inside_diameter >= 12.0 && inside_diameter < 14.0 && outside_diameter <= 20.0 {
		money_out0 = money_out_math(choose_num_out[8][0], choose_num_out[8][1], length0, i)
	}

	if inside_diameter >= 14.0 && inside_diameter < 16.0 && outside_diameter <= 22.0 {
		money_out0 = money_out_math(choose_num_out[9][0], choose_num_out[9][1], length0, i)
	}

	if inside_diameter >= 16.0 && inside_diameter < 18.0 && outside_diameter <= 25.0 {
		money_out0 = money_out_math(choose_num_out[10][0], choose_num_out[10][1], length0, i)
	}

	if inside_diameter >= 18.0 && inside_diameter < 20.0 && outside_diameter <= 28.0 {
		money_out0 = money_out_math(choose_num_out[11][0], choose_num_out[11][1], length0, i)
	}

	//建议零售价1
	
	return money_out0
}

func money_out_math(ml, mm, long, i float32) float32 {
	var singe float32										//单价
	if i >= 0.8 {
		singe = long * ml
	} else if i >= 0.6 && i < 0.8 {
		singe = long * mm
	}
	return singe
}
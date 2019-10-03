package main
import (
	"fmt"
)

func straight_math_in (size string, material int) float32 {
	var normal_num [12][2] float32			//普通材质
	var skd61_num [12][2] float32			//国产SKD61
	var choose_num [][] float32				//材质选择
	normal_num = {
		{0.06, 0.09},
		{0.035, 0.06},
		{0.045, 0.07},
		{0.05, 0.09},
		{0.07, 0.11},
		{0.075, 0.13},
		{0.11, 0.14},
		{0.15, 0.23},
		{0.18, 0.28},
		{0.24, 0.38},
		{0.32, 0.48},
		{0.40, 0.55},
	}
	skd61_num = {
		{0.09, 0.14},
		{0.05, 0.09},
		{0.06, 0.10},
		{0.07, 0.12},
		{0.09, 0.14},
		{0.10, 0.17},
		{0.14, 0.18},
		{0.20, 0.30},
		{0.23, 0.35},
		{0.30, 0.45},
		{0.38, 0.55},
		{0.45, 0.65},
	}

	var straight_format = "%f*%f*%f*%f"			//自身司筒格式
	var money_in float32						//进货价

	var (
		inside_diameter float32					//内径
		outside_diameter float32				//外径
		length0 float32							//外筒长度
		length1 float32							//内针长度
	)

	//从输入读取规格
	fmt.Sscanf(size, straight_format, &inside_diameter, &outside_diameter, &length0, &length1)

	// 材质 -->  单价
	if material == 0 {
		// 材质为国产SKD-61
		choose_num = skd61_num
	} else if material == 2 {
		// 材质为普通
		choose_num = normal_num
	}

	i := (outside_diameter - inside_diameter) / 2		//司筒单边

	if inside_diameter >= 1.0 && inside_diameter < 1.5 && outside_diameter <= 5.0 {
		money_in = money_math(choose_num[0][0], choose_num[0][1], length0, i)
	}

	if inside_diameter >= 1.5 && inside_diameter < 4.0 && outside_diameter <= 7.0 {
		money_in = money_math(choose_num[1][0], choose_num[1][1], length0, i)
	}

	if inside_diameter >= 4.0 && inside_diameter < 5.0 && outside_diameter <= 9.0 {
		money_in = money_math(choose_num[2][0], choose_num[2][1], length0, i)
	}

	if inside_diameter >= 5.0 && inside_diameter < 6.0 && outside_diameter <= 9.0 {
		money_in = money_math(choose_num[3][0], choose_num[3][1], length0, i)
	}

	if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 10.0 {
		money_in = money_math(choose_num[4][0], choose_num[4][1], length0, i)
	} else if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 12.0 {
		money_in = money_math(choose_num[5][0], choose_num[5][1], length0, i)
	}

	if inside_diameter >= 8.0 && inside_diameter < 10.0 && outside_diameter <= 15.0 {
		money_in = money_math(choose_num[6][0], choose_num[6][1], length0, i)
	}

	if inside_diameter >= 10.0 && inside_diameter < 12.0 && outside_diameter <= 18.0 {
		money_in = money_math(choose_num[7][0], choose_num[7][1], length0, i)
	}

	if inside_diameter >= 12.0 && inside_diameter < 14.0 && outside_diameter <= 20.0 {
		money_in = money_math(choose_num[8][0], choose_num[8][1], length0, i)
	}

	if inside_diameter >= 14.0 && inside_diameter < 16.0 && outside_diameter <= 22.0 {
		money_in = money_math(choose_num[9][0], choose_num[9][1], length0, i)
	}

	if inside_diameter >= 16.0 && inside_diameter < 18.0 && outside_diameter <= 25.0 {
		money_in = money_math(choose_num[10][0], choose_num[10][1], length0, i)
	}

	if inside_diameter >= 18.0 && inside_diameter < 20.0 && outside_diameter <= 28.0 {
		money_in = money_math(choose_num[11][0], choose_num[11][1], length0, i)
	}

	return money_in
}

func money_math(ml, mm, long, i float32) float32 {
	var singe float32										//单价
	if i >= 0.8 {
		singe = long * ml
	} else if i >= 0.6 && i < 0.8 {
		singe = long * mm
	}
	return singe
}

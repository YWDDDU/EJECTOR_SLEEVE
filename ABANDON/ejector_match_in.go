package main
import (
	"fmt"
)

/* test
func main() {
	var size string
	var money float32
	fmt.Scanln(&size)
	money = ejector_match_in(size)
	fmt.Println(size)
	fmt.Println(money)
}*/

func ejector_match_in (size string) float32 {
	var straight_format = "%f*%f*%f*%f"		//自身司筒格式
	var money float32
	var (
		inside_diameter float32			//内径
		outside_diameter float32		//外径
		length0 float32					//外筒长度
		length1 float32					//内针长度
	)
	fmt.Sscanf(size, straight_format, &inside_diameter, &outside_diameter, &length0, &length1)

	var ml, mm float32			//ml: 大于0.8 ; mm: 0.6-0.8
	i := (outside_diameter - inside_diameter) / 2

	//if 1 <= inside_diameter < 1.5 && outside_diameter <= 5.0 {
	if inside_diameter >= 1.0 && inside_diameter < 1.5 && outside_diameter <= 5.0 {
		ml = 0.06
		mm = 0.09
		money = money_out(ml, mm, length0, i)
		/*if i >= 0.8 {
			money = length0 * ml
		} else if 0.6 <= i < 0.8 {
			money = length0 * mm
		}*/
	}

	//if 1.5 <= inside_diameter < 4.0 && outside_diameter <= 7.0 {
	if inside_diameter >= 1.5 && inside_diameter < 4.0 && outside_diameter <= 7.0 {
		ml = 0.035
		mm = 0.06
		money = money_out(ml, mm, length0, i)
	}

	//if 4.0 <= inside_diameter < 5.0 && outside_diameter <= 9.0 {
	if inside_diameter >= 4.0 && inside_diameter < 5.0 && outside_diameter <= 9.0 {
		ml = 0.045
		mm = 0.07
		money = money_out(ml, mm, length0, i)
	}

	//if 5.0 <= inside_diameter < 6.0 && outside_diameter <= 9.0 {
	if inside_diameter >= 5.0 && inside_diameter < 6.0 && outside_diameter <= 9.0 {
		ml = 0.05
		mm = 0.09
		money = money_out(ml, mm, length0, i)
	}

	//if 6.0 <= inside_diameter < 8.0 && outside_diameter <= 10.0 {
	if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 10.0 {
		ml = 0.07
		mm = 0.11
		money = money_out(ml, mm, length0, i)
	//} else if 6.0 <= inside_diameter < 8.0 && outside_diameter <= 12.0 {
	} else if inside_diameter >= 6.0 && inside_diameter < 8.0 && outside_diameter <= 12.0 {
		ml = 0.075
		mm = 0.13
		money = money_out(ml, mm, length0, i)
	}

	//if 8.0 <= inside_diameter < 10.0 && outside_diameter <= 15.0 {
	if inside_diameter >= 8.0 && inside_diameter < 10.0 && outside_diameter <= 15.0 {
		ml = 0.11
		mm = 0.14
		money = money_out(ml, mm, length0, i)
	}
	return money
}

func money_out (ml, mm, long, i float32) float32 {
	var num float32
	if i >= 0.8 {
		num = long * ml
	//} else if 0.6 <= i < 0.8 {
	} else if i >= 0.6 && i < 0.8 {
		num = long * mm
	}
	return num
}

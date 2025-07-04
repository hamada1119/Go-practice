package main

import "fmt"

func main() {
	// var i int = 0;
	// for {
	// 	i++;
	// 	if i == 5{
	// 		break;
	// 	}
	// }
	// fmt.Println(i)
	// fmt.Println("Loop")

	// count := 0;
	// for count <=10 {
	// 	fmt.Println(count)
	// 	count++;
	// }
	// fmt.Println(count);

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	arr := [3]int{1, 2, 4}
	// for i :=0; i<len(arr); i++{
	// 	fmt.Println(arr[i]);
	// }

	// for i, v := range arr {
	// 	fmt.Println(i,v)
	// }

	for _, v := range arr {
		fmt.Println(v)
	}

	m := map[string]int{"1": 100, "2": 200}
	for _, i := range m {
		fmt.Println(i)
	}

	for i := range m {
		fmt.Println(i)
	}
}

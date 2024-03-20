package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 4

	fmt.Println(slice)

	slice = append(slice[0:i], slice[i+1:]...) // если важен порядок
	fmt.Println(slice)

	slice = remove(slice, 0) // если порядок не важен
	fmt.Println(slice)
	nam := RemoveIndex(slice, 1) //если важно сохранить изначачльный слайс без изменений
	fmt.Println(nam)
	fmt.Println(slice)
	nam = RemoveCopy(slice, 0)
	fmt.Println(nam)
	fmt.Println(slice)
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1] // на место i ставим последний
	return s[:len(s)-1]
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func RemoveCopy(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

package main

//配列の重複を除去する関数
func distinct(arr []string) []string {
	m := make(map[string]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	var newArr []string
	for k, _ := range m {
		newArr = append(newArr, k)
	}
	return newArr
}

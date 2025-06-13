package main

import "fmt"

/*
map
map_variable := make(map[KeyType]ValueType, initialCapacity)
*/
func main() {

	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"]
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}

	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("删除元素后地图")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

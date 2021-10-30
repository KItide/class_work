package main



func m_type(i interface{}) {
	switch i.(type) {
	case string:
		print("这个数据的类型为string\n")
	case int:
		print("这个数据的类型为int\n")
	case bool:
		print("这个数据的类型为bool\n")
	case float64:
		print("这个数据的类型float64\n")
	case float32:
		print("这个数据的类型为flaot32\n")
	}

	return
}
func main() {

	m_type(1)
	m_type("你好吗？？")
	m_type(true)
	m_type(0.3262)

}

package main

func main() {
	var x,y = 10,20
	x = y
	println("= ", x)

	x = 15
	x += y //35 = 15 + 20
	println("+= ", x)

	x = 30
	x -= y //10 = 30 - 20
	println("-= ", x)

	x = 5
	x *= y //100 = 5 * 20
	println("*= ", x)

	x = 100
	x /= y //x = 100 / 20
	println("/= ", x)

	x = 100
	x %= y //x = 100 % 20
	println("%= ", x)
}
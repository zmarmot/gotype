package gotype

func Abs(x int)int{
    if x < 0{
        x = -x
    }
    return x
}

func Min(a,b int)int{
    if a > b{
        return b
    }
    return a
}

func Max(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func Min3(a,b,c int)int{
    t := Min(a,b)
    return Min(t,c)
}

func Max3(a,b,c int)int{
    t := Max(a,b)
    return Max(t,c)
}

//判断数字n的二进制数从右往左数第i位是否为1
func IsOne(n, i int) bool {
	return (uint(n) & (uint(1) << uint(i))) == 1
}

func SwapInt(data []int, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}
func SwapRune(data []rune, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}
func IsNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

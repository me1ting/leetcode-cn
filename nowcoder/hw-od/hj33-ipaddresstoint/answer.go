package hj33ipaddresstoint

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(IpToInt(scanner.Text()))
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Println(IntToIp(n))
}

func IpToInt(ip string) int {
	slices := strings.Split(ip, ".")
	r := 0
	for i, nStr := range slices {
		n, _ := strconv.Atoi(nStr)
		r |= (n << ((3 - i) * 8))
	}
	return r
}

func IntToIp(n int) string {
	slices := []string{}
	for i := 0; i < 4; i++ {
		field := (n >> ((3 - i) * 8)) & 0xff
		slices = append(slices, strconv.Itoa(field))
	}

	return strings.Join(slices, ".")
}

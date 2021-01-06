package main

// https://leetcode-cn.com/problems/add-strings/
// 简单题，不需要任何技巧
func addStrings(num1 string, num2 string) string {
	i, j, k := len(num1)-1, len(num2)-1, len(num1)

	if i < j {
		k = len(num2)
	}

	carry := 0
	sum := make([]byte, k+1)

	for i >= 0 && j >= 0 {
		bit := int(num1[i]) + int(num2[j]) - 2*'0' + carry
		if bit >= 10 {
			bit = bit % 10
			carry = 1
		} else {
			carry = 0
		}
		sum[k] = byte(bit + '0')
		i--
		j--
		k--
	}

	if i != 0 {
		for i >= 0 {
			bit := int(num1[i]) - '0' + carry
			if bit >= 10 {
				bit = bit % 10
				carry = 1
			} else {
				carry = 0
			}
			sum[k] = byte(bit + '0')
			i--
			k--
		}
	}

	if j != 0 {
		for j >= 0 {
			bit := int(num2[j]) - '0' + carry
			if bit >= 10 {
				bit = bit % 10
				carry = 1
			} else {
				carry = 0
			}
			sum[k] = byte(bit + '0')
			j--
			k--
		}
	}

	if carry == 1 {
		sum[k] = '1'
		k--
	}

	return string(sum[k+1:])
}

func main() {
	addStrings("9", "99")
}

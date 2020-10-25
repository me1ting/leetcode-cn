package long_pressed_name

// https://leetcode-cn.com/problems/long-pressed-name/
func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	if name[0] != typed[0] {
		return false
	}

	var i,j int
	for i, j = 0, 0; i < len(name) && j < len(typed); {
		if name[i] != typed[j] {
			if typed[j] != typed[j-1] {
				return false
			}

			j++
			for j < len(typed) && typed[j-1] == typed[j] {
				j++
			}

			if j == len(typed) || name[i] != typed[j] {
				return false
			}
		}
		i++
		j++
	}

	if i!=len(name) {
		return false
	}

	if j!=len(typed) {
		for j < len(typed) && typed[j-1] == typed[j] {
			j++
		}
		if j != len(typed) {
			return false
		}
	}
	return true
}

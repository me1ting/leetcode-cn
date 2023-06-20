package asteroidcollision

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		if asteroid > 0 {
			stack = append(stack, asteroid)
			continue
		}

		alive := true

		for alive && len(stack) > 0 {
			n := len(stack)
			top := stack[n-1]

			if top < 0 {
				break
			}

			sum := top + asteroid
			if sum == 0 {
				stack = stack[:n-1]
				alive = false
			} else if sum > 0 {
				alive = false
			} else {
				stack = stack[:n-1]
			}
		}

		if alive {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

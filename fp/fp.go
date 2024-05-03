package fp

func SimulateTrafficLight(start string, iterations int) []string {
	var lights []string
	var state = start

	for i := 0; i < iterations; i++ {
		lights = append(lights, state)

		switch state {
		case "green":
			state = "yellow"
		case "yellow":
			state = "red"
		case "red":
			state = "green"
		}
	}

	return lights
}

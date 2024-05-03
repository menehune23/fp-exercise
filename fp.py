from typing import List


def simulate_traffic_light(start: str, iterations: int) -> List[str]:
    lights = []
    state = start

    for _ in range(iterations):
        lights.append(state)

        if state == 'green':
            state = 'yellow'
        elif state == 'yellow':
            state = 'red'
        elif state == 'red':
            state = 'green'
    
    return lights

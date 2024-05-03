from typing import List
from itertools import accumulate


# This is a pure function
def next_state(prev: str) -> str:
    if prev == 'green':
        return 'yellow'
    
    if prev == 'yellow':
        return 'red'
    
    if prev == 'red':
        return 'green'
    
    return 'invalid'



def simulate_traffic_light(start: str, iterations: int) -> List[str]:
    ## Middle-ground: extracts most state management to pure function
    # state = start
    # lights = []
    # 
    # for _ in range(iterations):
    #     lights.append(state)
    #     state = next_state(state)
    # 
    # return lights

    # Most "functional"
    return list(accumulate(range(iterations-1), lambda state, _: next_state(state), initial=start))

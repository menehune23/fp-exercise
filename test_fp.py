from fp import simulate_traffic_light


def test_simulate_traffic_lights():
    assert simulate_traffic_light('red', 5) == ['red', 'green', 'yellow', 'red', 'green']
    
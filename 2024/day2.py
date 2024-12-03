# Advent of Code 2024
# Day 2
# Ed Salisbury
# 2024-12-02


def is_valid(levels, max_dist=3):

    trend = 0

    # Turn the levels into pairs of numbers and iterate over them
    # [(83, 84), (84, 87), (87, 90)...]
    for prev, curr in zip(levels, levels[1:]):

        # Check the distance between the pairs
        # If 0, the numbers are the same and it is invalid
        # If greater than MAX_DIST, the distance is too far, so it's invalid
        dist = abs(curr - prev)
        if dist == 0 or dist > max_dist:
            return False

        # Check the trend - it should only go up or down, not change midstream
        current_trend = curr - prev
        if trend * current_trend < 0:
            return False

        # Update the trend
        trend = current_trend

    # The report is valid if it passes the above checks
    return True


def is_valid_with_dampener(levels, max_dist=3):
    # The "dampener" allows for one item to be incorrect within the list and it can still be valid
    # This will iterate through and remove a single item to see if it might be valid
    for i in range(len(levels)):
        if is_valid(levels[:i] + levels[i + 1 :], max_dist):
            return True

    return False


with open("day2.txt", "r") as file:
    safe = 0
    safe_with_dampener = 0

    for line in file:
        # Get levels, i.e. 83 84 87 90 91 94 96 98
        levels = list(map(int, line.split()))

        # Check to see if the line is valid as-is
        safe += is_valid(levels)

        # Check to see if it's valid with the dampener
        safe_with_dampener += 1 if is_valid(levels) else is_valid_with_dampener(levels)

    print(f"The number of safe reports is {safe}.")
    print(f"The number of safe reports with dampener is {safe_with_dampener}.")

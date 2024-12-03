# Advent of Code 2024
# Day 3
# Ed Salisbury
# 2024-12-03
import re

# Process each instruction and calculate the totals
with open("day3.txt", "r") as file:
    data = file.read()

    # Capture do(), don't(), or mul(X,Y)
    instructions = re.findall(r"do\(\)|don't\(\)|mul\(\d+,\d+\)", data)

    first_total = 0
    second_total = 0
    enabled = True

    for instruction in instructions:
        if instruction == "don't()":
            enabled = False
        elif instruction == "do()":
            enabled = True
        else:
            # Split mul(XX,YY) into left and right digits for multiplying
            left, right = map(int, instruction[4:-1].split(","))
            # Always add the product for the first total
            first_total += left * right
            if enabled:
                # Only add the product to the second total if it's enabled
                second_total += left * right

    print(f"The total of all of the multiplications is: {first_total}")
    print(f"The total of all of the multiplications with modifiers is: {second_total}")

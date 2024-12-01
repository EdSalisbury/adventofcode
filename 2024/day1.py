# Advent of Code 2024
# Day 1
# Ed Salisbury
# 2024-12-01

# Read a list of two columns of numbers as strings from the input file and split them into two lists (left and right)
# 'map(int, line.split())' converts the space-separated strings into integers
# 'zip(*...)' transposes the list of pairs, giving you two separate lists: left and right
with open("day1.txt", "r") as file:
    left, right = zip(*[map(int, line.split()) for line in file])

# Sort both lists to make sure we can pair the smallest numbers first (and so on)
left = sorted(left)
right = sorted(right)

# Ensure that both lists are of the same size before proceeding
if len(left) != len(right):
    raise Exception("Left and right lists are different sizes!")

# Calculate the total "difference" by iterating over the paired elements from both lists
# The zip function combines the two sorted lists element-wise and the absolute difference is computed
diff = sum(abs(l - r) for l, r in zip(left, right))

# Get the count for each item in the right list
right_count = {}
for num in right:
    right_count[num] = right_count.get(num, 0) + 1

# Get the similarities between the left and right lists
sim = 0
for num in left:
    sim += num * right_count.get(num, 0)

# Output the difference and similarity totals
print(f"The total difference between the lists is {diff}.")
print(f"The total similarity between the lists is {sim}.")

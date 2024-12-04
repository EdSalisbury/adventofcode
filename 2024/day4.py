# Advent of Code 2024
# Day 4
# Ed Salisbury
# 2024-12-04


def search(grid, row, col, word, pos, row_dir, col_dir):
    if col >= len(grid[0]) or row >= len(grid) or col < 0 or row < 0:
        return False
    if pos == len(word) - 1 and grid[row][col] == word[pos]:
        return True

    if grid[row][col] == word[pos]:
        return search(
            grid, row + row_dir, col + col_dir, word, pos + 1, row_dir, col_dir
        )
    return False


def xsearch(grid, row, col, word):
    if col >= len(grid[0]) - 1 or row >= len(grid) - 1 or col <= 0 or row <= 0:
        return False

    middle = word[int(len(word) / 2)]
    first = word[0]
    last = word[-1]

    x = grid[row][col]
    ul = grid[row - 1][col - 1]
    ur = grid[row + 1][col - 1]
    ll = grid[row - 1][col + 1]
    lr = grid[row + 1][col + 1]

    if x == middle:
        if ul == first and lr == last and ur == first and ll == last:
            return True
        if ul == last and lr == first and ur == last and ll == first:
            return True
        if ul == first and lr == last and ur == last and ll == first:
            return True
        if ul == last and lr == first and ur == first and ll == last:
            return True

    return False


with open("day4.txt", "r") as file:
    lines = map(str.strip, file.readlines())
    grid = [list(line) for line in lines]

    width = len(grid[0])
    height = len(grid)

    word = "XMAS"
    rev_word = word[::-1]
    found = 0
    xfound = 0

    # Assumes a square matrix
    for row in range(height):
        for col in range(width):
            found += search(grid, row, col, word, 0, 0, 1)  # Forward horizontal
            found += search(grid, row, col, rev_word, 0, 0, 1)  # Backward horizontal
            found += search(grid, row, col, word, 0, 1, 0)  # Forward vertical
            found += search(grid, row, col, rev_word, 0, 1, 0)  # Backward vertical
            found += search(grid, row, col, word, 0, 1, 1)  # Forward diagonal (down)
            found += search(
                grid, row, col, rev_word, 0, 1, 1
            )  # Backward diagonal (down)
            found += search(grid, row, col, word, 0, 1, -1)  # Forward diagonal (up)
            found += search(
                grid, row, col, rev_word, 0, 1, -1
            )  # Backward diagonal (down)

            # Look for X-MAS (X) pattern
            xfound += xsearch(grid, row, col, "MAS")

    print(f"The word {word} was found {found} times.")
    print(f"The X-MAS pattern was found {xfound} times.")

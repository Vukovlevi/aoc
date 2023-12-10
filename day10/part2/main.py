# Puzzle:            https://adventofcode.com/2023/day/10
# Video explanation: https://youtu.be/YiX9clrJBXA

FILE = 'test.txt'

# 1. part - How many steps along the loop does it take to get from the starting 
#           position to the point farthest from the starting position?
NEXT_PIPE = {                  # (dx, dy)
    '|': {( 0,  1), ( 0, -1)}, # down, up
    '-': {( 1,  0), (-1,  0)}, # right, left
    'L': {( 0, -1), ( 1,  0)}, # up, right
    'J': {( 0, -1), (-1,  0)}, # up, left
    '7': {(-1,  0), ( 0,  1)}, # left, down
    'F': {( 0,  1), ( 1,  0)}, # down, right
}

with open(FILE) as f:
    GRID = [list(line.strip()) for line in f.readlines()]
    DIST = {}
    start = None

    for y, row in enumerate(GRID):
        if 'S' in row:
            start = (row.index('S'), y)
            break

    GRID[start[1]][start[0]] = 'F' # <-- magic constant
    dist = 0
    pipe = start
    while pipe not in DIST:
        DIST[pipe] = dist
        dist += 1

        x, y = pipe
        c = GRID[y][x]

        for dx, dy in NEXT_PIPE[c]:
            nx, ny = x + dx, y + dy
            if (nx, ny) not in DIST:
                pipe = (nx, ny)
                break

    print(dist // 2)

    # 2. part - How many tiles are enclosed by the loop?
    tile_count = 0
    for y, row in enumerate(GRID):
        parity = 0
        for x, c in enumerate(row):
            if (x, y) not in DIST: # ground or junk pipe
                if parity % 2 == 1:
                    tile_count += 1
                continue

            # if c in ['|', 'F', '7']: # can be also used instead of following line
            if c in ['|', 'L', 'J']:   # L---J and F----7 do not increase parity
                parity += 1            # L---7 and F----J do increase parity

    print(tile_count)
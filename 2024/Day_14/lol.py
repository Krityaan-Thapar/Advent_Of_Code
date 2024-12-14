data = []
with open("input.txt", 'r') as f:
    data = f.read().strip('\n').split('\n')

robots = []
for line in data:
    p_str, v_str = line.split()
    px, py = [int(x) for x in p_str[2:].split(',')]
    vx, vy = [int(x) for x in v_str[2:].split(',')]

    robots.append([(px, py), (vx, vy)])

def calc_togetherness(a):
    togetherness = 0
    for i in range(len(a)):
        for j in range(len(a[0])):
            if a[i][j] != '#':
                continue
            for dx, dy in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
                nx, ny = (i + dx, j + dy)
                if 0 <= nx < len(a) and 0 <= ny < len(a[0]) and a[nx][ny] == '#':
                    togetherness += 1
                    break
    return togetherness

max_togetherness = 0
seconds = 0
for i in range(10000):
    image = [['.' for _ in range(101)] for _ in range(103)]
    for robot in robots:
        px, py = robot[0]
        image[py][px] = '#'
        vx, vy = robot[1]
        nx = (px + vx + 101) % 101
        ny = (py + vy + 103) % 103
        robot[0] = (nx, ny)
    togetherness = calc_togetherness(image)
    max_togetherness = max(max_togetherness, togetherness)
    if togetherness == max_togetherness:
        for line in image:
            print(''.join(line))
        print(seconds)
    seconds += 1
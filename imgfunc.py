from PIL import Image

def get_coordinate(path):
    """将图片转换成位图矩阵"""
    im = Image.open(path)
    l = list(im.getdata())
    size_x = im.size[0]
    res = []
    for i in range(len(l)):
        if l[i] != (255, 255, 255, 255):
            x = i % size_x
            y = i // size_x
            res.append([x, y])
    return res


def cut_coordinate(cod: list[list]):
    """裁剪图片的函数"""
    x0 = min([i[0] for i in cod])
    y0 = min([i[1] for i in cod])
    for i in cod:
        i[0] -= x0
        i[1] -= y0


def zoom(cod:list[list]):
    """缩放图片的函数"""
    x1 = max([i[0] for i in cod])
    y1 = max([i[1] for i in cod])
    for i in cod:
        i[0] = int((460/x1)*i[0])
        i[1] = int((658/y1)*i[1])

# 预留数据空列表和数字字典
data = []
name = {0: "one", 1: "two", 2: "three", 3: "four", 4: "five", 5: "six",
        6: "seven", 7: "eight", 8: "nine", 9: "zero"}
for i in name:
    l = []

    for j in range(31):
        try:
            l += get_coordinate(f"./ori_data/{name[i]}/{j}.png")
        except FileNotFoundError:
            break
    cut_coordinate(l)
    zoom(l)
    data.append(l)



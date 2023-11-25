from Neigh import *
from imgfunc import *
from collections import Counter

radis = 2
neigh = Neigh(radis)

# 预留数据和值的空列表
da = []
va = []

# 生成数据和值
for i in range(10):
    d = data[i]
    value = [name[i] for k in range(len(d))]
    da += d
    va += value

neigh.init_board()
neigh.fit(da, va)

def deal_img(path:str):
    t = get_coordinate(path)
    cut_coordinate(t)
    zoom(t)
    l = neigh.predict(t)
    ll = Counter(l)
    d = {"one":1,
         "two":2,
         "three":3,
         "four":4,
         "five":5,
         "six":6,
         "seven":7,
         "eight":8,
         "nine":9,
         "zero":0}
    return d[ll.most_common(1)[0][0]]



# for i in name:
#     t = get_coordinate(f"./try/test_0/{name[i]}.png")
#     cut_coordinate(t)
#     zoom(t)
#     l = neigh.predict(t)
#     ll = Counter(l)
#     print(name[i], "识别为：", ll.most_common(6))

# 这是一个减小了图片识别区域的版本，识别效果要有很大提升
# 在大的选取量意义下，这更像是一种KNN-排除算法，而不是分类-KNN算法
# 但是在这种基础下，准确度大大提高了，一方面是因为数据集的增大，另一方面是算法的改良

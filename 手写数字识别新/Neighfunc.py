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

for i in name:
    t = get_coordinate(f"./try/test_0/{name[i]}.png")
    cut_coordinate(t)
    zoom(t)
    l = neigh.predict(t)
    ll = Counter(l)
    print(name[i], "识别为：", ll)

# 这是一个减小了图片识别区域的版本，识别效果要有很大提升
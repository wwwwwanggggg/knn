from sklearn.neighbors import KNeighborsClassifier
from imgfunc import *
from collections import Counter


# 设置计算半径，创建KNN模型
radis = 25
neigh = KNeighborsClassifier(n_neighbors=radis)

# 预留数据和值的空列表
da = []
va = []

# 生成数据和值
for i in range(10):
    d = data[i]
    value = [name[i] for k in range(len(d))]
    da += d
    va += value

# 训练模型
neigh.fit(da, va)

# 获取数据测试模型
for i in name:
    t = get_coordinate(f"./try/test_1/test_{name[i]}_0.png")
    cut_coordinate(t)
    zoom(t)
    l = neigh.predict(t)
    l = list(l)
    ll = Counter(l)
    print(name[i], "识别为：", ll.most_common(1))

from Neigh import *
from imgfunc import *
from collections import Counter

radis = 1
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
    t = get_coordinate(f"./try/test_1/test_{name[i]}_0.png")
    cut_coordinate(t)
    zoom(t)
    l = neigh.predict(t)

    ll = Counter(l)
    print(name[i], "识别为：", ll.most_common(4))

# 得出结论：测试细的字体要大数拟合，测试粗的字体要小数拟合(这是在加入粗字体之前得出的结论)
# 但是三个测试集在radis=1以及严格KNN的条件下表现的要最好，实在是令人费解！！！
# 看来事情没有那么简单，重复运行的结果不一样，且这几次的结果错误最多达到了7，实在是不靠谱
# 也还好，我只是像在img-func.py中提前分配好后面训练集的位置，结果仅仅是改了range的值就导致了预测的值不相同，非常抽象
# 加入粗体字的训练集之后，效果甚至有所下降，非常抽象哈

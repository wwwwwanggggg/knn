from collections import Counter


def distance(x1, y1, x2, y2, r):
    if (x1 - x2) ** 2 + (y1 - y2) ** 2 <= r ** 2:
        return True
    else:
        return False


class Point:
    """描述单个点的类"""

    def __init__(self):
        """初始化点的属性"""
        self.value = []


class Neigh:
    """邻近类"""

    def __init__(self, radis):
        """初始化模型的属性"""
        self.x = 300
        self.y = 300
        self.board = []
        self.r = radis

    def init_board(self):
        """生成一个数据板"""
        for i in range(self.x):
            lin = []
            for j in range(self.y):
                lin.append(Point())
            self.board.append(lin)

    def fit(self, data: list[list[int, int]], value: list[list]):
        """训练方法"""
        for i in range(len(data)):
            self.board[data[i][0]][data[i][1]].value.append(value[i])
            self.board[data[i][0]][data[i][1]].value = list(set(self.board[data[i][0]][data[i][1]].value))

    def predict(self, data: list[list[int, int]]):
        """预测方法"""
        res = []
        for i in range(len(data)):
            tem = []
            x_min = data[i][0] - self.r if data[i][0] - self.r > 0 else 0
            y_min = data[i][1] - self.r if data[i][1] - self.r > 0 else 0
            x_max = data[i][0] + self.r if data[i][0] + self.r < self.x else self.x
            y_max = data[i][1] + self.r if data[i][1] + self.r < self.y else self.y
            for j in range(x_min, x_max):
                for k in range(y_min, y_max):
                    tem.extend(self.board[j][k].value)
            count = Counter(tem)
            va = count.most_common(2)
            for l in va:
                res.append(l[0])

        return res

#
# x = Neigh(2)
# x.init_board()
# x.fit([[1,2],[3,4]],[["one"],["one","two"]])
#
# z = x.predict([[1,2],[3,4]])
# print(z)

from tkinter import *
from tkinter.filedialog import *

root = Tk()
root.geometry("400x300+400+150")

img = PhotoImage(file="D:/coding/各种项目/手写数字识别/ori_data/six/0.png",width=10,height=10)
fl = {"path":""}

Label(root,image=img).pack(side="left")
def btn_get_file():
    t = askopenfilename()
    fl["path"] = t


btn = Button(text="导入图片", width=10, font=("华文行楷", 10), bg="lightblue", command=btn_get_file)
btn.place(x=30, y=30)

if __name__ == "__main__":
    mainloop()

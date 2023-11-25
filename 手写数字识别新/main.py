from tkinter import *
from tkinter.filedialog import *
from PIL import Image, ImageTk
from Neighfunc import deal_img

root = Tk()
root.geometry("400x400+400+150")
Label(text="手写数字识别", font=("楷体", 20)).place(x=120, y=10)


def btn_get_file():
    t = askopenfilename()
    img = Image.open(t)
    img = img.resize((200, 200))
    img = ImageTk.PhotoImage(img)
    l = Label(root, image=img)
    l.image = img
    l.place(x=10, y=70)
    Label(text="识别为：", font=("华文新魏", 15)).place(x=10, y=300)
    res = deal_img(t)
    Label(text=res, font=("华文新魏", 15)).place(x=70, y=300)


btn = Button(text="导入图片", width=10, font=("华文行楷", 10), bg="lightblue", command=btn_get_file)
btn.place(x=10, y=45)

if __name__ == "__main__":
    mainloop()

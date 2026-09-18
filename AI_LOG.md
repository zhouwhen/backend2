##核心提问记录
1.bcrypt的使用
https://www.doubao.com/thread/xUhCfu6kqRIwUyVg2

2.全局变量和局部变量的选择
https://www.doubao.com/thread/xUSqghsllvBNejsX1

3.token的生成和验证方法
https://www.doubao.com/thread/x0ozOmZBMhnZJjMKK

4.关于项目运行的问题
https://www.doubao.com/thread/x2NQdHl1hXY2j3wNw

(其它都是一些语法问题和帮忙找报错的提问，就不上传了)

##遇到的问题
---第一次写的那个版本，因为只有两个分支，所以我顺手把变量定义在main了，但是后面我想试一下新的东西，就变成了五个分支，数据传输就好麻烦
   -全部把要用的变成全局变量了，就不需要在每个要用的函数都传一次了
---bcrypt完全不知道
   -问豆包搞懂了一点
---关于deleteUser，我第一次加这个功能的时候，为了省事，直接把这个函数命名成delete了，然后就报错，又看不出来为什么，因为我不看不懂它警告我的英文
   -翻译了一下，原来是它把delete（）当成删掉map的delete了，然后我才想起来不能取和功能一样的名字。。
---不理解token，之前只听说过ai里的token
   -懂了

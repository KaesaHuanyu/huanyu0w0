### 布局/Layout
MDL的布局/Layout组件用来作为整个页面其他元素的容器，可以自动适应不同的浏览器、 屏幕尺寸和设备。
![image](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/2/1/img/mdl-layout.png)
---
布局/Layout组件需要按特定的HTML结构进行声明：
```
<any class="mdl-layout mdl-js-layout">
    <any class="mdl-layout__header">...</any>
    <any class="mdl-layout__drawer">...</any>
    <any class="mdl-layout__content">...</any>
</any>
```
需要指出的是，在一个布局声明中，header等子元素不一定全部使用，比如你可以不要 侧栏菜单：
```
<any class="mdl-layout mdl-js-layout">
    <any class="mdl-layout__header">...</any>
    <any class="mdl-layout__content">...</any>
</any>
```
布局组件简化了创建可伸缩页面的过程。确切的说，MDL可以根据屏幕的尺寸设定样式类 的不同显示效果：
- 桌面 - 当屏幕宽度大于840px时，MDL按桌面环境应对
- 平板 - 当屏幕尺寸大于480px，但小于840px时，MDL按平板环境应对。比如，自动隐藏 header、drawer区域等
- 手机 - 当屏幕尺寸小于480px时，MDL按手机环境应对
#### Class Options

MDL class | 说明
---|---
mdl-layout | 声明元素为布局组件
mdl-js-layout | 为布局实现基本的行为逻辑
mdl-layout__header | 声明元素为布局头/header元素
mdl-layout__drawer | 声明元素为侧栏菜单/drawer元素
mdl-layout__content | 声明元素为布局内容/content元素
mdl-layout--fixed-drawer | 将侧栏菜单/drawer声明为固定式
mdl-layout--fixed-header | 将头部/header声明为固定式
mdl-layout--large-screen-only | 在小尺寸屏幕上隐藏头部/header
mdl-layout--overlay-drawer-button | 为布局添加激活侧栏菜单按钮

---
### 头部/Header
布局组件的header子元素由一系列header-row组成：
![image](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/2/2/img/mdl-layout__header.png)
#### Class Options

MDL class | 说明
--- | ---
mdl-layout__header-row | 声明元素为行容器
mdl-layout-title | 声明元素为标题
mdl-layout-icon | 声明元素为菜单图标
mdl-layout-spacer | 声明元素自动填充行容器剩余空间
mdl-layout__header--transparent | 声明布局头为透明背景
mdl-layout__header--scroll | 声明布局头为可滚动
mdl-layout__header--waterfall | 对多行标题，当滚动内容时，仅显示第一行

---
### 导航/Navigation
在header子元素内可以使用导航/navigation，导航块由一个导航容器 和若干导航链接构成：
```
<div class="mdl-layout__header-row">
    <!--导航容器-->
    <nav class="mdl-navigation">
        <!--导航链接-->
        <a href="...">link</a>
        <a href="...">link</a>
        <a href="...">link</a>
    </nav>
</div>
```
如上例所示，导航块使用nav元素建立。在头部的导航块自动按水平排列各 链接项。

一个常见的UI模式是标题居左，导航居右，如下图所示：
![image](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/2/3/img/mdl-navigation.png)
mdl-layout-spacer可以自动地填充行容器（mdl-layout__header-row） 的剩余空间（扣除title和navigation的宽度），因此可以简单地实现为：
```
<div class="mdl-layout__header-row">
    <span class="mdl-layout-title">title</span>
    <div class="mdl-layout-spacer"></div>
    <nav class="mdl-navigation">...</nav>
</div>
```
#### Class Options

MDL class | 说明
--- | ---
mdl-navigation | 声明元素为MDL导航组
mdl-navigation__link | 声明锚点元素为MDL导航链接

---
### 选项卡/Tabs
在布局的头部可以嵌入选项栏/tab-bar，内容区域可以嵌入选项面板/tab-panel。当用户点击 选项栏中的链接/tab*时，自动显示对应的选项面板：
![image](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/2/4/img/mdl-layout__tab.png)
---
在布局头部声明选项栏，需要遵循特定的HTML结构：
```
<header class="mdl-layout__header">
    <!--声明选项栏-->
    <div class="mdl-layout__tab-bar">
        <!--声明选项，通过href绑定对应的面板，对要激活的选项声明is-active-->
        <a href="#panel-1" class="mdl-layout__tab is-active">tab-1</a>
        <a href="#panel-2" class="mdl-layout__tab">tab-2</a>
        <a href="#panel-3" class="mdl-layout__tab">tab-3</a>
    </div>
</header>
```
在布局的内容区域声明选项面板，也依赖于特定的HTML结构：
```
<main class="mdl-layout__content">
    <!--声明选项面板，使用id属性指定锚点，对要初始显示的面板声明is-active-->
    <div class="mdl-layout__tab-panel is-active" id="panel-1">...</div>
    <div class="mdl-layout__tab-panel" id="panel-2">...</div>
    <div class="mdl-layout__tab-panel" id="panel-3">...</div>
</main>
```
#### Class Options

MDL class | 说明
--- | ---
mdl-layout__tab-bar | 声明元素为选项栏
mdl-layout__tab | 声明锚点元素为选项链接
mdl-layout__tab-panel | 声明元素为选项面板
is-active | 将选项链接/tab或选项面板/tab-panel声明为激活
mdl-layout--fixed-tabs | 将头部tab条声明为固定式

---
### 侧拉菜单/Drawer
侧拉菜单默认情况下是隐藏的，需要用户点击按钮：
![image](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/2/5/img/mdl-layout__drawer.png)
可以设置修饰样式类mdl-layout--fixed-drawer来强制显示侧拉菜单（在小尺寸 屏幕下，侧拉菜单总是隐藏的）:
```
<div class="mdl-layout mdl-layout--fixed-drawer">
    <div class="mdl-layout__drawer">...</div>
</div>
```
在侧拉菜单中也可以使用导航，这时所有的链接自动按垂直方向排列：
```
<div class="mdl-layout__drawer">
    <nav class="mdl-navigation">
        <a href="..." class="mdl-navigation__link">link 1</a>
        <a href="..." class="mdl-navigation__link">link 2</a>
    </nav>
</div>
```

#### Class Options

MDL class | 说明
--- | ---
mdl-layout__drawer | 声明元素为侧栏菜单/drawer元素
mdl-layout-title | 声明元素为标题
mdl-navigation | 声明元素为MDL导航组
mdl-navigation__link | 声明锚点元素为MDL导航链接
mdl-layout--fixed-drawer | 将侧栏菜单/drawer声明为固定式

---
#### CONTINUE...
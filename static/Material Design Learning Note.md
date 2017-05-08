## 布局组件
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
## 容器组件
### 单行页脚/Mini footer
MDL的单行页脚/Mini footer组件以单行水平方式组织所有的信息：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/1/img/mdl-mini-footer.png)
单行页脚同样采用flexbox布局，将整行分割为左右两种区域，并 以空格填充剩余的行空间：
```
<footer class="mdl-mini-footer">
    <div class="mdl-mini-footer--left-section">...</div>
    <div class="mdl-mini-footer--right-section">...</div>
</footer>
```
left-section总是向左边对齐，而right-section总是向右边对齐。 单行页脚内可以放置多个left-section或right-section。

在每个区域内，MDL预定义了两种交互元素：链接和社交按钮。

链接/link-list样式应用在列表元素ul上，自动将列表成员水平排列：
```
<div class="mdl-mini-footer--left-section">
    <ul class="mdl-mini-footer--link-list">
        <li><a href="...">link 1</a></li>
        <li><a href="...">link 2</a></li>
        <li><a href="...">link 3</a></li>
    </ul>
</div>
```
社交按钮/social-btn样式将元素修饰为36px正方大小的容器，可以设置其 背景图片来构造图标式按钮。
#### Class Options

MDL class | 说明
--- | ---
mdl-mini-footer | 声明元素为单行页脚组件
mdl-mini-footer--left-section | 声明元素为左区域容器
mdl-mini-footer--right-section | 声明元素为右区域容器
mdl-logo | 声明元素为logo区
mdl-mini-footer--link-list | 声明元素为链接容器
mdl-mini-footer--social-btn | 声明元素为36px大小的方块区域

---
### 多行页脚/Mega footer
MDL的多行页脚/Mega footer组件可以包含多个垂直排列的区域。当我们需要一个复杂 的页脚区域来呈现信息及提供交互手段时，可以使用这个组件：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/2/img/mdl-mega-footer.png)
从上图容易看出，单行页脚/Mini footer组件相当于仅适用多行页脚/Mega footer 组件的bottom-section区域。

当声明为mdl-mega-footer--link-list样式的列表元素出现在drop-down-section 区域时，其列表项是垂直排列的。
#### Class Options

MDL class | 说明
--- | ---
mdl-mega-footer | 声明元素为多行页脚组件
mdl-mega-footer--top-section | 声明元素为顶部区域
mdl-mega-footer--middle-section | 声明元素为中部区域
mdl-mega-footer--bottom-section | 声明元素为底部区域
mdl-mega-footer--left-section | 声明元素在容器内居左
mdl-mega-footer--right-section | 声明元素在容器内居右
mdl-mega-footer--drop-down-section | 声明元素为垂直内容区
mdl-mega-footer--link-list | 声明元素为链接容器
mdl-mega-footer--heading | 声明元素为标题
mdl-logo | 声明元素为logo区
mdl-mega-footer--social-btn | 声明元素为36px正方大小

---
### 栅格/Grid
MDL的栅格/Grid组件是响应式的，可以适应不同屏幕分辨率的布局要求：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/3/img/mdl-grid.png)
栅格/Grid组件根据屏幕尺寸大小，自动地分割行宽：

- 桌面（ > 840px） - 12个单元格
- 平板（ 480px ~ 840px）- 8个单元格
- 手机（ < 480px）- 4个单元格
可以使用mdl-cell--N-col样式声明单元格的宽度：
```
<div class="mdl-grid">
    <div class="mdl-cell mdl-cell--4-col">...</div>
    <div class="mdl-cell mdl-cell--4-col">...</div>
    <div class="mdl-cell mdl-cell--4-col">...</div>
</div>
```
如果我们希望在任何情况下，示例栅格总是显示为相同的列数，那么 可以声明单元格在不同环境下的样式：
```
<div class="mdl-grid">
    <div class="mdl-cell mdl-cell--6-col-desktop mdl-cell--4-col-tablet mdl-cell--2-col-phone">...</div>
    <div class="mdl-cell mdl-cell--6-col-desktop mdl-cell--4-col-tablet mdl-cell--2-col-phone">...</div>
</div>
```
在同一行的各单元格，默认情况下总是拉伸/stretch其高度（采用同一行中最高单元格的高度），可以使用 mdl-cell--bottom样式使单元格不拉伸，并将底部对齐：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/3/img/mdl-cell--bottom.png)
与之类似，mdl-cell--top使单元格顶部对齐，mdl-cell--middle使单元格居中对齐：
```
<div class="mdl-grid">
    <!--顶部对齐-->
    <div class="mdl-cell mdl-cell--top">...</div>
    <!--居中对齐-->
    <div class="mdl-cell mdl-cell--middle">...</div>
    <!--底部对齐-->
    <div class="mdl-cell mdl-cell--bottom">...</div>
</div>
```

#### Class Options
MDL class | 说明
--- | ---
mdl-grid | 将元素声明为grid组件
mdl-cell | 将元素声明为grid组件的单元格cell
mdl-cell--N-col | 设置单元格宽为N（1-12），默认为4。可选
mdl-cell--N-col-desktop | 在桌面环境下设置单元格宽为N（1-12）。可选
mdl-cell--N-col-tablet | 在平板环境下设置单元格宽为N（1-8）。可选
mdl-cell--N-col-phone | 在手机环境下设置单元格宽为N（1-4）。可选
mdl-cell--hide-desktop | 在桌面环境下隐藏该单元格 。可选
mdl-cell--hide-tablet | 在平板环境下隐藏该单元格。可选
mdl-cell--hide-phone | 在手机环境下隐藏该单元格。可选
mdl-cell--stretch | 在垂直方向拉伸单元格以充满父元素。 这是单元格的默认值
mdl-cell--top | 在垂直方向单元格顶部对齐。可选
mdl-cell--middle | 在垂直方向单元格居中对齐。可选
mdl-cell--bottom | 在垂直方向单元格底部对齐。可选

---
### 选项卡/Tabs
MDL的选项卡/Tabs组件用来在多个内容间进行切换：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/4/img/mdl-tabs.png)
选项卡/Tabs组件具有固定的HTML结构，由选项栏、选项面板等元素构成：
```
<!--1. 声明组件-->
<div class="mdl-tabs mdl-js-tabs">
    <!--2. 声明选项栏-->
    <div class="mdl-tabs__tab-bar">
        <!--2.1 声明选项，使用href属性指向选项面板，为要激活的选项应用is-active样式-->
        <a class="mdl-tabs__tab is-active" href="#panel-1">tab-1</a>
        <a class="mdl-tabs__tab" href="#panel-2">tab-2</a>
    </div>
    <!--3. 声明选项面板，使用id属性声明锚点 , 为要显示的面板应用is-active样式-->
    <div class="mdl-tabs__panel is-active" id="panel-1">...</div>
    <div class="mdl-tabs__panel" id="panel-2">...</div>
</div>
```
可以为组件元素应用mdl-js-ripple-effect样式，使点击时具有水纹动效。
#### Class Options
MDL class | 功能
--- | ---
mdl-tabs | 将元素声明为tabs组件
mdl-js-tabs | 实现tabs组件的基本逻辑
mdl-tabs__tab-bar | 将元素声明为tabs组件的导航条容器
mdl-tabs__tab | 将链接元素声明为tabs组件的tab页触发器
mdl-tabs__panel | 将元素声明为tabs组件的tab页内容面板
is-active | 将tab页内容面板或tab页触发器设置为活动状态
mdl-js-ripple-effect | 为tab页触发器增加点击时水纹效果。可选

---
### 卡片/Cards
MDL的卡片/Card组件非常适合显示复杂的、包含多种类型信息的内容：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/3/5/img/mdl-card.png)
卡片通常具有固定的宽度，而高度则根据场景不同，可以固定，也可以变化。使用mdl-card样式类将外层元素声明为卡片组件，使用mdl-card__title 等样式类将内层元素声明为标题、媒体、动作等容器：
```
<any class="mdl-card">
    <any class="mdl-card__title">...</any>
    <any class="mdl-card__media">...</any>
    <any class="mdl-card__supporting-text">...</any>
    <any class="mdl-card__actions">...</any>
    <any class="mdl-card__menu">...</any>
</any>
```
卡片组件默认为330px宽，最小200px高，是一个主轴为竖向的flex容器。可以显式地 设置其宽度和高度。

title、media、supporting-text和actions作为flex容器成员在垂直方向 上依次排列，其高度是由内容决定，或者被显式地设定。例如，很多时候，我们希望给 title区域增加背景图片以增强感染力，那么将照片设置为title区域的背景之后，还需要 设置title区域的高度：
```
<div class="mdl-card">
    <div class="mdl-card__title"
        style="background:url(img/bg.jpg) no-repeat;backgroud-size:cover;height:150px;">
        ...
    </div>
</div>
```
menu块被设置为绝对定位，总是居于卡片组件的右上角。

#### Class Options
MDL class | 说明
--- | ---
mdl-card | 应用在外层容器，声明元素为卡片
mdl-card--border | 为区域增加顶部边框，应用于actions区域和title区域
mdl-shadow--Ndp | 为卡片添加N(2~8)dp的阴影，应用在外层容器
mdl-card__title | 声明容器为卡片标题区域，应用在内层容器
mdl-card__title-text | 为卡片标题设置合适的样式，应用在卡片标题区域的h1~h6
mdl-card__subtitle-text | 为卡片子标题设置合适的样式
mdl-card__media | 声明容器为卡片媒体区域，应用在内层容器
mdl-card__supporting-text | 声明容器为卡片正文区域，应用在内层容器
mdl-card__actions | 声明容器为卡片正文区域，应用在内层容器
mdl-card__menu | 声明元素为卡片菜单按钮区
mdl-card--expand | 声明区域的flex-grow为1，使区域自动增长以填充卡片剩余空间

---
## 交互组件
### 徽章/Badge
徽章/Badge向用户提供了发现额外信息的视觉线索，它通常是圆型，内容为数字 或其他字母，紧贴在宿主元素旁边:
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/1/img/mdl-badge.png)
徽章可以用来无侵入地吸引用户的注意力，例如：

一个新消息通知可以使用徽章提醒有几条未读信息
一个购物车未付款提醒可以使用徽章提醒购物车内的商品数量
一个加入讨论！按钮可以使用徽章提示当前已经加入讨论的用户数
使用MDL徽章组件很简单，为宿主元素添加mdl-badge样式类，然后在data-badge中设置 徽章内容：
```
<any class="mdl-badge" data-badge="1">...</any>
```
因为徽章组件的尺寸很小，所以不要放太多内容，通常data-badge的值设置为1~3个 字符。
#### Class Options
MDL class | 说明
--- | ---
mdl-badge | 声明当前元素为MDL徽章组件
mdl-badge--no-background | 声明徽章组件不使用背景色
data-badge | 徽章组件使用宿主元素上这个属性值来设置显示内容

---
### 提示框/Tooltip
当鼠标移动到元素上方时，提示框/Tooltip组件可以为界面元素提供额外的信息：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/2/img/mdl-tooltip.png)
在MDL中，为一个元素添加Tooltip的步骤如下：
```
<!--1. 为宿主元素定义一个id -->
<button id="test">TEST</button>
<!--2. 声明一个tooltip组件，使用*for*属性绑定到宿主元素上-->
<div class="mdl-tooltip" for="test">这个按钮没什么用;-(</div>
```
尽管在提示框内可以使用HTML片段，但是Material Design设计语言不建议在 提示框中加入图片等复杂的元素。

#### Class Options
MDL class | 说明
--- | ---
mdl-tooltip | 声明元素为MDL提示框组件
mdl-tooltip--large | 为MDL提示框组件应用大字体(14px)

---
### 按钮/Button
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/3/img/mdl-button.png)
MDL按钮的显示类型包括：flat, raised, fab, mini-fab, icon. 这些类型 都可以设置为浅灰或彩色，也可以禁止。fab, mini-fab和icon类型的按钮通常 使用一个小图像而不是文字来表征其功能。

使用按钮组件很简单，为button元素声明mdl-button、mdl-js-button及 其他可选的修饰样式类即可：
```
<!--缺省的扁平/flat按钮-->
<button class="mdl-button mdl-js-button">Save</button>
<!--凸起/raised按钮-->
<button class="mdl-button mdl-js-button mdl-button--raised">Save</button>
<!--浮动动作/FAB按钮-->
<button class="mdl-button mdl-js-button mdl-button--fab">Save</button>
<!--迷你浮动动作/MINI-FAB按钮-->
<button class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab">Save</button>
<!--彩色凸起/raised按钮-->
<button class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored">Save</button>
<!--具有点击动效的凸起/raised按钮-->
<button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect">Save</button>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-button | 将元素声明为MDL按钮组件
mdl-js-button | 为按钮添加基本的行为逻辑
mdl-button--raised | 为按钮应用凸起效果
mdl-button--fab | 将按钮设置为圆形，直径56px
mdl-button--mini-fab | 将fab按钮设置为原型，直径40px。
mdl-button--icon | 为按钮应用图标效果，直径32px
mdl-button--colored | 为按钮应用色彩，使用强调色
mdl-button--primary | 为按钮应用基准色
mdl-button--accent | 为按钮应用强调色
mdl-js-ripple-effect | 为点击动作应用水纹效果

---
### 菜单/Menus
菜单/menu组件提供一组选项供用户选择，用户的选择将执行一个动作、变化设置或 其他可以观察到的效果。当需要用户选择时，显示菜单，当用户完成选择时，关闭菜单：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/4/img/mdl-menu.png)
菜单是成熟然而未标准化的界面组件。

使用mdl-menu样式类声明菜单，使用mdl-menu__item样式类声明菜单项：
```
<any class="mdl-menu mdl-js-menu">
    <any class="mdl-menu__item">...</any>
    <any class="mdl-menu__item">...</any>
</any>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-button | 声明元素为按钮组件
mdl-js-button | 为按钮组件添加基本的逻辑
mdl-button--icon | 使按钮适配图标显示
material-icons | 声明元素为图标
mdl-menu | 声明元素为菜单组件
mdl-menu__item | 声明元素为菜单项
mdl-js-ripple-effect | 为点击动作添加水纹效果
mdl-menu--top-left | 在按钮之上显示菜单，菜单左边框与按钮对齐
mdl-menu--top-right | 在按钮之上显示菜单，菜单右边框与按钮对齐
mdl-menu--bottom-left | 在按钮之下显示菜单，菜单左边框与按钮对齐
mdl-menu--bottom-right | 在按钮之下显示菜单，菜单右边框与按钮对齐

---
### 滑动条/Sliders
MDL的滑动条/slider组件是HTML5新增元素range input的增强版本。 滑动条由一条水平线及其上的可移动滑块构成。当用户移动滑块时，就可以 从预设范围中选择一个值（左边是下界，右边是上界）：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/5/img/mdl-slider.png)
使用MDL的滑动条组件很简单，为range input元素应用样式类mdl-slider 和mdl-js-slider即可：
```
<input type="range" min="0" max="100" class="mdl-slider mdl-js-slider" >
```
使用range input元素的min和max属性来设定值的范围，使用value属性来 设置滑动条的初始值：
```
<input type="range" min="0" max="100" value="25" class="mdl-slider mdl-js-slider" >
```
#### Class Options
MDL class | 说明
--- | ---
mdl-slider | 声明元素为滑动条组件
mdl-js-slider | 为滑动条添加基本的行为逻辑

---
### 复选按钮/Checkbox
MDL的复选按钮/Checkbox组件是标准HTML元素checkbox input的增强版本。 复选按钮组件包含一个标签和一个开关选择按钮：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/6/img/mdl-checkbox.png)
MDL的复选按钮/Checkbox组件具有预定义的HTML结构：
```
<!--1. 声明组件容器-->
<label class="mdl-checkbox">
    <!--2. 为checkbox input元素应用mdl样式类-->
    <input type="checkbox" class="mdl-checkbox__input"/>
    <!--3. 为标签元素应用mdl样式类-->
    <span class="mdl-checkbox__label">标签</span>
</label>
```
可以使用checkbox input元素的checked属性设置复选按钮组件的初始选中状态。

#### Class Options
MDL class | 说明
--- | ---
mdl-checkbox | 声明元素为复选按钮
mdl-js-checkbox | 为复选按钮添加基本的行为逻辑
mdl-checkbox__input | 为组件的input子元素使用此样式
mdl-checkbox__label | 为组件的label子元素使用此样式
mdl-js-ripple-effect | 为点击动作添加水纹效果

---
### 单选按钮/Radio button
MDL的单选按钮/RadioButton组件是标准HTML元素radio input的增强版本。 单选按钮组件包含一个标签和一个开关选择按钮：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/7/img/mdl-radio.png)
MDL的单选按钮组件具有固定的HTML结构：
```
<!--1. 声明组件容器-->
<label class="mdl-radio mdl-js-radio">
    <!--2.为input子元素应用mdl样式类-->
    <input type="radio" class="mdl-radio__button" name="options" value="1"/>
    <!--3.为label子元素应用mdl样式类-->
    <span class="mdl-radio__label">选项1</span>
</label>
<!--选项2-->
<label class="mdl-radio mdl-js-radio">
    <input type="radio" class="mdl-radio__button" name="options" value="2"/>
    <span class="mdl-radio__label">选项2</span>
</label>
```
和复选按钮不同，多个同时出现的单选按钮组件，其选中状态是互斥的，任何时刻最多只有一个 可以被选中。

和复选按钮类似，使用radio input元素的checked属性设置单选按钮的选中状态。

#### Class Options
MDL class | 说明
--- | ---
mdl-radio | 声明元素为单选按钮
mdl-js-radio | 为单选按钮添加基本的行为逻辑
mdl-radio__button | 为input元素声明此样式
mdl-radio__label | 为label元素声明此样式
mdl-js-ripple-effect | 为点击动作应用水纹效果

---
### 图标开关/Icon toggle
MDL的图标开关/IconToggle组件是标准HTML元素checkbox input的增强版本。 图标开关组件包含一个标签和一个用户指定的图标按钮，图标的着色与否用来传达 当前选项是否被选中：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/8/img/mdl-icon-toggle.png)
MDL的图标开关/IconToggle组件具有预定义的HTML结构：
```
<!--1. 声明组件容器-->
<label class="mdl-icon-toggle mdl-js-icon-toggle">
    <!--2. 为checkbox input元素应用mdl样式类-->
    <input type="checkbox" class="mdl-icon-toggle__input"/>
    <!--3. 为图标元素应用mdl样式类-->
    <i class="mdl-icon-toggle__label material-icons">format_bold</i>
</label>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-icon-toggle | 声明元素为图标开关
mdl-js-icon-toggle | 为图标开关添加基本的行为逻辑
mdl-icon-toggle__input | 为input元素声明此样式
mdl-icon-toggle__label | 为label元素声明此样式
mdl-js-ripple-effect | 为点击动作添加水纹效果

---
### 进度条/Progress bar
MDL的进度条/progress bar组件用来提供后台活动的可视化反馈。进度条是一个水平的 长条，可以包含动画以传递工作中的感觉：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/9/img/mdl-progress.png)
使用MDL进度条/Progress bar组件很简单：
```
<any class="mdl-progress mdl-js-progress "></any>
```
如果不需要提供给用户进度完成的具体百分比，可以附加一个动画：
```
<any class="mdl-progress mdl-js-progress mdl-progress__indeterminate"></any>
```
如果需要显示进度百分比，需要使用挂接在DOM对象上的MaterialProgress 变量的setProgress()方法：
```
var el = document.querySelector("#p1");
//setProgress()方法接受一个0~100的值
el.MaterialProgress.setProgress(80);
```
如果需要同时显示一个视频流的缓冲及播放情况，可以使用MaterialProgress变量的 setBuffer()方法，这个方法将对未缓冲的部分播放一个动画来表达缓冲效果：
```
var el = document.querySelector("#p1");
//setBuffer()方法接受一个0~100的值
el.MaterialProgress.setBuffer(80);
```
#### Config Options
MDL class | 说明
--- | ---
mdl-progress | 声明元素为进度条组件
mdl-js-progress | 为进度条组件添加基本的行为逻辑
mdl-progress__indeterminate | 为元素应用动画效果。可选

---
### 等待指示器/Spinner
MDL的等待指示器/spinner组件是等待图标的增强版本，它使用一个边框色彩动态变化 的圆框，清晰地向用户传达作业已经开始、还未完成的状况：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/10/img/mdl-spinner.png)
使用spinner组件非常简单：
```
<any class="mdl-spinner mdl-js-spinner"></any>
```
spinner默认是隐藏的，为其应用is-active样式进行激活：
```
<any class="mdl-spinner mdl-js-spinner is-active"></any>
```
#### Class Options

MDL class | 说明
--- | ---
mdl-spinner | 声明元素为spinner组件
mdl-js-spinner | 为spinner增加基本的行为逻辑
is-active | 显示spinner组件并激活动画
mdl-spinner--single-color | 只使用单一色彩

---
### 文本输入Text Field
MDL的文本输入/Text Field组件是对标准HTMLtext input元素的封装：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/11/img/mdl-textfield.png)
文本输入组件有特定的HTML结构：
```
<!--1.声明组件-->
<div class="mdl-textfield mdl-js-textfield">
    <!--2.声明组件的input元素-->
    <input type="text" class="mdl-textfield__input"/>
    <!--3.声明组件的label元素-->
    <label class="mdl-textfield__label">Your Name</label>
    <!--4.声明组件的error元素-->
    <span class="mdl-textfield__error">Error!</span>
</div>
```
error元素默认是隐藏的，用来向用户反馈输入的错误。可以为input元素设置 pattern属性（这是HTML5的新特性），当用户的输入与pattern指定的正则 表达式不符时，将显示error元素：
```
<input type="text" pattern="-?[0-9]*(\.[0-9]+)?"/>
```
上面的正则表达式将检测用户的输入是否是一个数值，例如：-123.456 。

默认情况下，当用户开始输入时，标签将消失。可以为组件应用mdl-textfield--floating-label 样式开启浮动标签模式：
```
<!--用户输入时，标签将浮动在输入行上方-->
<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">...</div>
```
也可以将input元素换成textarea元素，这样将允许多行输入：
```
<div class="mdl-textfield mdl-js-textfield">
    <!--使用rows属性声明行数-->
    <textarea class="mdl-textfield__input" rows="3"></textarea>
    <label class="mdl-textfield__label">Memo</label>
</div>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-textfield | 声明元素为文本输入组件
mdl-js-textfield | 为文本输入组件添加基本的行为逻辑
mdl-textfield__input | 为input元素应用此样式
mdl-textfield__label | 为label元素应用此样式
mdl-textfield--floating-label | 为文本输入组件应用浮动标签效果
mdl-textfield__error | 声明span元素为MDL错误信息容器

---
### 文本输入 - 动态展开式
一种常见的文本输入模式具有一个按钮，点击这个按钮将展开输入框，如果 没有输入内容，那么当输入框失去焦点时将自动隐藏：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/12/img/mdl-textfield--expandable.png)
动态展开的文本输入组件有特定的HTML结构：
```
<!--1.使用expandable样式类声明动态展开的文本输入组件-->
<div class="mdl-textfield mdl-js-textfield mdl-textfield--expandable">
    <!--2. 声明触发按钮，使用属性for绑定到input元素-->
    <button class="mdl-button mdl-js-button" for="kw_inp">search</button>
    <!--3. 声明文本输入框容器-->
    <div class="mdl-textfield__expandable-holder">
        <!--4.声明input元素，使用属性id声明锚点-->
        <input type="text" class="mdl-textfield__input" id="kw_inp"/>
        <!--5.声明label元素-->
        <label class="mdl-textfield__label">keywords</label>
    </div>
</div>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-textfield--expandable | 声明元素为可展开文本输入组件
mdl-input__expandable-holder | 声明元素为文本输入元素的容器

---
### 数据表/Data table
MDL的数据表/Data table组件用来呈现密集的数据集：
![images](http://cw.hubwiz.com/card/c/55adae643ad79a1b05dcbf77/1/4/13/img/mdl-data-table.png)
```
<table class="mdl-data-table mdl-js-data-table">
    <thead>...</thead>
    <tbody>...</tbody>
</table>
```
#### Class Options
MDL class | 说明
--- | ---
mdl-data-table | 声明元素为数据表组件
mdl-js-data-table | 为数据表添加基本的行为逻辑
mdl-data-table--selectable | 为数据表的每一条记录添加复选按钮，应用在table元素上
mdl-data-table__cell--non-numeric | 声明单元格内容非数字，使单元格文字左对齐
---
## END
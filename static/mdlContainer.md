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
#### CONTINUE...
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
#### ...END
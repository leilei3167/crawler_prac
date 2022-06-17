package main

//https://cloud.tencent.com/developer/article/1196783
func main() {
	/* 元素选择器:
	   基于html基本元素名称进行选择,如Find("div") Find("p")等
	*/

	/*
		ID选择器:
		常用,假设某个标签有ID,如果需要快速定位 则可使用id选择器,以#为开头,后跟id名,如

		html := `<body>

						<div id="div1">DIV1</div>
						<div class="name">DIV2</div>
						<span>SPAN</span>

					</body>
					`

			使用id选择器:Find("#div1")
	*/

	/*
		class选择器:
		我们可以通过class选择器来快速的筛选需要的HTML元素，它的用法和ID选择器类似，
		为Find(".name"),筛选出class为name的这个div元素

	*/
	/*
		元素,class选择器:
		以上两者可以结合使用:Find("元素名.class名")

	*/
	/*
		属性选择器:
		有些元素都有自己的属性及属性值
		Find("div[rel]") 表示筛选有class的div元素
		也可以精确指定:Find("div[class=name]"),筛选出属性为某个值的元素
		属性值除了有完全相等的匹配方式外,还有多种:

		*Find(“div[lang]“):筛选含有lang属性的div元素
		*Find(“div[lang=zh]“):筛选lang属性为zh的div元素
		Find(“div[lang!=zh]“):筛选lang属性不等于zh的div元素
		Find(“div[lang¦=zh]“):筛选lang属性为zh或者zh-开头的div元素
		*Find(“div[lang*=zh]“):筛选lang属性包含zh这个字符串的div元素
		Find(“div[lang~=zh]“):筛选lang属性包含zh这个单词的div元素，单词以空格分开的
		*Find(“div[lang$=zh]“):筛选lang属性以zh结尾的div元素，区分大小写
		*Find(“div[lang^=zh]“):筛选lang属性以zh开头的div元素，区分大小写

		也可以使用多个属性筛选器组合:
		Find("div[id][lang=zh]"),使用多个中括号连接即可,代表具有id属性,并且lang属性值为zh的div元素





	*/

	/*
		parent>child 父子选择器:
		如果我们想筛选出某个元素下符合条件的子元素，我们就可以使用子元素筛选器，它的语法为Find("parent>child"),
		表示筛选parent这个父元素下，符合child这个条件的最直接（一级）的子元素。
		带>的语句只会找到下一级的元素,再往下不会返回,如果要返回某个元素所有的子元素,将>替换为空格即可:
		Find("parent child")

	*/
	/*
		prev+next 相邻选择器:
		如果我们要筛选的元素没有规律,但是他上一个元素有规律,则可以使用此选择器,加号前后都是选择器

				   html := `<body>

							<div lang="zh">DIV1</div>
							<p>P1</p>
							<div lang="zh-cn">DIV2</div>
							<div lang="en">DIV3</div>
							<span>
								<div>DIV4</div>
							</span>
							<p>P2</p>

						</body>
						`
		Find("div[lang=zh]+p"):我们想选择<p>P1</p>这个元素，但是没啥规律,但是不想选到P2

	*/

	/*
		prev~next 兄弟选择器:
		与相邻选择器类似,只是说这个是将所有同级的兄弟都选进来

	*/

	/*
		内容过滤器(contains,empty,has):
		有时候选择器选择出来后,我们希望再进行过滤,就需要用到过滤器,过滤器有很多种,通常标志就是:
		如:
		Find("div:contains(DIV2)") 表示筛选出来的元素要包含指定的文本,这里表示选出的div要包含
		DIV2这个字符串
		Find("div:empty") 表示筛选出的div元素不能有子元素(包括文本)
		还有一个has的用法:
		Find("div:has(选择器)") 表示选出的元素必须要有某些元素
	*/

	/*
		first-child过滤器:
		过滤出作为父元素第一个的子元素的元素
		html := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div style="display:none;">DIV4</div>
					<div>DIV5</div>
				</span>
				<p>P2</p>
				<div></div>

			</body>
			`
		Find("div:first-child"):会筛选出所有的div元素，但是我们加了:first-child后，就只有DIV1和DIV4了，
		因为只有这两个是他们父元素的第一个子元素，其他的DIV都不满足。

	*/

	/*
		:first-of-type过滤器:
		first-child限制必须是第一个子元素,如果该元素前面有其他元素,则无法选出;而first-of-type
		则只要求是这一类的第一个就可以

		相反的,还有last-of-type和last-child
	*/

	/*
		:nth-child(n)过滤器:
		这个表示筛选出的元素是其父元素的第n个元素，n以1开始。
		所以我们可以知道:first-child和:nth-child(1)是相等的

		:nth-of-type(n) 过滤器:
		它表示的是同类型元素的第n个,所以:nth-of-type(1) 和 :first-of-type是相等的

		nth-last-child(n) 和:nth-last-of-type(n) 过滤器:
		倒数第n个
	*/
}

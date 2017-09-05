class svgShower{
	//min生成数组最小值
	//max生成数组最大值
	//count数组长度，点的个数
	constructor(min, max, count)
	{
		this.width = 1000;
		this.height = 500;
		this.min = min;
		this.max = max;
		this.range = max - min;
		this.count = count;
		this.point_prefix = 'point_';
		this.point_width = parseInt(this.width / this.count)-4;
		this.point_width_half = parseInt(this.width/(this.count*2));
		this.swap_arr = [];
		this.scale_x = d3.scaleLinear().domain([0, count]).range([0, this.width]);
		this.scale_y = d3.scaleLinear().domain([this.min, this.max]).range([this.height, 0]);
		this.root = d3.select('svg').append('g').attr('id', 'root');
		this.init_xy_axis()
		this.points = this.root.append('g').attr('id', 'points');
	}
	//初始化坐标轴
	init_xy_axis()
	{
		var x_axis = d3.axisBottom().scale(this.scale_x).tickSize(12);
		var y_axis = d3.axisLeft().scale(this.scale_y).ticks(10, 's').tickSize(10);
		var root = this.root;
		root.attr('transform', 'translate(50,50)');
		root.append('g').call(x_axis).attr('transform', 'translate(0,500)');
		root.append("g").call(y_axis);
	}
	//增加数组中的值
	add_point(index, value)
	{

		this.points.append('rect')
			.attr('x', this.scale_x(index)+2)
			.attr('y', this.scale_y(value))
			.attr('width', this.point_width)
			.attr('height', this.scale_y(this.max-value))
			.attr('style', 'stroke:#000;stroke-width:1;fill:green;')
			.attr('id', this.point_prefix + index);
	}
	add_swap(index_1, index_2)
	{
		this.swap_arr.push({key1:index_1, key2:index_2});
	}
	show_swap()
	{
		if(this.swap_arr.length == 0) return;
		var data = this.swap_arr[0];
		var point_1 = d3.select('#'+this.point_prefix+data.key1);
		var point_2 = d3.select('#'+this.point_prefix+data.key2);
		var path = 'M' + point_1.attr('x') + ',' + point_1.attr('height') + ' L' + point_2.attr('x') + ',' + point_2.attr('height');
		point_1.style('fill','#9cb10e');
		point_2.style('fill', '#9cb10e');
		setTimeout(()=>{
			this.swap_arr.shift();
			var height = point_1.attr('height');
			var y = point_1.attr('y');
			point_1.attr('height', point_2.attr('height'));
			point_1.attr('y', point_2.attr('y'))
			point_2.attr('height', height);
			point_2.attr('y', y);
			point_1.style('fill','green');
			point_2.style('fill', 'green');
			this.show_swap();
		}, 500);
	}
	init_points(arr){
		
	}
	create_arr(){
		var arr = [];
		var length = this.count;
		for(var i=0; i<length; i++){
			var num = this.getRandInt();
			arr[i] = num;
			this.add_point(i, num);
		}
		return arr;
	}
	getRandInt(){
		var rand = Math.random();
		return Math.round(this.range * rand);
	}
}

var svgShowerObj = new svgShower(0, 100, 20);
var arr = svgShowerObj.create_arr();

//
var length = arr.length;
for(var i=0; i<length; i++){
	var end = length - i -1;
	for(var j=0 ;j<end; j++){
		if (arr[j] > arr[j+1]){
			var tmp = arr[j];
			arr[j] = arr[j+1];
			arr[j+1] = tmp;
			svgShowerObj.add_swap(j, j+1);
		}
	}
}
console.log(arr);
svgShowerObj.show_swap();

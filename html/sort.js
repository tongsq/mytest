var scale_x = d3.scaleLinear().domain([0,100]).range([0,1000]);
var scale_y = d3.scaleLinear().domain([0,100]).range([500,0]);

var x_axis = d3.axisBottom().scale(scale_x).tickSize(12);
var y_axis = d3.axisLeft().scale(scale_y).ticks(10,"s").tickSize(10);
var axis_g = d3.select("#svg").append('g');
axis_g.attr('transform', 'translate(50,50)');
axis_g.append("g").call(x_axis).attr('transform', 'translate(0,500)');
axis_g.append("g").call(y_axis);
			

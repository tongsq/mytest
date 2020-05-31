import React,{Component} from 'react';
import ReactDOM from 'react-dom';

import './index.css';

class Hello extends Component{
	render(){
		return (
			<div><h1>hello</h1></div>
			);
	}
}

ReactDOM.render(
  <Hello />,
  document.getElementById('root')
);
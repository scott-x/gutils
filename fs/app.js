/*
* @Author: scottxiong
* @Date:   2020-04-15 16:23:06
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-04-15 16:23:08
*/
import React from 'react'
import ReactDOM from 'react-dom'

import App from './app/component/App'

let data = [
 {id: 0, text: 'go swimming!!!', complete: false},
 {id: 1, text: 'go hiking!!!', complete: false},
 {id: 2, text: 'go shopping!!!', complete: true},
]

ReactDOM.render(
  <App data={data}/>,
  document.getElementById('app')
)
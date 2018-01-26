import React from 'react'
import ReactDOM from 'react-dom'
import Root from '@/views/Root'

import axios from 'axios'
axios.defaults.baseURL = process.env.API_HOST

import 'normalize.css'
import 'font-awesome/css/font-awesome.css'
import '@/components/variables.css'
import '@/components/basic.css'

ReactDOM.render(<Root />, document.getElementById('root'));

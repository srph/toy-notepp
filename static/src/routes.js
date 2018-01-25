import App from '@/views/Root'
import {Route, Switch} from 'react-router'
import App from '@/views/app'
import AppMain from '@/views/app.main'
import AppHome from '@/views/app.main.home'
import Login from '@/views/login'

export default (
  <App>
    <Route path="/login" component={Login} />
    
    <AppMain>
      <Route path="/" component={AppHome} />
    </App>
  </App>
)
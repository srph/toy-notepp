import {Route, Switch} from 'react-router'
import App from '@/views/app'
import AppMain from '@/views/app.main'
import AppHome from '@/views/app.main.home'
import AppLogin from '@/views/app.login'

export default (
  <App>
    <Route path="/login" component={Login} />
    
    <AppMain>
      <Route path="/" component={AppHome} />
    </AppMain>
  </App>
)
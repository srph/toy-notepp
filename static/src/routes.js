import React from 'react'
import {Route, Switch} from 'react-router'
import App from '@/views/app'
import AppMain from '@/views/app.main'
import AppHome from '@/views/app.main.home'
import AppLogin from '@/views/app.login'
import AppLogout from '@/views/app.logout'
import Compose from '@/components/Compose'
import Permission from '@/components/Permission'

export default (
  <App>
    <Switch>
      <Route path="/login" component={Permission.guest(AppLogin)} />
      <Route path="/logout" component={AppLogout} />
      
      <Compose hoc={Permission.auth} component={AppMain}>
        <Route path="/" component={AppHome} />
        <Route path="/trash" component={AppHome} />
      </Compose>
    </Switch>
  </App>
)
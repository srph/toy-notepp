import React from 'react'
import store from '@/store'
import * as auth from '@/store/auth'

export default class AppLogout extends React.Component {
  componentDidMount() {
    store.dispatch(auth.actions.logout())
  }
  
  render() {
    return (
      <div />
    )
  }
}
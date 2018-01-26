import React from 'react'
import Helmet from 'react-helmet'

import {Div} from 'glamorous'
import UiInput from '@/components/UiInput'
import UiButton from '@/components/UiButton'
import linkState from 'linkstate'

import * as auth from '@/store/auth'
import {connect} from 'react-redux'

class AppLogin extends React.Component {
  state = {
    username: '',
    password: ''
  }

  render() {
    const {loading, error} = this.props

    return (
      <div>
        <Helmet title="Welcome Back" />

        <Div
          width="320px"
          margin="0 auto"
          paddingTop="40px"
          paddingBottom="40px">
          <form onSubmit={this.handleSubmit}>
            <UiInput
              value={this.state.username}
              onChange={linkState(this, 'username')}
              name="username"
              placeholder="johndoe03"
              marginBottom="8px" />

            <UiInput
              value={this.state.password}
              onChange={linkState(this, 'password')}
              placeholder="*****"
              marginBottom="16px" />

            <Div textAlign="right">
              <UiButton disabled={loading}>Login</UiButton>
            </Div>
          </form>
        </Div>
      </div>
    )
  }

  handleSubmit = evt => {
    evt.preventDefault()
    this.props.dispatch(auth.actions.login(this.state))
  }
}

export default connect(state => {
  return {
    loading: state.auth.loading,
    error: state.auth.error
  }
})(AppLogin)
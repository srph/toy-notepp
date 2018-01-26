import React from 'react'
import Resolver from '@/components/Resolver'
import store from '@/store'
import * as auth from '@/store/auth'

class App extends React.Component {
  render() {
    return (
      <div>
        <Resolver promise={() => store.dispatch(auth.actions.init())}>
          {this.props.children}
        </Resolver>
      </div>
    )
  }
}

export default App
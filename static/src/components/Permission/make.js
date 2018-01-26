import React, {cloneElement} from 'react'
import store from '@/store'

/**
 * Permission high-order component
 * @param {Function} condition
 */
function make(condition) {
  return (Component) => {
    class Permission extends React.Component {
      state = {
        resolved: false
      }

      componentDidMount() {
        if (condition(store.getState()) !== false) {
          this.setState({ resolved: true })
        }
      }

      render() {
        if (this.state.resolved) {
          return <Component {...this.props} />;
        } else {
          return <div />;
        }
      }
    }

    return Permission
  }
}

export default make
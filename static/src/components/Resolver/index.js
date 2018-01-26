import React, {Children, cloneElement} from 'react'

/**
 * Delay component from mounting until promise is resolved.
 */
class Resolver extends React.Component {
  state = {
    resolved: false,
    data: null
  }

  componentDidMount() {
    this.props.promise()
      .then(data => {
        this.setState({
          resolved: true,
          data
        })
      })
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.shouldRun(nextProps)) {
      this.props.promise()
        .then(data => {
          this.setState({
            resolved: true,
            data
          })
        })
      }
  }

  render() {
    if (!this.state.resolved) {
      return <div />
    }

    if (this.props.key) {
      return this.props.children({
        ...this.props,
        [key]: this.state.data
      })
    }

    return Children.map(this.props.children, child => {
      return cloneElement(child, this.getProps())
    })
  }

  getProps(props = this.props) {
    const {key, promise, shouldRun, ...waterfall} = this.props
    return waterfall
  }
}

Resolver.defaultProps = {
  shouldRun: () => true
}

Resolver.resolver = function(key, promise, shouldRun) {
  // resolver(Function)
  if (promise == null) {
    key = null
    promise = key
  }

  return Component => {
    return class extends React.Component {
      render() {
        <Resolver key={key}
          promise={promise}
          shouldRun={shouldRun}
          {...this.props} />
      }
    }
  }
}

export default Resolver
import React, {cloneElement} from 'react'

/**
 * Make it easier to apply HOC to a component
 * without having to assign it to a variable
 */
class Compose extends React.Component {
  applied = this.props.hoc(this.props.component)

  render() {
    const {hoc, component, ...props} = this.props
    const Applied = this.applied
    return <Applied {...props} />
  }
}

export default Compose
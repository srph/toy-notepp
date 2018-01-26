import React from 'react'
import {Input} from 'glamorous'

class UiInput extends React.Component {
  render() {
    return (
      <Input {...this.props}
        display="block"
        height="var(--form-size)"
        paddingLeft="8px"
        paddingRight="8px"
        marginRight="16px"
        width="100%"
        color="var(--color-black)"
        borderRadius="var(--border-radius)"
        border="1px solid var(--color-silver)"
        outline="0" />
    )
  }
}

export default UiInput
import React from 'react'
import {Div, Input, Button} from 'glamorous'

class Searchbar extends React.Component {
  render() {
    return (
      <Div
        flexShrink="0"
        display="flex"
        padding="16"
        borderBottom="1px solid var(--color-silver)">
        <Input display="block"
          height="var(--form-size)"
          paddingLeft="8px"
          paddingRight="8px"
          marginRight="16px"
          width="100%"
          borderRadius="var(--border-radius)"
          border="1px solid var(--color-silver)"
          outline="0"
          placeholder="Search notes..." />

        <Button
          flexShrink="0"
          display="inline-block"
          padding="0"
          background="transparent"
          border="0"
          outline="0"
          color="var(--color-silver)"
          cursor="pointer"
          title="Create new note">
          <i className="fa fa-pencil-square-o" />
        </Button>
      </Div>
    )
  }
}

export default Searchbar
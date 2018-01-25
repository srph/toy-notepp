import React from 'react'
import {Div, A} from 'glamorous'

function UiSidebarLink({link, icon, text}) {
  return (
    <A href="#"
      display="flex"
      alignItems="center"
      padding="8px 16px"
      color="var(--color-white)"
      backgroundColor="transparent"
      textDecoration="none"
      transition="200ms all ease"
      css={{
        ':hover': {
          backgroundColor: 'var(--color-gray)'
        }
      }}>
      <Div marginRight="8px">
        <i className={`fa fa-${icon}`} />
      </Div>
      <span>{text}</span>
    </A>
  )
}

export default UiSidebarLink
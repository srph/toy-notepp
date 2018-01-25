import React from 'react'
import {Div, A} from 'glamorous'

function UiSidebarLink({link, icon, text, active}) {
  return (
    <A href="#"
      display="flex"
      marginBottom="8px"
      alignItems="center"
      padding="8px 16px"
      fontFamily="var(--font-montserrat)"
      fontSize="12px"
      color="var(--color-white)"
      backgroundColor={active ? 'var(--color-blue)' : 'transparent'}
      textDecoration="none"
      borderRadius="var(--border-radius)"
      transition="200ms all ease"
      css={{
        ':hover': {
          color: active ? '' : 'var(--color-blue)'
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
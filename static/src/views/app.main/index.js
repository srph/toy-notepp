import React from 'react'
import {Div, A} from 'glamorous'
import NotePane from './NotePane'

class AppMain extends React.Component {
  render() {
    return (
      <UiWrapper>
        <UiSidebar>
          <UiSidebarLink icon="book" text="Notebook" active />
          <UiSidebarLink icon="trash" text="Trash" />
        </UiSidebar>

        <NotePane />
      </UuiWrapper>
    )
  }
}

function UiWrapper({children}) {
  return <Div
    display="flex"
    height="100vh"
    width="100%">
    {children}
  </Div>
}

function UiSidebar({children}) {
  return (
    <Div
      height="100%"
      width="160px"
      paddingTop="40px"
      paddingLeft="8px"
      paddingRight="8px"
      backgroundColor="var(--color-black)"
      overflowX="auto">
      {children}
    </Div>
  )
}

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

export default AppMain
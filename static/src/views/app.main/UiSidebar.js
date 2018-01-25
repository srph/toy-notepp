import React from 'react'
import {Div} from 'glamorous'

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

export default UiSidebar
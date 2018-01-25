import React from 'react'
import {Div} from 'glamorous'

function UiSidebar({children}) {
  return (
    <Div
      height="100%"
      width="160px"
      paddingTop="40px"
      backgroundColor="var(--color-dark-gray)"
      overflowX="auto">
      {children}
    </Div>
  )
}

export default UiSidebar
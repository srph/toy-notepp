import React from 'react'
import {Div, A, Input, H4, P} from 'glamorous'
import Searchbar from './Searchbar'
import NoteList from './NoteList'

class NotePane extends React.Component {
  render() {
    return (
      <UiWrapper>
        <Searchbar />
        <NoteList />
      </UiWrapper>
    )
  }
}

function UiWrapper({children}) {
  return <Div
    display="flex"
    flexDirection="column"
    height="100%"
    width="280px"
    borderRight="1px solid var(--color-silver)">
    {children}
  </Div>
}

export default NotePane
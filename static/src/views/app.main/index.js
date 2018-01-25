import React from 'react'
import {Div, A, I} from 'glamorous'
import UiSidebar from './UiSidebar'
import UiSidebarLink from './UiSidebarLink'
import NotePane from './NotePane'

export default class AppMain extends React.Component {
  render() {
    return (
      <Div
        display="flex"
        height="100vh"
        width="100%">
        <UiSidebar>
          <UiSidebarLink icon="book" text="Notebook" />
          <UiSidebarLink icon="trash" text="Trash" />
        </UiSidebar>

        <NotePane />
      </Div>
    )
  }
}
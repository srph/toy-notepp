import React from 'react'
import {Div, A, Input, H4, P} from 'glamorous'
import Searchbar from './Searchbar'

class NotePane extends React.Component {
  render() {
    return (
      <Div
        height="100%"
        width="280px"
        borderRight="1px solid var(--color-gray)"
        overflowY="auto">
        <Searchbar />
        {new Array(10).fill(null).map((note, i) =>
          <A
            display="flex"
            alignItems="center"
            height="120px"
            paddingTop="16px"
            paddingLeft="16px"
            width="100%"
            textDecoration="none"
            href="#">
            <Div
              alignSelf="flex-start"
              flexShrink="0"
              marginRight="16px"
              width="32px"
              textAlign="center">
              <Div
                marginBottom="16"
                color="var(--color-gray)">
                1h
              </Div>

              <Div color="var(--color-red)">
                <i className="fa fa-thumb-tack" />
              </Div>
            </Div>

            <Div
              paddingBottom="16px"
              paddingRight="16px"
              borderBottom="1px solid var(--color-gray)">
              <H4
                marginTop="0"
                marginBottom="8px"
                color="var(--color-black)">
                Grocery List
              </H4>

              <P
                margin="0"
                lineHeight="1.5"
                color="var(--color-gray)">
                Bear is the most elegant way to compose, sync,
                and store your notes.
              </P>
            </Div>
          </A>
        )}
      </Div>
    )
  }
}

export default NotePane
import React from 'react'
import {Button} from 'glamorous'

class UiButton extends React.Component {
  render() {
    return (
      <Button {...this.props}
        display="inline-block"
        padding="0px 8px"
        color="var(--color-lavender)"
        background="transparent"
        height="26px"
        fontSize="10px"
        fontWeight="500"
        lineHeight="25px"
        border="1px solid var(--color-silver)"
        borderRadius="var(--border-radius)"
        boxShadow="var(--drop-shadow)"
        cursor="pointer"
        textTransform="uppercase"
        fontFamily="var(--font-montserrat)"
        transition="all .2s cubic-bezier(.06,.67,.37,.99)"
        transform="translateY(0)"
        outline="0"
        opacity="1"
        css={{
          ':hover': {
            transform: 'translateY(-4px)',
            boxShadow: 'var(--drop-shadow-lower)'
          },
          ':disabled': {
            opacity: '0.7'
          }
        }} />
    )
  }
}

export default UiButton
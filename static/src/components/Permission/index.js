import React from 'react'
import history from '@/history'
import make from './make'

export default {
  /**
   * Only guests can enter this page
   */
  guest: make((state) => {
    if (state.auth.authenticated) {
      history.push('/')
      return false;
    }
  }),

  /**
   * Only authenticated users can enter this page
   */
  auth: make((state) => {
    if (!state.auth.authenticated) {
      history.push('/login')
      return false;
    }
  })
}
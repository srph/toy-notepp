import immer from 'immer'
import actions from './actions'

const init = {
  data: {},
  loading: false,
  error: ''
}

export function reducer(state = init, action) {
  return immer(state, (draft) => {
    switch(action.type) {
      case 'auth:init': {
        //
      }

      case 'auth:login!': {
        //
      }

      case 'auth:login.success': {
        //
      }

      case 'auth:login.error': {
        //
      }

      case 'auth:logout': {
        //
      }
    }
  })
}

export actions
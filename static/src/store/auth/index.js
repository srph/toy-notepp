import immer from 'immer'

const init = {
  data: {},
  authenticated: false,
  loading: false,
  error: ''
}

export function reducer(state = init, action) {
  return immer(state, (draft) => {
    switch(action.type) {
      case 'auth:init': {
        if (action.payload == null) {
          draft.authenticated = false
        } else {
          draft.data = action.payload
          draft.authenticated = true
        }
        break
      }

      case 'auth:login!': {
        draft.loading = true
        break
      }

      case 'auth:login.success': {
        draft.loading = false
        draft.data = action.payload
        draft.authenticated = true
        break
      }

      case 'auth:login.error': {
        draft.error = action.payload.error
        draft.loading = false
        break
      }

      case 'auth:logout': {
        draft.data = {}
        draft.authenticated = false
        break
      }
    }
  })
}

export * as actions from './actions'
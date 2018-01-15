import { ADD_URL, TOGGLE_MODAL } from '../actions'
import { combineReducers } from 'redux'

function urls(state = {}, action) {
  switch (action.type) {
    case ADD_URL: {
      return {
        originURL: action.originURL,
        shortURL: action.shortURL
      };
    }

    default: return state
  }
}

function visibilityFilter(state = false, action) {
  switch (action.type) {
    case TOGGLE_MODAL: {
      return !state;
    }

    default: return state
  }
}


export default combineReducers({ urls, visibilityFilter });
import { ADD_URL, ADD_ERROR, REMOVE_ERRORS, TOGGLE_MODAL } from '../actions';
import { combineReducers } from 'redux';

const urls = (state = {}, action) => {
  switch (action.type) {
    case ADD_URL: {
      return {
        originURL: action.originURL,
        shortURL: action.shortURL
      };
    }

    default: return state;
  }
};

const modalWindowVisibility = (state = false, action) => {
  switch (action.type) {
    case TOGGLE_MODAL: {
      return !state;
    }

    default: return state;
  }
};

const errors = (state = [], action) => {
  switch (action.type) {
    case ADD_ERROR: {
      return state.concat(action.error);
    };

    case REMOVE_ERRORS: {
      return [];
    };

    default: return state;
  }
}


export default combineReducers({ urls, modalWindowVisibility, errors });
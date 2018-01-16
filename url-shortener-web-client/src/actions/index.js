import axios from 'axios';

export const ADD_URL = 'ADD_URL';
export const ADD_ERROR = 'ADD_ERROR';
export const REMOVE_ERRORS = 'REMOVE_ERRORS';
export const TOGGLE_MODAL = 'TOGGLE_MODAL';

const SERVER_URL = 'http://localhost:3000';

export function addURL(originURL, shortURL) {
  return { type: ADD_URL, originURL, shortURL };
}

export function addError(error) {
  return { type: ADD_ERROR, error };
}

export function toggleModal() {
  return { type: TOGGLE_MODAL };
}

export function encodeURL(originURL) {
  return dispatch => {
    dispatch({
      type: REMOVE_ERRORS
    });

    return axios.post(`${SERVER_URL}/encode`, { 'url': originURL }).then(resp => {
      dispatch(addURL(originURL, resp.data.shortUrl));
      dispatch(toggleModal());
    }).catch(err => {
      dispatch(addError(err));
    });
  };
}
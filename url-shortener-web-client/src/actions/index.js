import axios from 'axios';

export const ADD_URL = 'ADD_URL'
export const TOGGLE_MODAL = 'TOGGLE_MODAL';

const SERVER_URL = "http://localhost:3000";

export function addURL(originURL, shortURL) {
  return { type: ADD_URL, originURL, shortURL };
}

export function toggleModal() {
  return { type: TOGGLE_MODAL }
}

export function encodeURL(originURL) {
  return dispatch => {
    return axios.post(`${SERVER_URL}/encode`, { "url": originURL }).then(resp => {
      dispatch(addURL(originURL, resp.data.shortUrl))
      dispatch(toggleModal());
    }).catch(err => {
      console.log(err);
    });
  }
}
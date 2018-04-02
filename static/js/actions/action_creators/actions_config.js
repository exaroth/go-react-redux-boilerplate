import * as actions from 'actions/types';
import { apiConnector } from 'utils/api';

export function setConfig(config) {
  return {
    type: actions.SET_CONFIG,
    payload: config,
  }
}

export function getConfig() {
  return (dispatch, getState)=> {
    apiConnector.getConfig().then(response=>{
      dispatch(setConfig(response))
    });
  }

}

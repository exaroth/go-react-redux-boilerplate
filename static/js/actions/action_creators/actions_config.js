import { SET_CONFIG } from "actions/types";
import { apiConnector } from "utils/api";

export function setConfig(config) {
  return {
    type: SET_CONFIG,
    payload: config
  };
}

export function getConfig() {
  return (dispatch, getState) => {
    apiConnector.getConfig().then(response => {
      dispatch(setConfig(response));
    });
  };
}

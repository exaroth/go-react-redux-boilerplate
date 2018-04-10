import { bindActionCreators, createStore, applyMiddleware } from "redux";
import { connect } from "react-redux";
import thunk from "redux-thunk";

import allActions from "actions/all";

export function connectedComponent(component, options) {
  return connect(
    state => state,
    dispatch => bindActionCreators(allActions, dispatch),
    (stateProps, dispatchProps, ownProps) =>
      Object.assign({}, ownProps, stateProps, dispatchProps),
    options
  )(component);
}

export function initStore(reducers, preloadedState) {
  return createStore(reducers, preloadedState, applyMiddleware(thunk));
}

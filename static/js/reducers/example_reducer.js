import * as actions from 'actions/types';

export default function(state = null, action) {
  switch(action.type) {
    case actions.EXAMPLE_ACTION:
      return action.payload;
    default:
      return state;
  }
}

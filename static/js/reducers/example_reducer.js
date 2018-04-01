import * as actions from 'actions/types';

export default function(state = null, action) {
  switch(action.type) {
    case actions.EXAMPLE_ACTION:
      console.log("Example has been set!")
      return state
    default:
      return state;
  }
}

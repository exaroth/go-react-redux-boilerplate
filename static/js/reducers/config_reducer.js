import { SET_CONFIG } from "actions/types";

// Set application config
export default function(state = null, action) {
  switch (action.type) {
    case SET_CONFIG:
      return action.payload;
    default:
      return state;
  }
}

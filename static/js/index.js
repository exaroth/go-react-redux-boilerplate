import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";

import { initStore } from "utils/react";
import allReducers from "reducers/all";
import App from "components/app";

function initApp() {
  let appContainer = (
    <Provider store={initStore(allReducers)}>
      <MuiThemeProvider>
        <App />
      </MuiThemeProvider>
    </Provider>
  );
  ReactDOM.render(appContainer, document.getElementById("app"));
}

document.addEventListener("DOMContentLoaded", initApp);

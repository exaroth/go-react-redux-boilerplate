import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { initStore } from 'utils/react' 
import reducersAggregator from 'reducers/store_aggregator';
import App from 'components/app';

function initApp() {
  let appContainer = (
    <Provider store={ initStore(reducersAggregator) }>
        <App />
    </Provider>
  );
  ReactDOM.render(appContainer, document.getElementById('app'));
}

document.addEventListener("DOMContentLoaded", initApp);

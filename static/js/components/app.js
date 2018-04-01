import React, { Component } from 'react';

import { connectedComponent } from 'utils/react';

class App extends Component {

  render() {
    return <h1>App</h1>
  }
}

export default connectedComponent(App)

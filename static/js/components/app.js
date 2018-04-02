import React, { Component } from 'react';

import { connectedComponent } from 'utils/react';

class App extends Component {

  componentDidMount() {
    this.props.getConfig();
  }

  componentDidUpdate() {
    console.log('updated')
    console.log(this.props)
  }

  render() {
    return <h1>App</h1>
  }
}

export default connectedComponent(App)

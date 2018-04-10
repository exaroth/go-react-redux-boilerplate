import React, { Component } from "react";

import { connectedComponent } from "utils/react";

class App extends Component {
  componentDidMount() {
    this.props.getConfig();
  }

  getVersion() {
    if (this.props.config) {
      return <span>{this.props.config.version}</span>;
    }
    return null;
  }

  render() {
    return <h1>App{this.getVersion()}</h1>;
  }
}

export default connectedComponent(App);

/*
 * Simple api middleware to streamline
 * calling of the API endpoints from
 * within the React app
 */

import axios from "axios";

class API {
  constructor() {
    this.apiEndpointBase = "/api";
    this.baseOptions = {
      responseType: "json",
      timeout: 60 * 1000,
      validateStatus: status => status >= 200 && status < 300,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json"
      }
    };
    this._call = this._call.bind(this);
    this._generateApiAddress = this._generateApiAddress.bind(this);
    this._getNewClient = this._getNewClient.bind(this);
  }

  _generateApiAddress(endpoint) {
    return `${this.apiEndpointBase}${endpoint}`;
  }

  _handleResponse(responsePromise) {
    return new Promise((resolve, reject) => {
      responsePromise.then(
        response => resolve(response.data),
        error => {
          if (error.response) {
            reject(error.response.data);
          } else {
            reject(error);
          }
        }
      );
    });
  }

  _getNewClient(overrideOptions = {}) {
    let options = Object.assign(overrideOptions, this.baseOptions);
    return axios.create(options);
  }

  _call(params) {
    const apiAddress = this._generateApiAddress(params.url);
    const makeRequest = this._getNewClient(params.options);

    let requestParams = {
      url: apiAddress,
      method: params.method,
      data: params.data
    };

    const response = makeRequest(requestParams);
    return this._handleResponse(response);
  }

  /*
  * Backend calls
  */
  getConfig() {
    return this._call({
      method: "get",
      url: "/config"
    });
  }
}

export const apiConnector = new API();

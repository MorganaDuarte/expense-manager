export default class RequestResponse{
  constructor(body, error) {
    this.body = body
    this.error = error
  }

  hasError() {
    return !!this.error;
  }

  getErrorMessage() {
    return this.error.message;
  }

  getBody() {
    return this.body;
  }
}